// Copyright 2024 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build windows

package hyperv

import (
	"fmt"
	"strings"

	"github.com/prometheus-community/windows_exporter/internal/perfdata"
	"github.com/prometheus-community/windows_exporter/internal/types"
	"github.com/prometheus/client_golang/prometheus"
)

// collectorHypervisorVirtualProcessor Hyper-V Hypervisor Virtual Processor metrics
type collectorHypervisorVirtualProcessor struct {
	perfDataCollectorHypervisorVirtualProcessor *perfdata.Collector

	// \Hyper-V Hypervisor Virtual Processor(*)\% Guest Run Time
	// \Hyper-V Hypervisor Virtual Processor(*)\% Hypervisor Run Time
	// \Hyper-V Hypervisor Virtual Processor(*)\% Remote Run Time
	hypervisorVirtualProcessorTimeTotal         *prometheus.Desc
	hypervisorVirtualProcessorTotalRunTimeTotal *prometheus.Desc // \Hyper-V Hypervisor Virtual Processor(*)\% Total Run Time
	hypervisorVirtualProcessorContextSwitches   *prometheus.Desc // \Hyper-V Hypervisor Virtual Processor(*)\CPU Wait Time Per Dispatch
}

const (
	hypervisorVirtualProcessorGuestIdleTimePercent     = "% Guest Idle Time"
	hypervisorVirtualProcessorHypervisorRunTimePercent = "% Hypervisor Run Time"
	hypervisorVirtualProcessorTotalRunTimePercent      = "% Total Run Time"
	hypervisorVirtualProcessorRemoteRunTimePercent     = "% Remote Run Time"
	hypervisorVirtualProcessorCPUWaitTimePerDispatch   = "CPU Wait Time Per Dispatch"
)

func (c *Collector) buildHypervisorVirtualProcessor() error {
	var err error

	c.perfDataCollectorHypervisorVirtualProcessor, err = perfdata.NewCollector("Hyper-V Hypervisor Virtual Processor", perfdata.InstancesAll, []string{
		hypervisorVirtualProcessorGuestIdleTimePercent,
		hypervisorVirtualProcessorHypervisorRunTimePercent,
		hypervisorVirtualProcessorTotalRunTimePercent,
		hypervisorVirtualProcessorRemoteRunTimePercent,
		hypervisorVirtualProcessorCPUWaitTimePerDispatch,
	})
	if err != nil {
		return fmt.Errorf("failed to create Hyper-V Hypervisor Virtual Processor collector: %w", err)
	}

	c.hypervisorVirtualProcessorTimeTotal = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "hypervisor_virtual_processor_time_total"),
		"Time that processor spent in different modes (hypervisor, guest_run, guest_idle, remote)",
		[]string{"vm", "core", "state"},
		nil,
	)
	c.hypervisorVirtualProcessorTotalRunTimeTotal = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "hypervisor_virtual_processor_total_run_time_total"),
		"Time that processor spent",
		[]string{"vm", "core"},
		nil,
	)
	c.hypervisorVirtualProcessorContextSwitches = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "hypervisor_virtual_processor_cpu_wait_time_per_dispatch_total"),
		"The average time (in nanoseconds) spent waiting for a virtual processor to be dispatched onto a logical processor.",
		[]string{"vm", "core"},
		nil,
	)

	return nil
}

func (c *Collector) collectHypervisorVirtualProcessor(ch chan<- prometheus.Metric) error {
	data, err := c.perfDataCollectorHypervisorVirtualProcessor.Collect()
	if err != nil {
		return fmt.Errorf("failed to collect Hyper-V Hypervisor Virtual Processor metrics: %w", err)
	}

	for coreName, coreData := range data {
		// The name format is <VM Name>:Hv VP <vcore id>
		parts := strings.Split(coreName, ":")
		if len(parts) != 2 {
			return fmt.Errorf("unexpected format of Name in Hyper-V Hypervisor Virtual Processor: %q, expected %q", coreName, "<VM Name>:Hv VP <vcore id>")
		}

		coreParts := strings.Split(parts[1], " ")
		if len(coreParts) != 3 {
			return fmt.Errorf("unexpected format of core identifier in Hyper-V Hypervisor Virtual Processor: %q, expected %q", parts[1], "Hv VP <vcore id>")
		}

		vmName := parts[0]
		coreId := coreParts[2]

		ch <- prometheus.MustNewConstMetric(
			c.hypervisorVirtualProcessorTimeTotal,
			prometheus.CounterValue,
			coreData[hypervisorVirtualProcessorHypervisorRunTimePercent].FirstValue,
			vmName, coreId, "hypervisor",
		)

		ch <- prometheus.MustNewConstMetric(
			c.hypervisorVirtualProcessorTimeTotal,
			prometheus.CounterValue,
			coreData[hypervisorVirtualProcessorGuestIdleTimePercent].FirstValue,
			vmName, coreId, "guest_idle",
		)

		ch <- prometheus.MustNewConstMetric(
			c.hypervisorVirtualProcessorTimeTotal,
			prometheus.CounterValue,
			coreData[hypervisorVirtualProcessorGuestIdleTimePercent].FirstValue,
			vmName, coreId, "guest_idle",
		)

		ch <- prometheus.MustNewConstMetric(
			c.hypervisorVirtualProcessorTotalRunTimeTotal,
			prometheus.CounterValue,
			coreData[hypervisorVirtualProcessorTotalRunTimePercent].FirstValue,
			vmName, coreId,
		)

		ch <- prometheus.MustNewConstMetric(
			c.hypervisorVirtualProcessorContextSwitches,
			prometheus.CounterValue,
			coreData[hypervisorVirtualProcessorCPUWaitTimePerDispatch].FirstValue,
			vmName, coreId,
		)
	}

	return nil
}

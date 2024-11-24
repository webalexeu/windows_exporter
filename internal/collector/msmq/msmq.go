//go:build windows

package msmq

import (
	"fmt"
	"log/slog"

	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus-community/windows_exporter/internal/mi"
	"github.com/prometheus-community/windows_exporter/internal/perfdata"
	"github.com/prometheus-community/windows_exporter/internal/types"
	"github.com/prometheus/client_golang/prometheus"
)

const Name = "msmq"

type Config struct{}

var ConfigDefaults = Config{}

// A Collector is a Prometheus Collector for WMI Win32_PerfRawData_MSMQ_MSMQQueue metrics.
type Collector struct {
	config            Config
	perfDataCollector *perfdata.Collector

	bytesInJournalQueue    *prometheus.Desc
	bytesInQueue           *prometheus.Desc
	messagesInJournalQueue *prometheus.Desc
	messagesInQueue        *prometheus.Desc
}

func New(config *Config) *Collector {
	if config == nil {
		config = &ConfigDefaults
	}

	c := &Collector{
		config: *config,
	}

	return c
}

func NewWithFlags(_ *kingpin.Application) *Collector {
	c := &Collector{
		config: ConfigDefaults,
	}

	return c
}

func (c *Collector) GetName() string {
	return Name
}

func (c *Collector) Close() error {
	return nil
}

func (c *Collector) Build(_ *slog.Logger, _ *mi.Session) error {
	var err error

	c.perfDataCollector, err = perfdata.NewCollector("MSMQ Queue", perfdata.InstancesAll, []string{
		BytesInJournalQueue,
		BytesInQueue,
		MessagesInJournalQueue,
		MessagesInQueue,
	})
	if err != nil {
		return fmt.Errorf("failed to create MSMQ Queue collector: %w", err)
	}

	c.bytesInJournalQueue = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "bytes_in_journal_queue"),
		"Size of queue journal in bytes",
		[]string{"name"},
		nil,
	)
	c.bytesInQueue = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "bytes_in_queue"),
		"Size of queue in bytes",
		[]string{"name"},
		nil,
	)
	c.messagesInJournalQueue = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "messages_in_journal_queue"),
		"Count messages in queue journal",
		[]string{"name"},
		nil,
	)
	c.messagesInQueue = prometheus.NewDesc(
		prometheus.BuildFQName(types.Namespace, Name, "messages_in_queue"),
		"Count messages in queue",
		[]string{"name"},
		nil,
	)

	return nil
}

// Collect sends the metric values for each metric
// to the provided prometheus Metric channel.
func (c *Collector) Collect(ch chan<- prometheus.Metric) error {
	perfData, err := c.perfDataCollector.Collect()
	if err != nil {
		return fmt.Errorf("failed to collect MSMQ Queue metrics: %w", err)
	}

	for name, data := range perfData {
		ch <- prometheus.MustNewConstMetric(
			c.bytesInJournalQueue,
			prometheus.GaugeValue,
			data[BytesInJournalQueue].FirstValue,
			name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.bytesInQueue,
			prometheus.GaugeValue,
			data[BytesInQueue].FirstValue,
			name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.messagesInJournalQueue,
			prometheus.GaugeValue,
			data[MessagesInJournalQueue].FirstValue,
			name,
		)

		ch <- prometheus.MustNewConstMetric(
			c.messagesInQueue,
			prometheus.GaugeValue,
			data[MessagesInQueue].FirstValue,
			name,
		)
	}

	return nil
}

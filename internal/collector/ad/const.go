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

package ad

const (
	abANRPerSec                                                      = "AB ANR/sec"
	abBrowsesPerSec                                                  = "AB Browses/sec"
	abClientSessions                                                 = "AB Client Sessions"
	abMatchesPerSec                                                  = "AB Matches/sec"
	abPropertyReadsPerSec                                            = "AB Property Reads/sec"
	abProxyLookupsPerSec                                             = "AB Proxy Lookups/sec"
	abSearchesPerSec                                                 = "AB Searches/sec"
	approximateHighestDNT                                            = "Approximate highest DNT"
	atqEstimatedQueueDelay                                           = "ATQ Estimated Queue Delay"
	atqOutstandingQueuedRequests                                     = "ATQ Outstanding Queued Requests"
	_                                                                = "ATQ Queue Latency"
	atqRequestLatency                                                = "ATQ Request Latency"
	atqThreadsLDAP                                                   = "ATQ Threads LDAP"
	atqThreadsOther                                                  = "ATQ Threads Other"
	atqThreadsTotal                                                  = "ATQ Threads Total"
	baseSearchesPerSec                                               = "Base searches/sec"
	databaseAddsPerSec                                               = "Database adds/sec"
	databaseDeletesPerSec                                            = "Database deletes/sec"
	databaseModifiesPerSec                                           = "Database modifys/sec"
	databaseRecyclesPerSec                                           = "Database recycles/sec"
	digestBindsPerSec                                                = "Digest Binds/sec"
	_                                                                = "DirSync session throttling rate"
	_                                                                = "DirSync sessions in progress"
	draHighestUSNCommittedHighPart                                   = "DRA Highest USN Committed (High part)"
	draHighestUSNCommittedLowPart                                    = "DRA Highest USN Committed (Low part)"
	draHighestUSNIssuedHighPart                                      = "DRA Highest USN Issued (High part)"
	draHighestUSNIssuedLowPart                                       = "DRA Highest USN Issued (Low part)"
	draInboundBytesCompressedBetweenSitesAfterCompressionSinceBoot   = "DRA Inbound Bytes Compressed (Between Sites, After Compression) Since Boot"
	draInboundBytesCompressedBetweenSitesAfterCompressionPerSec      = "DRA Inbound Bytes Compressed (Between Sites, After Compression)/sec"
	draInboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot  = "DRA Inbound Bytes Compressed (Between Sites, Before Compression) Since Boot"
	draInboundBytesCompressedBetweenSitesBeforeCompressionPerSec     = "DRA Inbound Bytes Compressed (Between Sites, Before Compression)/sec"
	draInboundBytesNotCompressedWithinSiteSinceBoot                  = "DRA Inbound Bytes Not Compressed (Within Site) Since Boot"
	draInboundBytesNotCompressedWithinSitePerSec                     = "DRA Inbound Bytes Not Compressed (Within Site)/sec"
	draInboundBytesTotalSinceBoot                                    = "DRA Inbound Bytes Total Since Boot"
	draInboundBytesTotalPerSec                                       = "DRA Inbound Bytes Total/sec"
	draInboundFullSyncObjectsRemaining                               = "DRA Inbound Full Sync Objects Remaining"
	draInboundLinkValueUpdatesRemainingInPacket                      = "DRA Inbound Link Value Updates Remaining in Packet"
	_                                                                = "DRA Inbound Link Values/sec"
	draInboundObjectUpdatesRemainingInPacket                         = "DRA Inbound Object Updates Remaining in Packet"
	draInboundObjectsAppliedPerSec                                   = "DRA Inbound Objects Applied/sec"
	draInboundObjectsFilteredPerSec                                  = "DRA Inbound Objects Filtered/sec"
	draInboundObjectsPerSec                                          = "DRA Inbound Objects/sec"
	draInboundPropertiesAppliedPerSec                                = "DRA Inbound Properties Applied/sec"
	draInboundPropertiesFilteredPerSec                               = "DRA Inbound Properties Filtered/sec"
	draInboundPropertiesTotalPerSec                                  = "DRA Inbound Properties Total/sec"
	_                                                                = "DRA Inbound Sync Link Deletion/sec"
	draInboundTotalUpdatesRemainingInPacket                          = "DRA Inbound Total Updates Remaining in Packet"
	draInboundValuesDNsOnlyPerSec                                    = "DRA Inbound Values (DNs only)/sec"
	draInboundValuesTotalPerSec                                      = "DRA Inbound Values Total/sec"
	_                                                                = "DRA number of NC replication calls since boot"
	_                                                                = "DRA number of successful NC replication calls since boot"
	draOutboundBytesCompressedBetweenSitesAfterCompressionSinceBoot  = "DRA Outbound Bytes Compressed (Between Sites, After Compression) Since Boot"
	draOutboundBytesCompressedBetweenSitesAfterCompressionPerSec     = "DRA Outbound Bytes Compressed (Between Sites, After Compression)/sec"
	draOutboundBytesCompressedBetweenSitesBeforeCompressionSinceBoot = "DRA Outbound Bytes Compressed (Between Sites, Before Compression) Since Boot"
	draOutboundBytesCompressedBetweenSitesBeforeCompressionPerSec    = "DRA Outbound Bytes Compressed (Between Sites, Before Compression)/sec"
	draOutboundBytesNotCompressedWithinSiteSinceBoot                 = "DRA Outbound Bytes Not Compressed (Within Site) Since Boot"
	draOutboundBytesNotCompressedWithinSitePerSec                    = "DRA Outbound Bytes Not Compressed (Within Site)/sec"
	draOutboundBytesTotalSinceBoot                                   = "DRA Outbound Bytes Total Since Boot"
	draOutboundBytesTotalPerSec                                      = "DRA Outbound Bytes Total/sec"
	draOutboundObjectsFilteredPerSec                                 = "DRA Outbound Objects Filtered/sec"
	draOutboundObjectsPerSec                                         = "DRA Outbound Objects/sec"
	draOutboundPropertiesPerSec                                      = "DRA Outbound Properties/sec"
	draOutboundValuesDNsOnlyPerSec                                   = "DRA Outbound Values (DNs only)/sec"
	draOutboundValuesTotalPerSec                                     = "DRA Outbound Values Total/sec"
	draPendingReplicationOperations                                  = "DRA Pending Replication Operations"
	draPendingReplicationSynchronizations                            = "DRA Pending Replication Synchronizations"
	draSyncFailuresOnSchemaMismatch                                  = "DRA Sync Failures on Schema Mismatch"
	draSyncRequestsMade                                              = "DRA Sync Requests Made"
	draSyncRequestsSuccessful                                        = "DRA Sync Requests Successful"
	draThreadsGettingNCChanges                                       = "DRA Threads Getting NC Changes"
	draThreadsGettingNCChangesHoldingSemaphore                       = "DRA Threads Getting NC Changes Holding Semaphore"
	_                                                                = "DRA total number of Busy failures since boot"
	_                                                                = "DRA total number of MissingParent failures since boot"
	_                                                                = "DRA total number of NotEnoughAttrs/MissingObject failures since boot"
	_                                                                = "DRA total number of Preempted failures since boot"
	_                                                                = "DRA total time of applying replication package since boot"
	_                                                                = "DRA total time of NC replication calls since boot"
	_                                                                = "DRA total time of successful NC replication calls since boot"
	_                                                                = "DRA total time of successfully applying replication package since boot"
	_                                                                = "DRA total time on waiting async replication packages since boot"
	_                                                                = "DRA total time on waiting sync replication packages since boot"
	dsPercentReadsFromDRA                                            = "DS % Reads from DRA"
	dsPercentReadsFromKCC                                            = "DS % Reads from KCC"
	dsPercentReadsFromLSA                                            = "DS % Reads from LSA"
	dsPercentReadsFromNSPI                                           = "DS % Reads from NSPI"
	dsPercentReadsFromNTDSAPI                                        = "DS % Reads from NTDSAPI"
	dsPercentReadsFromSAM                                            = "DS % Reads from SAM"
	dsPercentReadsOther                                              = "DS % Reads Other"
	dsPercentSearchesFromDRA                                         = "DS % Searches from DRA"
	dsPercentSearchesFromKCC                                         = "DS % Searches from KCC"
	dsPercentSearchesFromLDAP                                        = "DS % Searches from LDAP"
	dsPercentSearchesFromLSA                                         = "DS % Searches from LSA"
	dsPercentSearchesFromNSPI                                        = "DS % Searches from NSPI"
	dsPercentSearchesFromNTDSAPI                                     = "DS % Searches from NTDSAPI"
	dsPercentSearchesFromSAM                                         = "DS % Searches from SAM"
	dsPercentSearchesOther                                           = "DS % Searches Other"
	dsPercentWritesFromDRA                                           = "DS % Writes from DRA"
	dsPercentWritesFromKCC                                           = "DS % Writes from KCC"
	dsPercentWritesFromLDAP                                          = "DS % Writes from LDAP"
	dsPercentWritesFromLSA                                           = "DS % Writes from LSA"
	dsPercentWritesFromNSPI                                          = "DS % Writes from NSPI"
	dsPercentWritesFromNTDSAPI                                       = "DS % Writes from NTDSAPI"
	dsPercentWritesFromSAM                                           = "DS % Writes from SAM"
	dsPercentWritesOther                                             = "DS % Writes Other"
	dsClientBindsPerSec                                              = "DS Client Binds/sec"
	dsClientNameTranslationsPerSec                                   = "DS Client Name Translations/sec"
	dsDirectoryReadsPerSec                                           = "DS Directory Reads/sec"
	dsDirectorySearchesPerSec                                        = "DS Directory Searches/sec"
	dsDirectoryWritesPerSec                                          = "DS Directory Writes/sec"
	dsMonitorListSize                                                = "DS Monitor List Size"
	dsNameCacheHitRate                                               = "DS Name Cache hit rate"
	dsNotifyQueueSize                                                = "DS Notify Queue Size"
	dsSearchSubOperationsPerSec                                      = "DS Search sub-operations/sec"
	dsSecurityDescriptorPropagationsEvents                           = "DS Security Descriptor Propagations Events"
	dsSecurityDescriptorPropagatorAverageExclusionTime               = "DS Security Descriptor Propagator Average Exclusion Time"
	dsSecurityDescriptorPropagatorRuntimeQueue                       = "DS Security Descriptor Propagator Runtime Queue"
	dsSecurityDescriptorSubOperationsPerSec                          = "DS Security Descriptor sub-operations/sec"
	dsServerBindsPerSec                                              = "DS Server Binds/sec"
	dsServerNameTranslationsPerSec                                   = "DS Server Name Translations/sec"
	dsThreadsInUse                                                   = "DS Threads in Use"
	_                                                                = "Error eventlogs since boot"
	_                                                                = "Error events since boot"
	externalBindsPerSec                                              = "External Binds/sec"
	fastBindsPerSec                                                  = "Fast Binds/sec"
	_                                                                = "Fatal events since boot"
	_                                                                = "Info eventlogs since boot"
	ldapActiveThreads                                                = "LDAP Active Threads"
	_                                                                = "LDAP Add Operations"
	_                                                                = "LDAP Add Operations/sec"
	_                                                                = "LDAP batch slots available"
	ldapBindTime                                                     = "LDAP Bind Time"
	_                                                                = "LDAP busy retries"
	_                                                                = "LDAP busy retries/sec"
	ldapClientSessions                                               = "LDAP Client Sessions"
	ldapClosedConnectionsPerSec                                      = "LDAP Closed Connections/sec"
	_                                                                = "LDAP Delete Operations"
	_                                                                = "LDAP Delete Operations/sec"
	_                                                                = "LDAP Modify DN Operations"
	_                                                                = "LDAP Modify DN Operations/sec"
	_                                                                = "LDAP Modify Operations"
	_                                                                = "LDAP Modify Operations/sec"
	ldapNewConnectionsPerSec                                         = "LDAP New Connections/sec"
	ldapNewSSLConnectionsPerSec                                      = "LDAP New SSL Connections/sec"
	_                                                                = "LDAP Outbound Bytes"
	_                                                                = "LDAP Outbound Bytes/sec"
	_                                                                = "LDAP Page Search Cache entries count"
	_                                                                = "LDAP Page Search Cache size"
	ldapSearchesPerSec                                               = "LDAP Searches/sec"
	ldapSuccessfulBindsPerSec                                        = "LDAP Successful Binds/sec"
	_                                                                = "LDAP Threads Sleeping on BUSY"
	ldapUDPOperationsPerSec                                          = "LDAP UDP operations/sec"
	ldapWritesPerSec                                                 = "LDAP Writes/sec"
	linkValuesCleanedPerSec                                          = "Link Values Cleaned/sec"
	_                                                                = "Links added"
	_                                                                = "Links added/sec"
	_                                                                = "Links visited"
	_                                                                = "Links visited/sec"
	_                                                                = "Logical link deletes"
	_                                                                = "Logical link deletes/sec"
	negotiatedBindsPerSec                                            = "Negotiated Binds/sec"
	ntlmBindsPerSec                                                  = "NTLM Binds/sec"
	_                                                                = "Objects returned"
	_                                                                = "Objects returned/sec"
	_                                                                = "Objects visited"
	_                                                                = "Objects visited/sec"
	oneLevelSearchesPerSec                                           = "Onelevel searches/sec"
	_                                                                = "PDC failed password update notifications"
	_                                                                = "PDC password update notifications/sec"
	_                                                                = "PDC successful password update notifications"
	phantomsCleanedPerSec                                            = "Phantoms Cleaned/sec"
	phantomsVisitedPerSec                                            = "Phantoms Visited/sec"
	_                                                                = "Physical link deletes"
	_                                                                = "Physical link deletes/sec"
	_                                                                = "Replicate Single Object operations"
	_                                                                = "Replicate Single Object operations/sec"
	_                                                                = "RID Pool invalidations since boot"
	_                                                                = "RID Pool request failures since boot"
	_                                                                = "RID Pool request successes since boot"
	samAccountGroupEvaluationLatency                                 = "SAM Account Group Evaluation Latency"
	samDisplayInformationQueriesPerSec                               = "SAM Display Information Queries/sec"
	samDomainLocalGroupMembershipEvaluationsPerSec                   = "SAM Domain Local Group Membership Evaluations/sec"
	samEnumerationsPerSec                                            = "SAM Enumerations/sec"
	samGCEvaluationsPerSec                                           = "SAM GC Evaluations/sec"
	samGlobalGroupMembershipEvaluationsPerSec                        = "SAM Global Group Membership Evaluations/sec"
	samMachineCreationAttemptsPerSec                                 = "SAM Machine Creation Attempts/sec"
	samMembershipChangesPerSec                                       = "SAM Membership Changes/sec"
	samNonTransitiveMembershipEvaluationsPerSec                      = "SAM Non-Transitive Membership Evaluations/sec"
	samPasswordChangesPerSec                                         = "SAM Password Changes/sec"
	samResourceGroupEvaluationLatency                                = "SAM Resource Group Evaluation Latency"
	samSuccessfulComputerCreationsPerSecIncludesAllRequests          = "SAM Successful Computer Creations/sec: Includes all requests"
	samSuccessfulUserCreationsPerSec                                 = "SAM Successful User Creations/sec"
	samTransitiveMembershipEvaluationsPerSec                         = "SAM Transitive Membership Evaluations/sec"
	samUniversalGroupMembershipEvaluationsPerSec                     = "SAM Universal Group Membership Evaluations/sec"
	samUserCreationAttemptsPerSec                                    = "SAM User Creation Attempts/sec"
	simpleBindsPerSec                                                = "Simple Binds/sec"
	subtreeSearchesPerSec                                            = "Subtree searches/sec"
	tombstonesGarbageCollectedPerSec                                 = "Tombstones Garbage Collected/sec"
	tombstonesVisitedPerSec                                          = "Tombstones Visited/sec"
	transitiveOperationsMillisecondsRun                              = "Transitive operations milliseconds run"
	transitiveOperationsPerSec                                       = "Transitive operations/sec"
	transitiveSubOperationsPerSec                                    = "Transitive suboperations/sec"
	_                                                                = "Warning eventlogs since boot"
	_                                                                = "Warning events since boot"
)

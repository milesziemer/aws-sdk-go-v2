// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CancellationStatus string

// Enum values for CancellationStatus
const (
	CancellationStatusCancelling CancellationStatus = "CANCELLING"
	CancellationStatusCancelled  CancellationStatus = "CANCELLED"
)

// Values returns all known values for CancellationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CancellationStatus) Values() []CancellationStatus {
	return []CancellationStatus{
		"CANCELLING",
		"CANCELLED",
	}
}

type CompressionType string

// Enum values for CompressionType
const (
	CompressionTypeNone CompressionType = "NONE"
	CompressionTypeGzip CompressionType = "GZIP"
)

// Values returns all known values for CompressionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CompressionType) Values() []CompressionType {
	return []CompressionType{
		"NONE",
		"GZIP",
	}
}

type DeviceStatus string

// Enum values for DeviceStatus
const (
	DeviceStatusOnline  DeviceStatus = "ONLINE"
	DeviceStatusOffline DeviceStatus = "OFFLINE"
	DeviceStatusRetired DeviceStatus = "RETIRED"
)

// Values returns all known values for DeviceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeviceStatus) Values() []DeviceStatus {
	return []DeviceStatus{
		"ONLINE",
		"OFFLINE",
		"RETIRED",
	}
}

type DeviceType string

// Enum values for DeviceType
const (
	DeviceTypeQpu       DeviceType = "QPU"
	DeviceTypeSimulator DeviceType = "SIMULATOR"
)

// Values returns all known values for DeviceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DeviceType) Values() []DeviceType {
	return []DeviceType{
		"QPU",
		"SIMULATOR",
	}
}

type HybridJobAdditionalAttributeName string

// Enum values for HybridJobAdditionalAttributeName
const (
	HybridJobAdditionalAttributeNameQueueInfo HybridJobAdditionalAttributeName = "QueueInfo"
)

// Values returns all known values for HybridJobAdditionalAttributeName. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (HybridJobAdditionalAttributeName) Values() []HybridJobAdditionalAttributeName {
	return []HybridJobAdditionalAttributeName{
		"QueueInfo",
	}
}

type InstanceType string

// Enum values for InstanceType
const (
	InstanceTypeMlM4Xlarge     InstanceType = "ml.m4.xlarge"
	InstanceTypeMlM42xlarge    InstanceType = "ml.m4.2xlarge"
	InstanceTypeMlM44xlarge    InstanceType = "ml.m4.4xlarge"
	InstanceTypeMlM410xlarge   InstanceType = "ml.m4.10xlarge"
	InstanceTypeMlM416xlarge   InstanceType = "ml.m4.16xlarge"
	InstanceTypeMlG4dnXlarge   InstanceType = "ml.g4dn.xlarge"
	InstanceTypeMlG4dn2xlarge  InstanceType = "ml.g4dn.2xlarge"
	InstanceTypeMlG4dn4xlarge  InstanceType = "ml.g4dn.4xlarge"
	InstanceTypeMlG4dn8xlarge  InstanceType = "ml.g4dn.8xlarge"
	InstanceTypeMlG4dn12xlarge InstanceType = "ml.g4dn.12xlarge"
	InstanceTypeMlG4dn16xlarge InstanceType = "ml.g4dn.16xlarge"
	InstanceTypeMlM5Large      InstanceType = "ml.m5.large"
	InstanceTypeMlM5Xlarge     InstanceType = "ml.m5.xlarge"
	InstanceTypeMlM52xlarge    InstanceType = "ml.m5.2xlarge"
	InstanceTypeMlM54xlarge    InstanceType = "ml.m5.4xlarge"
	InstanceTypeMlM512xlarge   InstanceType = "ml.m5.12xlarge"
	InstanceTypeMlM524xlarge   InstanceType = "ml.m5.24xlarge"
	InstanceTypeMlC4Xlarge     InstanceType = "ml.c4.xlarge"
	InstanceTypeMlC42xlarge    InstanceType = "ml.c4.2xlarge"
	InstanceTypeMlC44xlarge    InstanceType = "ml.c4.4xlarge"
	InstanceTypeMlC48xlarge    InstanceType = "ml.c4.8xlarge"
	InstanceTypeMlP2Xlarge     InstanceType = "ml.p2.xlarge"
	InstanceTypeMlP28xlarge    InstanceType = "ml.p2.8xlarge"
	InstanceTypeMlP216xlarge   InstanceType = "ml.p2.16xlarge"
	InstanceTypeMlP32xlarge    InstanceType = "ml.p3.2xlarge"
	InstanceTypeMlP38xlarge    InstanceType = "ml.p3.8xlarge"
	InstanceTypeMlP316xlarge   InstanceType = "ml.p3.16xlarge"
	InstanceTypeMlP3dn24xlarge InstanceType = "ml.p3dn.24xlarge"
	InstanceTypeMlP4d24xlarge  InstanceType = "ml.p4d.24xlarge"
	InstanceTypeMlC5Xlarge     InstanceType = "ml.c5.xlarge"
	InstanceTypeMlC52xlarge    InstanceType = "ml.c5.2xlarge"
	InstanceTypeMlC54xlarge    InstanceType = "ml.c5.4xlarge"
	InstanceTypeMlC59xlarge    InstanceType = "ml.c5.9xlarge"
	InstanceTypeMlC518xlarge   InstanceType = "ml.c5.18xlarge"
	InstanceTypeMlC5nXlarge    InstanceType = "ml.c5n.xlarge"
	InstanceTypeMlC5n2xlarge   InstanceType = "ml.c5n.2xlarge"
	InstanceTypeMlC5n4xlarge   InstanceType = "ml.c5n.4xlarge"
	InstanceTypeMlC5n9xlarge   InstanceType = "ml.c5n.9xlarge"
	InstanceTypeMlC5n18xlarge  InstanceType = "ml.c5n.18xlarge"
)

// Values returns all known values for InstanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceType) Values() []InstanceType {
	return []InstanceType{
		"ml.m4.xlarge",
		"ml.m4.2xlarge",
		"ml.m4.4xlarge",
		"ml.m4.10xlarge",
		"ml.m4.16xlarge",
		"ml.g4dn.xlarge",
		"ml.g4dn.2xlarge",
		"ml.g4dn.4xlarge",
		"ml.g4dn.8xlarge",
		"ml.g4dn.12xlarge",
		"ml.g4dn.16xlarge",
		"ml.m5.large",
		"ml.m5.xlarge",
		"ml.m5.2xlarge",
		"ml.m5.4xlarge",
		"ml.m5.12xlarge",
		"ml.m5.24xlarge",
		"ml.c4.xlarge",
		"ml.c4.2xlarge",
		"ml.c4.4xlarge",
		"ml.c4.8xlarge",
		"ml.p2.xlarge",
		"ml.p2.8xlarge",
		"ml.p2.16xlarge",
		"ml.p3.2xlarge",
		"ml.p3.8xlarge",
		"ml.p3.16xlarge",
		"ml.p3dn.24xlarge",
		"ml.p4d.24xlarge",
		"ml.c5.xlarge",
		"ml.c5.2xlarge",
		"ml.c5.4xlarge",
		"ml.c5.9xlarge",
		"ml.c5.18xlarge",
		"ml.c5n.xlarge",
		"ml.c5n.2xlarge",
		"ml.c5n.4xlarge",
		"ml.c5n.9xlarge",
		"ml.c5n.18xlarge",
	}
}

type JobEventType string

// Enum values for JobEventType
const (
	JobEventTypeWaitingForPriority           JobEventType = "WAITING_FOR_PRIORITY"
	JobEventTypeQueuedForExecution           JobEventType = "QUEUED_FOR_EXECUTION"
	JobEventTypeStartingInstance             JobEventType = "STARTING_INSTANCE"
	JobEventTypeDownloadingData              JobEventType = "DOWNLOADING_DATA"
	JobEventTypeRunning                      JobEventType = "RUNNING"
	JobEventTypeDeprioritizedDueToInactivity JobEventType = "DEPRIORITIZED_DUE_TO_INACTIVITY"
	JobEventTypeUploadingResults             JobEventType = "UPLOADING_RESULTS"
	JobEventTypeCompleted                    JobEventType = "COMPLETED"
	JobEventTypeFailed                       JobEventType = "FAILED"
	JobEventTypeMaxRuntimeExceeded           JobEventType = "MAX_RUNTIME_EXCEEDED"
	JobEventTypeCancelled                    JobEventType = "CANCELLED"
)

// Values returns all known values for JobEventType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JobEventType) Values() []JobEventType {
	return []JobEventType{
		"WAITING_FOR_PRIORITY",
		"QUEUED_FOR_EXECUTION",
		"STARTING_INSTANCE",
		"DOWNLOADING_DATA",
		"RUNNING",
		"DEPRIORITIZED_DUE_TO_INACTIVITY",
		"UPLOADING_RESULTS",
		"COMPLETED",
		"FAILED",
		"MAX_RUNTIME_EXCEEDED",
		"CANCELLED",
	}
}

type JobPrimaryStatus string

// Enum values for JobPrimaryStatus
const (
	JobPrimaryStatusQueued     JobPrimaryStatus = "QUEUED"
	JobPrimaryStatusRunning    JobPrimaryStatus = "RUNNING"
	JobPrimaryStatusCompleted  JobPrimaryStatus = "COMPLETED"
	JobPrimaryStatusFailed     JobPrimaryStatus = "FAILED"
	JobPrimaryStatusCancelling JobPrimaryStatus = "CANCELLING"
	JobPrimaryStatusCancelled  JobPrimaryStatus = "CANCELLED"
)

// Values returns all known values for JobPrimaryStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JobPrimaryStatus) Values() []JobPrimaryStatus {
	return []JobPrimaryStatus{
		"QUEUED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
		"CANCELLING",
		"CANCELLED",
	}
}

type QuantumTaskAdditionalAttributeName string

// Enum values for QuantumTaskAdditionalAttributeName
const (
	QuantumTaskAdditionalAttributeNameQueueInfo QuantumTaskAdditionalAttributeName = "QueueInfo"
)

// Values returns all known values for QuantumTaskAdditionalAttributeName. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (QuantumTaskAdditionalAttributeName) Values() []QuantumTaskAdditionalAttributeName {
	return []QuantumTaskAdditionalAttributeName{
		"QueueInfo",
	}
}

type QuantumTaskStatus string

// Enum values for QuantumTaskStatus
const (
	QuantumTaskStatusCreated    QuantumTaskStatus = "CREATED"
	QuantumTaskStatusQueued     QuantumTaskStatus = "QUEUED"
	QuantumTaskStatusRunning    QuantumTaskStatus = "RUNNING"
	QuantumTaskStatusCompleted  QuantumTaskStatus = "COMPLETED"
	QuantumTaskStatusFailed     QuantumTaskStatus = "FAILED"
	QuantumTaskStatusCancelling QuantumTaskStatus = "CANCELLING"
	QuantumTaskStatusCancelled  QuantumTaskStatus = "CANCELLED"
)

// Values returns all known values for QuantumTaskStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (QuantumTaskStatus) Values() []QuantumTaskStatus {
	return []QuantumTaskStatus{
		"CREATED",
		"QUEUED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
		"CANCELLING",
		"CANCELLED",
	}
}

type QueueName string

// Enum values for QueueName
const (
	QueueNameQuantumTasksQueue QueueName = "QUANTUM_TASKS_QUEUE"
	QueueNameJobsQueue         QueueName = "JOBS_QUEUE"
)

// Values returns all known values for QueueName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (QueueName) Values() []QueueName {
	return []QueueName{
		"QUANTUM_TASKS_QUEUE",
		"JOBS_QUEUE",
	}
}

type QueuePriority string

// Enum values for QueuePriority
const (
	QueuePriorityNormal   QueuePriority = "Normal"
	QueuePriorityPriority QueuePriority = "Priority"
)

// Values returns all known values for QueuePriority. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (QueuePriority) Values() []QueuePriority {
	return []QueuePriority{
		"Normal",
		"Priority",
	}
}

type SearchJobsFilterOperator string

// Enum values for SearchJobsFilterOperator
const (
	SearchJobsFilterOperatorLt       SearchJobsFilterOperator = "LT"
	SearchJobsFilterOperatorLte      SearchJobsFilterOperator = "LTE"
	SearchJobsFilterOperatorEqual    SearchJobsFilterOperator = "EQUAL"
	SearchJobsFilterOperatorGt       SearchJobsFilterOperator = "GT"
	SearchJobsFilterOperatorGte      SearchJobsFilterOperator = "GTE"
	SearchJobsFilterOperatorBetween  SearchJobsFilterOperator = "BETWEEN"
	SearchJobsFilterOperatorContains SearchJobsFilterOperator = "CONTAINS"
)

// Values returns all known values for SearchJobsFilterOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SearchJobsFilterOperator) Values() []SearchJobsFilterOperator {
	return []SearchJobsFilterOperator{
		"LT",
		"LTE",
		"EQUAL",
		"GT",
		"GTE",
		"BETWEEN",
		"CONTAINS",
	}
}

type SearchQuantumTasksFilterOperator string

// Enum values for SearchQuantumTasksFilterOperator
const (
	SearchQuantumTasksFilterOperatorLt      SearchQuantumTasksFilterOperator = "LT"
	SearchQuantumTasksFilterOperatorLte     SearchQuantumTasksFilterOperator = "LTE"
	SearchQuantumTasksFilterOperatorEqual   SearchQuantumTasksFilterOperator = "EQUAL"
	SearchQuantumTasksFilterOperatorGt      SearchQuantumTasksFilterOperator = "GT"
	SearchQuantumTasksFilterOperatorGte     SearchQuantumTasksFilterOperator = "GTE"
	SearchQuantumTasksFilterOperatorBetween SearchQuantumTasksFilterOperator = "BETWEEN"
)

// Values returns all known values for SearchQuantumTasksFilterOperator. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (SearchQuantumTasksFilterOperator) Values() []SearchQuantumTasksFilterOperator {
	return []SearchQuantumTasksFilterOperator{
		"LT",
		"LTE",
		"EQUAL",
		"GT",
		"GTE",
		"BETWEEN",
	}
}

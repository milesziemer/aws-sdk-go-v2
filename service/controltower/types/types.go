// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An operation performed by the control.
type ControlOperation struct {

	// The time that the operation finished.
	EndTime *time.Time

	// One of ENABLE_CONTROL or DISABLE_CONTROL .
	OperationType ControlOperationType

	// The time that the operation began.
	StartTime *time.Time

	// One of IN_PROGRESS , SUCEEDED , or FAILED .
	Status ControlOperationStatus

	// If the operation result is FAILED , this string contains a message explaining
	// why the operation failed.
	StatusMessage *string

	noSmithyDocumentSerde
}

// The drift summary of the enabled control. AWS Control Tower expects the enabled
// control configuration to include all supported and governed Regions. If the
// enabled control differs from the expected configuration, it is defined to be in
// a state of drift. You can repair this drift by resetting the enabled control.
type DriftStatusSummary struct {

	// The drift status of the enabled control. Valid values:
	//   - DRIFTED : The enabledControl deployed in this configuration doesn’t match
	//   the configuration that AWS Control Tower expected.
	//   - IN_SYNC : The enabledControl deployed in this configuration matches the
	//   configuration that AWS Control Tower expected.
	//   - NOT_CHECKING : AWS Control Tower does not check drift for this enabled
	//   control. Drift is not supported for the control type.
	//   - UNKNOWN : AWS Control Tower is not able to check the drift status for the
	//   enabled control.
	DriftStatus DriftStatus

	noSmithyDocumentSerde
}

// Information about the enabled control.
type EnabledControlDetails struct {

	// The ARN of the enabled control.
	Arn *string

	// The control identifier of the enabled control. For information on how to find
	// the controlIdentifier , see the overview page (https://docs.aws.amazon.com/controltower/latest/APIReference/Welcome.html)
	// .
	ControlIdentifier *string

	// The drift status of the enabled control.
	DriftStatusSummary *DriftStatusSummary

	// The deployment summary of the enabled control.
	StatusSummary *EnablementStatusSummary

	// The ARN of the organizational unit. For information on how to find the
	// targetIdentifier , see the overview page (https://docs.aws.amazon.com/controltower/latest/APIReference/Welcome.html)
	// .
	TargetIdentifier *string

	// Target AWS Regions for the enabled control.
	TargetRegions []Region

	noSmithyDocumentSerde
}

// Returns a summary of information about an enabled control.
type EnabledControlSummary struct {

	// The ARN of the enabled control.
	Arn *string

	// The controlIdentifier of the enabled control.
	ControlIdentifier *string

	// The drift status of the enabled control.
	DriftStatusSummary *DriftStatusSummary

	// A short description of the status of the enabled control.
	StatusSummary *EnablementStatusSummary

	// The ARN of the organizational unit.
	TargetIdentifier *string

	noSmithyDocumentSerde
}

// The deployment summary of the enabled control.
type EnablementStatusSummary struct {

	// The last operation identifier for the enabled control.
	LastOperationIdentifier *string

	// The deployment status of the enabled control. Valid values:
	//   - SUCCEEDED : The enabledControl configuration was deployed successfully.
	//   - UNDER_CHANGE : The enabledControl configuration is changing.
	//   - FAILED : The enabledControl configuration failed to deploy.
	Status EnablementStatus

	noSmithyDocumentSerde
}

// An AWS Region in which AWS Control Tower expects to find the control deployed.
// The expected Regions are based on the Regions that are governed by the landing
// zone. In certain cases, a control is not actually enabled in the Region as
// expected, such as during drift, or mixed governance (https://docs.aws.amazon.com/controltower/latest/userguide/region-how.html#mixed-governance)
// .
type Region struct {

	// The AWS Region name.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

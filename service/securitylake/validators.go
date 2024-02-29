// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateAwsLogSource struct {
}

func (*validateOpCreateAwsLogSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAwsLogSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAwsLogSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAwsLogSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateCustomLogSource struct {
}

func (*validateOpCreateCustomLogSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCustomLogSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCustomLogSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCustomLogSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataLakeExceptionSubscription struct {
}

func (*validateOpCreateDataLakeExceptionSubscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataLakeExceptionSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataLakeExceptionSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataLakeExceptionSubscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataLake struct {
}

func (*validateOpCreateDataLake) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataLake) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataLakeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataLakeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataLakeOrganizationConfiguration struct {
}

func (*validateOpCreateDataLakeOrganizationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataLakeOrganizationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataLakeOrganizationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataLakeOrganizationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSubscriber struct {
}

func (*validateOpCreateSubscriber) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSubscriber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSubscriberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSubscriberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSubscriberNotification struct {
}

func (*validateOpCreateSubscriberNotification) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSubscriberNotification) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSubscriberNotificationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSubscriberNotificationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAwsLogSource struct {
}

func (*validateOpDeleteAwsLogSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAwsLogSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAwsLogSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAwsLogSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCustomLogSource struct {
}

func (*validateOpDeleteCustomLogSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCustomLogSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCustomLogSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCustomLogSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDataLake struct {
}

func (*validateOpDeleteDataLake) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDataLake) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDataLakeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDataLakeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDataLakeOrganizationConfiguration struct {
}

func (*validateOpDeleteDataLakeOrganizationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDataLakeOrganizationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDataLakeOrganizationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDataLakeOrganizationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSubscriber struct {
}

func (*validateOpDeleteSubscriber) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSubscriber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSubscriberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSubscriberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSubscriberNotification struct {
}

func (*validateOpDeleteSubscriberNotification) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSubscriberNotification) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSubscriberNotificationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSubscriberNotificationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSubscriber struct {
}

func (*validateOpGetSubscriber) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSubscriber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSubscriberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSubscriberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterDataLakeDelegatedAdministrator struct {
}

func (*validateOpRegisterDataLakeDelegatedAdministrator) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterDataLakeDelegatedAdministrator) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterDataLakeDelegatedAdministratorInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterDataLakeDelegatedAdministratorInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDataLakeExceptionSubscription struct {
}

func (*validateOpUpdateDataLakeExceptionSubscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDataLakeExceptionSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDataLakeExceptionSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDataLakeExceptionSubscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDataLake struct {
}

func (*validateOpUpdateDataLake) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDataLake) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDataLakeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDataLakeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSubscriber struct {
}

func (*validateOpUpdateSubscriber) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSubscriber) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSubscriberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSubscriberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSubscriberNotification struct {
}

func (*validateOpUpdateSubscriberNotification) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSubscriberNotification) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSubscriberNotificationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSubscriberNotificationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateAwsLogSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAwsLogSource{}, middleware.After)
}

func addOpCreateCustomLogSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCustomLogSource{}, middleware.After)
}

func addOpCreateDataLakeExceptionSubscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataLakeExceptionSubscription{}, middleware.After)
}

func addOpCreateDataLakeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataLake{}, middleware.After)
}

func addOpCreateDataLakeOrganizationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataLakeOrganizationConfiguration{}, middleware.After)
}

func addOpCreateSubscriberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSubscriber{}, middleware.After)
}

func addOpCreateSubscriberNotificationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSubscriberNotification{}, middleware.After)
}

func addOpDeleteAwsLogSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAwsLogSource{}, middleware.After)
}

func addOpDeleteCustomLogSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCustomLogSource{}, middleware.After)
}

func addOpDeleteDataLakeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDataLake{}, middleware.After)
}

func addOpDeleteDataLakeOrganizationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDataLakeOrganizationConfiguration{}, middleware.After)
}

func addOpDeleteSubscriberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSubscriber{}, middleware.After)
}

func addOpDeleteSubscriberNotificationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSubscriberNotification{}, middleware.After)
}

func addOpGetSubscriberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSubscriber{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRegisterDataLakeDelegatedAdministratorValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterDataLakeDelegatedAdministrator{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateDataLakeExceptionSubscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDataLakeExceptionSubscription{}, middleware.After)
}

func addOpUpdateDataLakeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDataLake{}, middleware.After)
}

func addOpUpdateSubscriberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSubscriber{}, middleware.After)
}

func addOpUpdateSubscriberNotificationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSubscriberNotification{}, middleware.After)
}

func validateAwsIdentity(v *types.AwsIdentity) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AwsIdentity"}
	if v.Principal == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Principal"))
	}
	if v.ExternalId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExternalId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAwsLogSourceConfiguration(v *types.AwsLogSourceConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AwsLogSourceConfiguration"}
	if v.Regions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Regions"))
	}
	if len(v.SourceName) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("SourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateAwsLogSourceConfigurationList(v []types.AwsLogSourceConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AwsLogSourceConfigurationList"}
	for i := range v {
		if err := validateAwsLogSourceConfiguration(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCustomLogSourceConfiguration(v *types.CustomLogSourceConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CustomLogSourceConfiguration"}
	if v.CrawlerConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CrawlerConfiguration"))
	} else if v.CrawlerConfiguration != nil {
		if err := validateCustomLogSourceCrawlerConfiguration(v.CrawlerConfiguration); err != nil {
			invalidParams.AddNested("CrawlerConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ProviderIdentity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProviderIdentity"))
	} else if v.ProviderIdentity != nil {
		if err := validateAwsIdentity(v.ProviderIdentity); err != nil {
			invalidParams.AddNested("ProviderIdentity", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCustomLogSourceCrawlerConfiguration(v *types.CustomLogSourceCrawlerConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CustomLogSourceCrawlerConfiguration"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataLakeAutoEnableNewAccountConfiguration(v *types.DataLakeAutoEnableNewAccountConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataLakeAutoEnableNewAccountConfiguration"}
	if v.Region == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Region"))
	}
	if v.Sources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sources"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataLakeAutoEnableNewAccountConfigurationList(v []types.DataLakeAutoEnableNewAccountConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataLakeAutoEnableNewAccountConfigurationList"}
	for i := range v {
		if err := validateDataLakeAutoEnableNewAccountConfiguration(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataLakeConfiguration(v *types.DataLakeConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataLakeConfiguration"}
	if v.Region == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Region"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataLakeConfigurationList(v []types.DataLakeConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataLakeConfigurationList"}
	for i := range v {
		if err := validateDataLakeConfiguration(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHttpsNotificationConfiguration(v *types.HttpsNotificationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpsNotificationConfiguration"}
	if v.Endpoint == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Endpoint"))
	}
	if v.TargetRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetRoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNotificationConfiguration(v types.NotificationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotificationConfiguration"}
	switch uv := v.(type) {
	case *types.NotificationConfigurationMemberHttpsNotificationConfiguration:
		if err := validateHttpsNotificationConfiguration(&uv.Value); err != nil {
			invalidParams.AddNested("[httpsNotificationConfiguration]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAwsLogSourceInput(v *CreateAwsLogSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAwsLogSourceInput"}
	if v.Sources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sources"))
	} else if v.Sources != nil {
		if err := validateAwsLogSourceConfigurationList(v.Sources); err != nil {
			invalidParams.AddNested("Sources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCustomLogSourceInput(v *CreateCustomLogSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCustomLogSourceInput"}
	if v.SourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceName"))
	}
	if v.Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Configuration"))
	} else if v.Configuration != nil {
		if err := validateCustomLogSourceConfiguration(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataLakeExceptionSubscriptionInput(v *CreateDataLakeExceptionSubscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataLakeExceptionSubscriptionInput"}
	if v.SubscriptionProtocol == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriptionProtocol"))
	}
	if v.NotificationEndpoint == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationEndpoint"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataLakeInput(v *CreateDataLakeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataLakeInput"}
	if v.Configurations == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Configurations"))
	} else if v.Configurations != nil {
		if err := validateDataLakeConfigurationList(v.Configurations); err != nil {
			invalidParams.AddNested("Configurations", err.(smithy.InvalidParamsError))
		}
	}
	if v.MetaStoreManagerRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MetaStoreManagerRoleArn"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataLakeOrganizationConfigurationInput(v *CreateDataLakeOrganizationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataLakeOrganizationConfigurationInput"}
	if v.AutoEnableNewAccount != nil {
		if err := validateDataLakeAutoEnableNewAccountConfigurationList(v.AutoEnableNewAccount); err != nil {
			invalidParams.AddNested("AutoEnableNewAccount", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSubscriberInput(v *CreateSubscriberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSubscriberInput"}
	if v.SubscriberIdentity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberIdentity"))
	} else if v.SubscriberIdentity != nil {
		if err := validateAwsIdentity(v.SubscriberIdentity); err != nil {
			invalidParams.AddNested("SubscriberIdentity", err.(smithy.InvalidParamsError))
		}
	}
	if v.SubscriberName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberName"))
	}
	if v.Sources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sources"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSubscriberNotificationInput(v *CreateSubscriberNotificationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSubscriberNotificationInput"}
	if v.SubscriberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberId"))
	}
	if v.Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Configuration"))
	} else if v.Configuration != nil {
		if err := validateNotificationConfiguration(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAwsLogSourceInput(v *DeleteAwsLogSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAwsLogSourceInput"}
	if v.Sources == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sources"))
	} else if v.Sources != nil {
		if err := validateAwsLogSourceConfigurationList(v.Sources); err != nil {
			invalidParams.AddNested("Sources", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCustomLogSourceInput(v *DeleteCustomLogSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCustomLogSourceInput"}
	if v.SourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDataLakeInput(v *DeleteDataLakeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDataLakeInput"}
	if v.Regions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Regions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDataLakeOrganizationConfigurationInput(v *DeleteDataLakeOrganizationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDataLakeOrganizationConfigurationInput"}
	if v.AutoEnableNewAccount != nil {
		if err := validateDataLakeAutoEnableNewAccountConfigurationList(v.AutoEnableNewAccount); err != nil {
			invalidParams.AddNested("AutoEnableNewAccount", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSubscriberInput(v *DeleteSubscriberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSubscriberInput"}
	if v.SubscriberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSubscriberNotificationInput(v *DeleteSubscriberNotificationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSubscriberNotificationInput"}
	if v.SubscriberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSubscriberInput(v *GetSubscriberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSubscriberInput"}
	if v.SubscriberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterDataLakeDelegatedAdministratorInput(v *RegisterDataLakeDelegatedAdministratorInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterDataLakeDelegatedAdministratorInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDataLakeExceptionSubscriptionInput(v *UpdateDataLakeExceptionSubscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDataLakeExceptionSubscriptionInput"}
	if v.SubscriptionProtocol == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriptionProtocol"))
	}
	if v.NotificationEndpoint == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationEndpoint"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDataLakeInput(v *UpdateDataLakeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDataLakeInput"}
	if v.Configurations == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Configurations"))
	} else if v.Configurations != nil {
		if err := validateDataLakeConfigurationList(v.Configurations); err != nil {
			invalidParams.AddNested("Configurations", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSubscriberInput(v *UpdateSubscriberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSubscriberInput"}
	if v.SubscriberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberId"))
	}
	if v.SubscriberIdentity != nil {
		if err := validateAwsIdentity(v.SubscriberIdentity); err != nil {
			invalidParams.AddNested("SubscriberIdentity", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSubscriberNotificationInput(v *UpdateSubscriberNotificationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSubscriberNotificationInput"}
	if v.SubscriberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubscriberId"))
	}
	if v.Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Configuration"))
	} else if v.Configuration != nil {
		if err := validateNotificationConfiguration(v.Configuration); err != nil {
			invalidParams.AddNested("Configuration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

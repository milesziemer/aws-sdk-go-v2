// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a configuration for running a SageMaker image as a KernelGateway app.
// The configuration specifies the Amazon Elastic File System storage volume on the
// image, and a list of the kernels in the image.
func (c *Client) CreateAppImageConfig(ctx context.Context, params *CreateAppImageConfigInput, optFns ...func(*Options)) (*CreateAppImageConfigOutput, error) {
	if params == nil {
		params = &CreateAppImageConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppImageConfig", params, optFns, c.addOperationCreateAppImageConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppImageConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppImageConfigInput struct {

	// The name of the AppImageConfig. Must be unique to your account.
	//
	// This member is required.
	AppImageConfigName *string

	// The CodeEditorAppImageConfig . You can only specify one image kernel in the
	// AppImageConfig API. This kernel is shown to users before the image starts. After
	// the image runs, all kernels are visible in Code Editor.
	CodeEditorAppImageConfig *types.CodeEditorAppImageConfig

	// The JupyterLabAppImageConfig . You can only specify one image kernel in the
	// AppImageConfig API. This kernel is shown to users before the image starts. After
	// the image runs, all kernels are visible in JupyterLab.
	JupyterLabAppImageConfig *types.JupyterLabAppImageConfig

	// The KernelGatewayImageConfig. You can only specify one image kernel in the
	// AppImageConfig API. This kernel will be shown to users before the image starts.
	// Once the image runs, all kernels are visible in JupyterLab.
	KernelGatewayImageConfig *types.KernelGatewayImageConfig

	// A list of tags to apply to the AppImageConfig.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAppImageConfigOutput struct {

	// The ARN of the AppImageConfig.
	AppImageConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppImageConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAppImageConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAppImageConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAppImageConfig"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAppImageConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppImageConfig(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateAppImageConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAppImageConfig",
	}
}

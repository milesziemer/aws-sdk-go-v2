// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the audit history about a specific case if it exists.
func (c *Client) GetCaseAuditEvents(ctx context.Context, params *GetCaseAuditEventsInput, optFns ...func(*Options)) (*GetCaseAuditEventsOutput, error) {
	if params == nil {
		params = &GetCaseAuditEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCaseAuditEvents", params, optFns, c.addOperationGetCaseAuditEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCaseAuditEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCaseAuditEventsInput struct {

	// A unique identifier of the case.
	//
	// This member is required.
	CaseId *string

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// The maximum number of audit events to return. The current maximum supported
	// value is 25. This is also the default when no other value is provided.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetCaseAuditEventsOutput struct {

	// A list of case audits where each represents a particular edit of the case.
	//
	// This member is required.
	AuditEvents []*types.AuditEvent

	// The token for the next set of results. This is null if there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCaseAuditEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCaseAuditEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCaseAuditEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetCaseAuditEvents"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpGetCaseAuditEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCaseAuditEvents(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
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

// GetCaseAuditEventsAPIClient is a client that implements the GetCaseAuditEvents
// operation.
type GetCaseAuditEventsAPIClient interface {
	GetCaseAuditEvents(context.Context, *GetCaseAuditEventsInput, ...func(*Options)) (*GetCaseAuditEventsOutput, error)
}

var _ GetCaseAuditEventsAPIClient = (*Client)(nil)

// GetCaseAuditEventsPaginatorOptions is the paginator options for
// GetCaseAuditEvents
type GetCaseAuditEventsPaginatorOptions struct {
	// The maximum number of audit events to return. The current maximum supported
	// value is 25. This is also the default when no other value is provided.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCaseAuditEventsPaginator is a paginator for GetCaseAuditEvents
type GetCaseAuditEventsPaginator struct {
	options   GetCaseAuditEventsPaginatorOptions
	client    GetCaseAuditEventsAPIClient
	params    *GetCaseAuditEventsInput
	nextToken *string
	firstPage bool
}

// NewGetCaseAuditEventsPaginator returns a new GetCaseAuditEventsPaginator
func NewGetCaseAuditEventsPaginator(client GetCaseAuditEventsAPIClient, params *GetCaseAuditEventsInput, optFns ...func(*GetCaseAuditEventsPaginatorOptions)) *GetCaseAuditEventsPaginator {
	if params == nil {
		params = &GetCaseAuditEventsInput{}
	}

	options := GetCaseAuditEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCaseAuditEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCaseAuditEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetCaseAuditEvents page.
func (p *GetCaseAuditEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCaseAuditEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetCaseAuditEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetCaseAuditEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetCaseAuditEvents",
	}
}

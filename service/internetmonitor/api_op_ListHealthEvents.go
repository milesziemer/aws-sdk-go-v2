// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all health events for a monitor in Amazon CloudWatch Internet Monitor.
// Returns information for health events including the event start and end time and
// the status. Health events that have start times during the time frame that is
// requested are not included in the list of health events.
func (c *Client) ListHealthEvents(ctx context.Context, params *ListHealthEventsInput, optFns ...func(*Options)) (*ListHealthEventsOutput, error) {
	if params == nil {
		params = &ListHealthEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHealthEvents", params, optFns, c.addOperationListHealthEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHealthEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHealthEventsInput struct {

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// The time when a health event ended. If the health event is still ongoing, then
	// the end time is not set.
	EndTime *time.Time

	// The status of a health event.
	EventStatus types.HealthEventStatus

	// TBD
	LinkedAccountId *string

	// The number of health event objects that you want to return with this call.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// The time when a health event started.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type ListHealthEventsOutput struct {

	// A list of health events.
	//
	// This member is required.
	HealthEvents []types.HealthEvent

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHealthEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHealthEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHealthEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListHealthEvents"); err != nil {
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
	if err = addOpListHealthEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHealthEvents(options.Region), middleware.Before); err != nil {
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

// ListHealthEventsAPIClient is a client that implements the ListHealthEvents
// operation.
type ListHealthEventsAPIClient interface {
	ListHealthEvents(context.Context, *ListHealthEventsInput, ...func(*Options)) (*ListHealthEventsOutput, error)
}

var _ ListHealthEventsAPIClient = (*Client)(nil)

// ListHealthEventsPaginatorOptions is the paginator options for ListHealthEvents
type ListHealthEventsPaginatorOptions struct {
	// The number of health event objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHealthEventsPaginator is a paginator for ListHealthEvents
type ListHealthEventsPaginator struct {
	options   ListHealthEventsPaginatorOptions
	client    ListHealthEventsAPIClient
	params    *ListHealthEventsInput
	nextToken *string
	firstPage bool
}

// NewListHealthEventsPaginator returns a new ListHealthEventsPaginator
func NewListHealthEventsPaginator(client ListHealthEventsAPIClient, params *ListHealthEventsInput, optFns ...func(*ListHealthEventsPaginatorOptions)) *ListHealthEventsPaginator {
	if params == nil {
		params = &ListHealthEventsInput{}
	}

	options := ListHealthEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHealthEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHealthEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHealthEvents page.
func (p *ListHealthEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHealthEventsOutput, error) {
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

	result, err := p.client.ListHealthEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHealthEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListHealthEvents",
	}
}

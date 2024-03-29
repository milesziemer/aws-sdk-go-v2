// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog/document"
	smithydocument "github.com/aws/smithy-go/document"
)

// Object that allows filtering on entity id of an AMI product.
type AmiProductEntityIdFilter struct {

	// A string array of unique entity id values to be filtered on.
	ValueList []string

	noSmithyDocumentSerde
}

// Object containing all the filter fields for AMI products. Client can add only
// one wildcard filter and a maximum of 8 filters in a single ListEntities request.
type AmiProductFilters struct {

	// Unique identifier for the AMI product.
	EntityId *AmiProductEntityIdFilter

	// The last date on which the AMI product was modified.
	LastModifiedDate *AmiProductLastModifiedDateFilter

	// The title of the AMI product.
	ProductTitle *AmiProductTitleFilter

	// The visibility of the AMI product.
	Visibility *AmiProductVisibilityFilter

	noSmithyDocumentSerde
}

// Object that allows filtering based on the last modified date of AMI products.
type AmiProductLastModifiedDateFilter struct {

	// Dates between which the AMI product was last modified.
	DateRange *AmiProductLastModifiedDateFilterDateRange

	noSmithyDocumentSerde
}

// Object that contains date range of the last modified date to be filtered on.
// You can optionally provide a BeforeValue and/or AfterValue . Both are inclusive.
type AmiProductLastModifiedDateFilterDateRange struct {

	// Date after which the AMI product was last modified.
	AfterValue *string

	// Date before which the AMI product was last modified.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Objects that allows sorting on AMI products based on certain fields and sorting
// order.
type AmiProductSort struct {

	// Field to sort the AMI products by.
	SortBy AmiProductSortBy

	// The sorting order. Can be ASCENDING or DESCENDING . The default value is
	// DESCENDING .
	SortOrder SortOrder

	noSmithyDocumentSerde
}

// Object that contains summarized information about an AMI product.
type AmiProductSummary struct {

	// The title of the AMI product.
	ProductTitle *string

	// The lifecycle of the AMI product.
	Visibility AmiProductVisibilityString

	noSmithyDocumentSerde
}

// Object that allows filtering on product title.
type AmiProductTitleFilter struct {

	// A string array of unique product title values to be filtered on.
	ValueList []string

	// A string that will be the wildCard input for product tile filter. It matches
	// the provided value as a substring in the actual value.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Object that allows filtering on the visibility of the product in the AWS
// Marketplace.
type AmiProductVisibilityFilter struct {

	// A string array of unique visibility values to be filtered on.
	ValueList []AmiProductVisibilityString

	noSmithyDocumentSerde
}

// An object that contains an error code and error message.
type BatchDescribeErrorDetail struct {

	// The error code returned.
	ErrorCode *string

	// The error message returned.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// An object that contains the ChangeType , Details , and Entity .
type Change struct {

	// Change types are single string values that describe your intention for the
	// change. Each change type is unique for each EntityType provided in the change's
	// scope. For more information about change types available for single-AMI
	// products, see Working with single-AMI products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/ami-products.html#working-with-single-AMI-products)
	// . Also, for more information about change types available for container-based
	// products, see Working with container products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/container-products.html#working-with-container-products)
	// .
	//
	// This member is required.
	ChangeType *string

	// The entity to be changed.
	//
	// This member is required.
	Entity *Entity

	// Optional name for the change.
	ChangeName *string

	// This object contains details specific to the change type of the requested
	// change. For more information about change types available for single-AMI
	// products, see Working with single-AMI products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/ami-products.html#working-with-single-AMI-products)
	// . Also, for more information about change types available for container-based
	// products, see Working with container products (https://docs.aws.amazon.com/marketplace-catalog/latest/api-reference/container-products.html#working-with-container-products)
	// .
	Details *string

	// Alternative field that accepts a JSON value instead of a string for ChangeType
	// details. You can use either Details or DetailsDocument , but not both.
	DetailsDocument document.Interface

	// The tags associated with the change.
	EntityTags []Tag

	noSmithyDocumentSerde
}

// A summary of a change set returned in a list of change sets when the
// ListChangeSets action is called.
type ChangeSetSummaryListItem struct {

	// The ARN associated with the unique identifier for the change set referenced in
	// this request.
	ChangeSetArn *string

	// The unique identifier for a change set.
	ChangeSetId *string

	// The non-unique name for the change set.
	ChangeSetName *string

	// The time, in ISO 8601 format (2018-02-27T13:45:22Z), when the change set was
	// finished.
	EndTime *string

	// This object is a list of entity IDs (string) that are a part of a change set.
	// The entity ID list is a maximum of 20 entities. It must contain at least one
	// entity.
	EntityIdList []string

	// Returned if the change set is in FAILED status. Can be either CLIENT_ERROR ,
	// which means that there are issues with the request (see the ErrorDetailList of
	// DescribeChangeSet ), or SERVER_FAULT , which means that there is a problem in
	// the system, and you should retry your request.
	FailureCode FailureCode

	// The time, in ISO 8601 format (2018-02-27T13:45:22Z), when the change set was
	// started.
	StartTime *string

	// The current status of the change set.
	Status ChangeStatus

	noSmithyDocumentSerde
}

// This object is a container for common summary information about the change. The
// summary doesn't contain the whole change structure.
type ChangeSummary struct {

	// Optional name for the change.
	ChangeName *string

	// The type of the change.
	ChangeType *string

	// This object contains details specific to the change type of the requested
	// change.
	Details *string

	// The JSON value of the details specific to the change type of the requested
	// change.
	DetailsDocument document.Interface

	// The entity to be changed.
	Entity *Entity

	// An array of ErrorDetail objects associated with the change.
	ErrorDetailList []ErrorDetail

	noSmithyDocumentSerde
}

// Object that allows filtering on entity id of a container product.
type ContainerProductEntityIdFilter struct {

	// A string array of unique entity id values to be filtered on.
	ValueList []string

	noSmithyDocumentSerde
}

// Object containing all the filter fields for container products. Client can add
// only one wildcard filter and a maximum of 8 filters in a single ListEntities
// request.
type ContainerProductFilters struct {

	// Unique identifier for the container product.
	EntityId *ContainerProductEntityIdFilter

	// The last date on which the container product was modified.
	LastModifiedDate *ContainerProductLastModifiedDateFilter

	// The title of the container product.
	ProductTitle *ContainerProductTitleFilter

	// The visibility of the container product.
	Visibility *ContainerProductVisibilityFilter

	noSmithyDocumentSerde
}

// Object that allows filtering based on the last modified date of container
// products.
type ContainerProductLastModifiedDateFilter struct {

	// Dates between which the container product was last modified.
	DateRange *ContainerProductLastModifiedDateFilterDateRange

	noSmithyDocumentSerde
}

// Object that contains date range of the last modified date to be filtered on.
// You can optionally provide a BeforeValue and/or AfterValue . Both are inclusive.
type ContainerProductLastModifiedDateFilterDateRange struct {

	// Date after which the container product was last modified.
	AfterValue *string

	// Date before which the container product was last modified.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Objects that allows sorting on container products based on certain fields and
// sorting order.
type ContainerProductSort struct {

	// Field to sort the container products by.
	SortBy ContainerProductSortBy

	// The sorting order. Can be ASCENDING or DESCENDING . The default value is
	// DESCENDING .
	SortOrder SortOrder

	noSmithyDocumentSerde
}

// Object that contains summarized information about a container product.
type ContainerProductSummary struct {

	// The title of the container product.
	ProductTitle *string

	// The lifecycle of the product.
	Visibility ContainerProductVisibilityString

	noSmithyDocumentSerde
}

// Object that allows filtering on product title.
type ContainerProductTitleFilter struct {

	// A string array of unique product title values to be filtered on.
	ValueList []string

	// A string that will be the wildCard input for product tile filter. It matches
	// the provided value as a substring in the actual value.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Object that allows filtering on the visibility of the product in the AWS
// Marketplace.
type ContainerProductVisibilityFilter struct {

	// A string array of unique visibility values to be filtered on.
	ValueList []ContainerProductVisibilityString

	noSmithyDocumentSerde
}

// Object that allows filtering on entity id of a data product.
type DataProductEntityIdFilter struct {

	// A string array of unique entity id values to be filtered on.
	ValueList []string

	noSmithyDocumentSerde
}

// Object containing all the filter fields for data products. Client can add only
// one wildcard filter and a maximum of 8 filters in a single ListEntities request.
type DataProductFilters struct {

	// Unique identifier for the data product.
	EntityId *DataProductEntityIdFilter

	// The last date on which the data product was modified.
	LastModifiedDate *DataProductLastModifiedDateFilter

	// The title of the data product.
	ProductTitle *DataProductTitleFilter

	// The visibility of the data product.
	Visibility *DataProductVisibilityFilter

	noSmithyDocumentSerde
}

// Object that allows filtering based on the last modified date of data products.
type DataProductLastModifiedDateFilter struct {

	// Dates between which the data product was last modified.
	DateRange *DataProductLastModifiedDateFilterDateRange

	noSmithyDocumentSerde
}

// Object that contains date range of the last modified date to be filtered on.
// You can optionally provide a BeforeValue and/or AfterValue . Both are inclusive.
type DataProductLastModifiedDateFilterDateRange struct {

	// Date after which the data product was last modified.
	AfterValue *string

	// Date before which the data product was last modified.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Objects that allows sorting on data products based on certain fields and
// sorting order.
type DataProductSort struct {

	// Field to sort the data products by.
	SortBy DataProductSortBy

	// The sorting order. Can be ASCENDING or DESCENDING . The default value is
	// DESCENDING .
	SortOrder SortOrder

	noSmithyDocumentSerde
}

// Object that contains summarized information about a data product.
type DataProductSummary struct {

	// The title of the data product.
	ProductTitle *string

	// The lifecycle of the data product.
	Visibility DataProductVisibilityString

	noSmithyDocumentSerde
}

// Object that allows filtering on product title.
type DataProductTitleFilter struct {

	// A string array of unique product title values to be filtered on.
	ValueList []string

	// A string that will be the wildCard input for product tile filter. It matches
	// the provided value as a substring in the actual value.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Object that allows filtering on the visibility of the product in the AWS
// Marketplace.
type DataProductVisibilityFilter struct {

	// A string array of unique visibility values to be filtered on.
	ValueList []DataProductVisibilityString

	noSmithyDocumentSerde
}

// An entity contains data that describes your product, its supported features,
// and how it can be used or launched by your customer.
type Entity struct {

	// The type of entity.
	//
	// This member is required.
	Type *string

	// The identifier for the entity.
	Identifier *string

	noSmithyDocumentSerde
}

// An object that contains metadata and details about the entity.
type EntityDetail struct {

	// An object that contains all the details of the entity.
	DetailsDocument document.Interface

	// The Amazon Resource Name (ARN) of the entity.
	EntityArn *string

	// The ID of the entity, in the format of EntityId@RevisionId .
	EntityIdentifier *string

	// The entity type of the entity, in the format of EntityType@Version .
	EntityType *string

	// The last time the entity was modified.
	LastModifiedDate *string

	noSmithyDocumentSerde
}

// An object that contains entity ID and the catalog in which the entity is
// present.
type EntityRequest struct {

	// The name of the catalog the entity is present in. The only value at this time
	// is AWSMarketplace .
	//
	// This member is required.
	Catalog *string

	// The ID of the entity.
	//
	// This member is required.
	EntityId *string

	noSmithyDocumentSerde
}

// This object is a container for common summary information about the entity. The
// summary doesn't contain the whole entity structure, but it does contain
// information common across all entities.
type EntitySummary struct {

	// An object that contains summary information about the AMI product.
	AmiProductSummary *AmiProductSummary

	// An object that contains summary information about the container product.
	ContainerProductSummary *ContainerProductSummary

	// An object that contains summary information about the data product.
	DataProductSummary *DataProductSummary

	// The ARN associated with the unique identifier for the entity.
	EntityArn *string

	// The unique identifier for the entity.
	EntityId *string

	// The type of the entity.
	EntityType *string

	// The last time the entity was published, using ISO 8601 format
	// (2018-02-27T13:45:22Z).
	LastModifiedDate *string

	// The name for the entity. This value is not unique. It is defined by the seller.
	Name *string

	// An object that contains summary information about the offer.
	OfferSummary *OfferSummary

	// An object that contains summary information about the Resale Authorization.
	ResaleAuthorizationSummary *ResaleAuthorizationSummary

	// An object that contains summary information about the SaaS product.
	SaaSProductSummary *SaaSProductSummary

	// The visibility status of the entity to buyers. This value can be Public
	// (everyone can view the entity), Limited (the entity is visible to limited
	// accounts only), or Restricted (the entity was published and then unpublished
	// and only existing buyers can view it).
	Visibility *string

	noSmithyDocumentSerde
}

// Object containing all the filter fields per entity type.
//
// The following types satisfy this interface:
//
//	EntityTypeFiltersMemberAmiProductFilters
//	EntityTypeFiltersMemberContainerProductFilters
//	EntityTypeFiltersMemberDataProductFilters
//	EntityTypeFiltersMemberOfferFilters
//	EntityTypeFiltersMemberResaleAuthorizationFilters
//	EntityTypeFiltersMemberSaaSProductFilters
type EntityTypeFilters interface {
	isEntityTypeFilters()
}

// A filter for AMI products.
type EntityTypeFiltersMemberAmiProductFilters struct {
	Value AmiProductFilters

	noSmithyDocumentSerde
}

func (*EntityTypeFiltersMemberAmiProductFilters) isEntityTypeFilters() {}

// A filter for container products.
type EntityTypeFiltersMemberContainerProductFilters struct {
	Value ContainerProductFilters

	noSmithyDocumentSerde
}

func (*EntityTypeFiltersMemberContainerProductFilters) isEntityTypeFilters() {}

// A filter for data products.
type EntityTypeFiltersMemberDataProductFilters struct {
	Value DataProductFilters

	noSmithyDocumentSerde
}

func (*EntityTypeFiltersMemberDataProductFilters) isEntityTypeFilters() {}

// A filter for offers.
type EntityTypeFiltersMemberOfferFilters struct {
	Value OfferFilters

	noSmithyDocumentSerde
}

func (*EntityTypeFiltersMemberOfferFilters) isEntityTypeFilters() {}

// A filter for Resale Authorizations.
type EntityTypeFiltersMemberResaleAuthorizationFilters struct {
	Value ResaleAuthorizationFilters

	noSmithyDocumentSerde
}

func (*EntityTypeFiltersMemberResaleAuthorizationFilters) isEntityTypeFilters() {}

// A filter for SaaS products.
type EntityTypeFiltersMemberSaaSProductFilters struct {
	Value SaaSProductFilters

	noSmithyDocumentSerde
}

func (*EntityTypeFiltersMemberSaaSProductFilters) isEntityTypeFilters() {}

// Object containing all the sort fields per entity type.
//
// The following types satisfy this interface:
//
//	EntityTypeSortMemberAmiProductSort
//	EntityTypeSortMemberContainerProductSort
//	EntityTypeSortMemberDataProductSort
//	EntityTypeSortMemberOfferSort
//	EntityTypeSortMemberResaleAuthorizationSort
//	EntityTypeSortMemberSaaSProductSort
type EntityTypeSort interface {
	isEntityTypeSort()
}

// A sort for AMI products.
type EntityTypeSortMemberAmiProductSort struct {
	Value AmiProductSort

	noSmithyDocumentSerde
}

func (*EntityTypeSortMemberAmiProductSort) isEntityTypeSort() {}

// A sort for container products.
type EntityTypeSortMemberContainerProductSort struct {
	Value ContainerProductSort

	noSmithyDocumentSerde
}

func (*EntityTypeSortMemberContainerProductSort) isEntityTypeSort() {}

// A sort for data products.
type EntityTypeSortMemberDataProductSort struct {
	Value DataProductSort

	noSmithyDocumentSerde
}

func (*EntityTypeSortMemberDataProductSort) isEntityTypeSort() {}

// A sort for offers.
type EntityTypeSortMemberOfferSort struct {
	Value OfferSort

	noSmithyDocumentSerde
}

func (*EntityTypeSortMemberOfferSort) isEntityTypeSort() {}

// A sort for Resale Authorizations.
type EntityTypeSortMemberResaleAuthorizationSort struct {
	Value ResaleAuthorizationSort

	noSmithyDocumentSerde
}

func (*EntityTypeSortMemberResaleAuthorizationSort) isEntityTypeSort() {}

// A sort for SaaS products.
type EntityTypeSortMemberSaaSProductSort struct {
	Value SaaSProductSort

	noSmithyDocumentSerde
}

func (*EntityTypeSortMemberSaaSProductSort) isEntityTypeSort() {}

// Details about the error.
type ErrorDetail struct {

	// The error code that identifies the type of error.
	ErrorCode *string

	// The message for the error.
	ErrorMessage *string

	noSmithyDocumentSerde
}

// A filter object, used to optionally filter results from calls to the
// ListEntities and ListChangeSets actions.
type Filter struct {

	// For ListEntities , the supported value for this is an EntityId . For
	// ListChangeSets , the supported values are as follows:
	Name *string

	// ListEntities - This is a list of unique EntityId s. ListChangeSets - The
	// supported filter names and associated ValueList s is as follows:
	//   - ChangeSetName - The supported ValueList is a list of non-unique
	//   ChangeSetName s. These are defined when you call the StartChangeSet action.
	//   - Status - The supported ValueList is a list of statuses for all change set
	//   requests.
	//   - EntityId - The supported ValueList is a list of unique EntityId s.
	//   - BeforeStartTime - The supported ValueList is a list of all change sets that
	//   started before the filter value.
	//   - AfterStartTime - The supported ValueList is a list of all change sets that
	//   started after the filter value.
	//   - BeforeEndTime - The supported ValueList is a list of all change sets that
	//   ended before the filter value.
	//   - AfterEndTime - The supported ValueList is a list of all change sets that
	//   ended after the filter value.
	ValueList []string

	noSmithyDocumentSerde
}

// Allows filtering on the AvailabilityEndDate of an offer.
type OfferAvailabilityEndDateFilter struct {

	// Allows filtering on the AvailabilityEndDate of an offer with date range as
	// input.
	DateRange *OfferAvailabilityEndDateFilterDateRange

	noSmithyDocumentSerde
}

// Allows filtering on the AvailabilityEndDate of an offer with date range as
// input.
type OfferAvailabilityEndDateFilterDateRange struct {

	// Allows filtering on the AvailabilityEndDate of an offer after a date.
	AfterValue *string

	// Allows filtering on the AvailabilityEndDate of an offer before a date.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the BuyerAccounts of an offer.
type OfferBuyerAccountsFilter struct {

	// Allows filtering on the BuyerAccounts of an offer with wild card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the entity id of an offer.
type OfferEntityIdFilter struct {

	// Allows filtering on entity id of an offer with list input.
	ValueList []string

	noSmithyDocumentSerde
}

// Object containing all the filter fields for offers entity. Client can add only
// one wildcard filter and a maximum of 8 filters in a single ListEntities request.
type OfferFilters struct {

	// Allows filtering on the AvailabilityEndDate of an offer.
	AvailabilityEndDate *OfferAvailabilityEndDateFilter

	// Allows filtering on the BuyerAccounts of an offer.
	BuyerAccounts *OfferBuyerAccountsFilter

	// Allows filtering on EntityId of an offer.
	EntityId *OfferEntityIdFilter

	// Allows filtering on the LastModifiedDate of an offer.
	LastModifiedDate *OfferLastModifiedDateFilter

	// Allows filtering on the Name of an offer.
	Name *OfferNameFilter

	// Allows filtering on the ProductId of an offer.
	ProductId *OfferProductIdFilter

	// Allows filtering on the ReleaseDate of an offer.
	ReleaseDate *OfferReleaseDateFilter

	// Allows filtering on the ResaleAuthorizationId of an offer. Not all offers have
	// a ResaleAuthorizationId . The response will only include offers for which you
	// have permissions.
	ResaleAuthorizationId *OfferResaleAuthorizationIdFilter

	// Allows filtering on the State of an offer.
	State *OfferStateFilter

	// Allows filtering on the Targeting of an offer.
	Targeting *OfferTargetingFilter

	noSmithyDocumentSerde
}

// Allows filtering on the LastModifiedDate of an offer.
type OfferLastModifiedDateFilter struct {

	// Allows filtering on the LastModifiedDate of an offer with date range as input.
	DateRange *OfferLastModifiedDateFilterDateRange

	noSmithyDocumentSerde
}

// Allows filtering on the LastModifiedDate of an offer with date range as input.
type OfferLastModifiedDateFilterDateRange struct {

	// Allows filtering on the LastModifiedDate of an offer after a date.
	AfterValue *string

	// Allows filtering on the LastModifiedDate of an offer before a date.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the Name of an offer.
type OfferNameFilter struct {

	// Allows filtering on the Name of an offer with list input.
	ValueList []string

	// Allows filtering on the Name of an offer with wild card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the ProductId of an offer.
type OfferProductIdFilter struct {

	// Allows filtering on the ProductId of an offer with list input.
	ValueList []string

	noSmithyDocumentSerde
}

// Allows filtering on the ReleaseDate of an offer.
type OfferReleaseDateFilter struct {

	// Allows filtering on the ReleaseDate of an offer with date range as input.
	DateRange *OfferReleaseDateFilterDateRange

	noSmithyDocumentSerde
}

// Allows filtering on the ReleaseDate of an offer with date range as input.
type OfferReleaseDateFilterDateRange struct {

	// Allows filtering on the ReleaseDate of offers after a date.
	AfterValue *string

	// Allows filtering on the ReleaseDate of offers before a date.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the ResaleAuthorizationId of an offer. Not all offers have
// a ResaleAuthorizationId . The response will only include offers for which you
// have permissions.
type OfferResaleAuthorizationIdFilter struct {

	// Allows filtering on the ResaleAuthorizationId of an offer with list input.
	ValueList []string

	noSmithyDocumentSerde
}

// Allows to sort offers.
type OfferSort struct {

	// Allows to sort offers.
	SortBy OfferSortBy

	// Allows to sort offers.
	SortOrder SortOrder

	noSmithyDocumentSerde
}

// Allows filtering on the State of an offer.
type OfferStateFilter struct {

	// Allows filtering on the State of an offer with list input.
	ValueList []OfferStateString

	noSmithyDocumentSerde
}

// Summarized information about an offer.
type OfferSummary struct {

	// The availability end date of the offer.
	AvailabilityEndDate *string

	// The buyer accounts in the offer.
	BuyerAccounts []string

	// The name of the offer.
	Name *string

	// The product ID of the offer.
	ProductId *string

	// The release date of the offer.
	ReleaseDate *string

	// The ResaleAuthorizationId of the offer.
	ResaleAuthorizationId *string

	// The status of the offer.
	State OfferStateString

	// The targeting in the offer.
	Targeting []OfferTargetingString

	noSmithyDocumentSerde
}

// Allows filtering on the Targeting of an offer.
type OfferTargetingFilter struct {

	// Allows filtering on the Targeting of an offer with list input.
	ValueList []OfferTargetingString

	noSmithyDocumentSerde
}

// Allows filtering on AvailabilityEndDate of a ResaleAuthorization.
type ResaleAuthorizationAvailabilityEndDateFilter struct {

	// Allows filtering on AvailabilityEndDate of a ResaleAuthorization with date
	// range as input
	DateRange *ResaleAuthorizationAvailabilityEndDateFilterDateRange

	// Allows filtering on AvailabilityEndDate of a ResaleAuthorization with date
	// value as input.
	ValueList []string

	noSmithyDocumentSerde
}

// Allows filtering on AvailabilityEndDate of a ResaleAuthorization with date
// range as input.
type ResaleAuthorizationAvailabilityEndDateFilterDateRange struct {

	// Allows filtering on AvailabilityEndDate of a ResaleAuthorization after a date.
	AfterValue *string

	// Allows filtering on AvailabilityEndDate of a ResaleAuthorization before a date.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Allows filtering on CreatedDate of a ResaleAuthorization.
type ResaleAuthorizationCreatedDateFilter struct {

	// Allows filtering on CreatedDate of a ResaleAuthorization with date range as
	// input.
	DateRange *ResaleAuthorizationCreatedDateFilterDateRange

	// Allows filtering on CreatedDate of a ResaleAuthorization with date value as
	// input.
	ValueList []string

	noSmithyDocumentSerde
}

// Allows filtering on CreatedDate of a ResaleAuthorization with date range as
// input.
type ResaleAuthorizationCreatedDateFilterDateRange struct {

	// Allows filtering on CreatedDate of a ResaleAuthorization after a date.
	AfterValue *string

	// Allows filtering on CreatedDate of a ResaleAuthorization before a date.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Allows filtering on EntityId of a ResaleAuthorization.
type ResaleAuthorizationEntityIdFilter struct {

	// Allows filtering on EntityId of a ResaleAuthorization with list input.
	ValueList []string

	noSmithyDocumentSerde
}

// Object containing all the filter fields for resale authorization entity. Client
// can add only one wildcard filter and a maximum of 8 filters in a single
// ListEntities request.
type ResaleAuthorizationFilters struct {

	// Allows filtering on the AvailabilityEndDate of a ResaleAuthorization.
	AvailabilityEndDate *ResaleAuthorizationAvailabilityEndDateFilter

	// Allows filtering on the CreatedDate of a ResaleAuthorization.
	CreatedDate *ResaleAuthorizationCreatedDateFilter

	// Allows filtering on the EntityId of a ResaleAuthorization.
	EntityId *ResaleAuthorizationEntityIdFilter

	// Allows filtering on the LastModifiedDate of a ResaleAuthorization.
	LastModifiedDate *ResaleAuthorizationLastModifiedDateFilter

	// Allows filtering on the ManufacturerAccountId of a ResaleAuthorization.
	ManufacturerAccountId *ResaleAuthorizationManufacturerAccountIdFilter

	// Allows filtering on the ManufacturerLegalName of a ResaleAuthorization.
	ManufacturerLegalName *ResaleAuthorizationManufacturerLegalNameFilter

	// Allows filtering on the Name of a ResaleAuthorization.
	Name *ResaleAuthorizationNameFilter

	// Allows filtering on the OfferExtendedStatus of a ResaleAuthorization.
	OfferExtendedStatus *ResaleAuthorizationOfferExtendedStatusFilter

	// Allows filtering on the ProductId of a ResaleAuthorization.
	ProductId *ResaleAuthorizationProductIdFilter

	// Allows filtering on the ProductName of a ResaleAuthorization.
	ProductName *ResaleAuthorizationProductNameFilter

	// Allows filtering on the ResellerAccountID of a ResaleAuthorization.
	ResellerAccountID *ResaleAuthorizationResellerAccountIDFilter

	// Allows filtering on the ResellerLegalName of a ResaleAuthorization.
	ResellerLegalName *ResaleAuthorizationResellerLegalNameFilter

	// Allows filtering on the Status of a ResaleAuthorization.
	Status *ResaleAuthorizationStatusFilter

	noSmithyDocumentSerde
}

// Allows filtering on the LastModifiedDate of a ResaleAuthorization.
type ResaleAuthorizationLastModifiedDateFilter struct {

	// Allows filtering on the LastModifiedDate of a ResaleAuthorization with date
	// range as input.
	DateRange *ResaleAuthorizationLastModifiedDateFilterDateRange

	noSmithyDocumentSerde
}

// Allows filtering on the LastModifiedDate of a ResaleAuthorization with date
// range as input.
type ResaleAuthorizationLastModifiedDateFilterDateRange struct {

	// Allows filtering on the LastModifiedDate of a ResaleAuthorization after a date.
	AfterValue *string

	// Allows filtering on the LastModifiedDate of a ResaleAuthorization before a date.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the ManufacturerAccountId of a ResaleAuthorization.
type ResaleAuthorizationManufacturerAccountIdFilter struct {

	// Allows filtering on the ManufacturerAccountId of a ResaleAuthorization with
	// list input.
	ValueList []string

	// Allows filtering on the ManufacturerAccountId of a ResaleAuthorization with
	// wild card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the ManufacturerLegalName of a ResaleAuthorization.
type ResaleAuthorizationManufacturerLegalNameFilter struct {

	// Allows filtering on the ManufacturerLegalName of a ResaleAuthorization with
	// list input.
	ValueList []string

	// Allows filtering on the ManufacturerLegalName of a ResaleAuthorization with
	// wild card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the Name of a ResaleAuthorization.
type ResaleAuthorizationNameFilter struct {

	// Allows filtering on the Name of a ResaleAuthorization with list input.
	ValueList []string

	// Allows filtering on the Name of a ResaleAuthorization with wild card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the OfferExtendedStatus of a ResaleAuthorization.
type ResaleAuthorizationOfferExtendedStatusFilter struct {

	// Allows filtering on the OfferExtendedStatus of a ResaleAuthorization with list
	// input.
	ValueList []string

	noSmithyDocumentSerde
}

// Allows filtering on the ProductId of a ResaleAuthorization.
type ResaleAuthorizationProductIdFilter struct {

	// Allows filtering on the ProductId of a ResaleAuthorization with list input.
	ValueList []string

	// Allows filtering on the ProductId of a ResaleAuthorization with wild card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the ProductName of a ResaleAuthorization.
type ResaleAuthorizationProductNameFilter struct {

	// Allows filtering on the ProductName of a ResaleAuthorization with list input.
	ValueList []string

	// Allows filtering on the ProductName of a ResaleAuthorization with wild card
	// input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the ResellerAccountID of a ResaleAuthorization.
type ResaleAuthorizationResellerAccountIDFilter struct {

	// Allows filtering on the ResellerAccountID of a ResaleAuthorization with list
	// input.
	ValueList []string

	// Allows filtering on the ResellerAccountID of a ResaleAuthorization with wild
	// card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows filtering on the ResellerLegalName of a ResaleAuthorization.
type ResaleAuthorizationResellerLegalNameFilter struct {

	// Allows filtering on the ResellerLegalNameProductName of a ResaleAuthorization
	// with list input.
	ValueList []string

	// Allows filtering on the ResellerLegalName of a ResaleAuthorization with wild
	// card input.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Allows to sort ResaleAuthorization.
type ResaleAuthorizationSort struct {

	// Allows to sort ResaleAuthorization.
	SortBy ResaleAuthorizationSortBy

	// Allows to sort ResaleAuthorization.
	SortOrder SortOrder

	noSmithyDocumentSerde
}

// Allows filtering on the Status of a ResaleAuthorization.
type ResaleAuthorizationStatusFilter struct {

	// Allows filtering on the Status of a ResaleAuthorization with list input.
	ValueList []ResaleAuthorizationStatusString

	noSmithyDocumentSerde
}

// Summarized information about a Resale Authorization.
type ResaleAuthorizationSummary struct {

	// The availability end date of the ResaleAuthorization.
	AvailabilityEndDate *string

	// The created date of the ResaleAuthorization.
	CreatedDate *string

	// The manufacturer account ID of the ResaleAuthorization.
	ManufacturerAccountId *string

	// The manufacturer legal name of the ResaleAuthorization.
	ManufacturerLegalName *string

	// The name of the ResaleAuthorization.
	Name *string

	// The offer extended status of the ResaleAuthorization
	OfferExtendedStatus *string

	// The product ID of the ResaleAuthorization.
	ProductId *string

	// The product name of the ResaleAuthorization.
	ProductName *string

	// The reseller account ID of the ResaleAuthorization.
	ResellerAccountID *string

	// The reseller legal name of the ResaleAuthorization
	ResellerLegalName *string

	// The status of the ResaleAuthorization.
	Status ResaleAuthorizationStatusString

	noSmithyDocumentSerde
}

// Object that allows filtering on entity id of a SaaS product.
type SaaSProductEntityIdFilter struct {

	// A string array of unique entity id values to be filtered on.
	ValueList []string

	noSmithyDocumentSerde
}

// Object containing all the filter fields for SaaS products. Client can add only
// one wildcard filter and a maximum of 8 filters in a single ListEntities request.
type SaaSProductFilters struct {

	// Unique identifier for the SaaS product.
	EntityId *SaaSProductEntityIdFilter

	// The last date on which the SaaS product was modified.
	LastModifiedDate *SaaSProductLastModifiedDateFilter

	// The title of the SaaS product.
	ProductTitle *SaaSProductTitleFilter

	// The visibility of the SaaS product.
	Visibility *SaaSProductVisibilityFilter

	noSmithyDocumentSerde
}

// Object that allows filtering based on the last modified date of SaaS products
type SaaSProductLastModifiedDateFilter struct {

	// Dates between which the SaaS product was last modified.
	DateRange *SaaSProductLastModifiedDateFilterDateRange

	noSmithyDocumentSerde
}

// Object that contains date range of the last modified date to be filtered on.
// You can optionally provide a BeforeValue and/or AfterValue . Both are inclusive.
type SaaSProductLastModifiedDateFilterDateRange struct {

	// Date after which the SaaS product was last modified.
	AfterValue *string

	// Date before which the SaaS product was last modified.
	BeforeValue *string

	noSmithyDocumentSerde
}

// Objects that allows sorting on SaaS products based on certain fields and
// sorting order.
type SaaSProductSort struct {

	// Field to sort the SaaS products by.
	SortBy SaaSProductSortBy

	// The sorting order. Can be ASCENDING or DESCENDING . The default value is
	// DESCENDING .
	SortOrder SortOrder

	noSmithyDocumentSerde
}

// Object that contains summarized information about a SaaS product.
type SaaSProductSummary struct {

	// The title of the SaaS product.
	ProductTitle *string

	// The lifecycle of the SaaS product.
	Visibility SaaSProductVisibilityString

	noSmithyDocumentSerde
}

// Object that allows filtering on product title.
type SaaSProductTitleFilter struct {

	// A string array of unique product title values to be filtered on.
	ValueList []string

	// A string that will be the wildCard input for product tile filter. It matches
	// the provided value as a substring in the actual value.
	WildCardValue *string

	noSmithyDocumentSerde
}

// Object that allows filtering on the visibility of the product in the AWS
// Marketplace.
type SaaSProductVisibilityFilter struct {

	// A string array of unique visibility values to be filtered on.
	ValueList []SaaSProductVisibilityString

	noSmithyDocumentSerde
}

// An object that contains two attributes, SortBy and SortOrder .
type Sort struct {

	// For ListEntities , supported attributes include LastModifiedDate (default) and
	// EntityId . In addition to LastModifiedDate and EntityId , each EntityType might
	// support additional fields. For ListChangeSets , supported attributes include
	// StartTime and EndTime .
	SortBy *string

	// The sorting order. Can be ASCENDING or DESCENDING . The default value is
	// DESCENDING .
	SortOrder SortOrder

	noSmithyDocumentSerde
}

// A list of objects specifying each key name and value.
type Tag struct {

	// The key associated with the tag.
	//
	// This member is required.
	Key *string

	// The value associated with the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isEntityTypeFilters() {}
func (*UnknownUnionMember) isEntityTypeSort()    {}

# ListInvoicesResponseInvoicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The id of the customer that the invoice is associated with | 
**ContractId** | Pointer to **string** | The id of the contract that the invoice is associated with | [optional] 
**Name** | **string** | The name of the invoice, usually a description of the billing period | 
**BillingCycle** | [**Period**](Period.md) |  | 
**LineItems** | [**[]GetInvoiceResponseInvoiceLineItemsInner**](GetInvoiceResponseInvoiceLineItemsInner.md) |  | 
**Amount** | **float32** | The total amount of the invoice | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewListInvoicesResponseInvoicesInner

`func NewListInvoicesResponseInvoicesInner(customerId string, name string, billingCycle Period, lineItems []GetInvoiceResponseInvoiceLineItemsInner, amount float32, id string, createdAt time.Time, updatedAt time.Time, ) *ListInvoicesResponseInvoicesInner`

NewListInvoicesResponseInvoicesInner instantiates a new ListInvoicesResponseInvoicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoicesResponseInvoicesInnerWithDefaults

`func NewListInvoicesResponseInvoicesInnerWithDefaults() *ListInvoicesResponseInvoicesInner`

NewListInvoicesResponseInvoicesInnerWithDefaults instantiates a new ListInvoicesResponseInvoicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *ListInvoicesResponseInvoicesInner) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ListInvoicesResponseInvoicesInner) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ListInvoicesResponseInvoicesInner) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetContractId

`func (o *ListInvoicesResponseInvoicesInner) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ListInvoicesResponseInvoicesInner) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ListInvoicesResponseInvoicesInner) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ListInvoicesResponseInvoicesInner) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetName

`func (o *ListInvoicesResponseInvoicesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInvoicesResponseInvoicesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInvoicesResponseInvoicesInner) SetName(v string)`

SetName sets Name field to given value.


### GetBillingCycle

`func (o *ListInvoicesResponseInvoicesInner) GetBillingCycle() Period`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *ListInvoicesResponseInvoicesInner) GetBillingCycleOk() (*Period, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *ListInvoicesResponseInvoicesInner) SetBillingCycle(v Period)`

SetBillingCycle sets BillingCycle field to given value.


### GetLineItems

`func (o *ListInvoicesResponseInvoicesInner) GetLineItems() []GetInvoiceResponseInvoiceLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *ListInvoicesResponseInvoicesInner) GetLineItemsOk() (*[]GetInvoiceResponseInvoiceLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *ListInvoicesResponseInvoicesInner) SetLineItems(v []GetInvoiceResponseInvoiceLineItemsInner)`

SetLineItems sets LineItems field to given value.


### GetAmount

`func (o *ListInvoicesResponseInvoicesInner) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ListInvoicesResponseInvoicesInner) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ListInvoicesResponseInvoicesInner) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *ListInvoicesResponseInvoicesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInvoicesResponseInvoicesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInvoicesResponseInvoicesInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListInvoicesResponseInvoicesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListInvoicesResponseInvoicesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListInvoicesResponseInvoicesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListInvoicesResponseInvoicesInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListInvoicesResponseInvoicesInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListInvoicesResponseInvoicesInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetInvoiceResponseInvoice

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
**DeletedAt** | **string** |  | 

## Methods

### NewGetInvoiceResponseInvoice

`func NewGetInvoiceResponseInvoice(customerId string, name string, billingCycle Period, lineItems []GetInvoiceResponseInvoiceLineItemsInner, amount float32, id string, createdAt time.Time, updatedAt time.Time, deletedAt string, ) *GetInvoiceResponseInvoice`

NewGetInvoiceResponseInvoice instantiates a new GetInvoiceResponseInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceResponseInvoiceWithDefaults

`func NewGetInvoiceResponseInvoiceWithDefaults() *GetInvoiceResponseInvoice`

NewGetInvoiceResponseInvoiceWithDefaults instantiates a new GetInvoiceResponseInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *GetInvoiceResponseInvoice) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *GetInvoiceResponseInvoice) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *GetInvoiceResponseInvoice) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetContractId

`func (o *GetInvoiceResponseInvoice) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *GetInvoiceResponseInvoice) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *GetInvoiceResponseInvoice) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *GetInvoiceResponseInvoice) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetName

`func (o *GetInvoiceResponseInvoice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInvoiceResponseInvoice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInvoiceResponseInvoice) SetName(v string)`

SetName sets Name field to given value.


### GetBillingCycle

`func (o *GetInvoiceResponseInvoice) GetBillingCycle() Period`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *GetInvoiceResponseInvoice) GetBillingCycleOk() (*Period, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *GetInvoiceResponseInvoice) SetBillingCycle(v Period)`

SetBillingCycle sets BillingCycle field to given value.


### GetLineItems

`func (o *GetInvoiceResponseInvoice) GetLineItems() []GetInvoiceResponseInvoiceLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *GetInvoiceResponseInvoice) GetLineItemsOk() (*[]GetInvoiceResponseInvoiceLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *GetInvoiceResponseInvoice) SetLineItems(v []GetInvoiceResponseInvoiceLineItemsInner)`

SetLineItems sets LineItems field to given value.


### GetAmount

`func (o *GetInvoiceResponseInvoice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetInvoiceResponseInvoice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetInvoiceResponseInvoice) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *GetInvoiceResponseInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInvoiceResponseInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInvoiceResponseInvoice) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetInvoiceResponseInvoice) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetInvoiceResponseInvoice) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetInvoiceResponseInvoice) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetInvoiceResponseInvoice) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetInvoiceResponseInvoice) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetInvoiceResponseInvoice) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *GetInvoiceResponseInvoice) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GetInvoiceResponseInvoice) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GetInvoiceResponseInvoice) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



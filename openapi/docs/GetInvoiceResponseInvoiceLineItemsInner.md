# GetInvoiceResponseInvoiceLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The id of the invoice that the line item is a part of | 
**Price** | **float32** | The price of the line item | 

## Methods

### NewGetInvoiceResponseInvoiceLineItemsInner

`func NewGetInvoiceResponseInvoiceLineItemsInner(invoiceId string, price float32, ) *GetInvoiceResponseInvoiceLineItemsInner`

NewGetInvoiceResponseInvoiceLineItemsInner instantiates a new GetInvoiceResponseInvoiceLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceResponseInvoiceLineItemsInnerWithDefaults

`func NewGetInvoiceResponseInvoiceLineItemsInnerWithDefaults() *GetInvoiceResponseInvoiceLineItemsInner`

NewGetInvoiceResponseInvoiceLineItemsInnerWithDefaults instantiates a new GetInvoiceResponseInvoiceLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *GetInvoiceResponseInvoiceLineItemsInner) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *GetInvoiceResponseInvoiceLineItemsInner) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *GetInvoiceResponseInvoiceLineItemsInner) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetPrice

`func (o *GetInvoiceResponseInvoiceLineItemsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetInvoiceResponseInvoiceLineItemsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetInvoiceResponseInvoiceLineItemsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



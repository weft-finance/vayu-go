# CreditLedgerEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**BillingCycleStatus**](BillingCycleStatus.md) |  | 
**Amount** | **float32** |  | 
**InvoiceId** | Pointer to **string** |  | [optional] 
**BalanceAfterEntry** | **float32** |  | 

## Methods

### NewCreditLedgerEntry

`func NewCreditLedgerEntry(type_ BillingCycleStatus, amount float32, balanceAfterEntry float32, ) *CreditLedgerEntry`

NewCreditLedgerEntry instantiates a new CreditLedgerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditLedgerEntryWithDefaults

`func NewCreditLedgerEntryWithDefaults() *CreditLedgerEntry`

NewCreditLedgerEntryWithDefaults instantiates a new CreditLedgerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreditLedgerEntry) GetType() BillingCycleStatus`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreditLedgerEntry) GetTypeOk() (*BillingCycleStatus, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreditLedgerEntry) SetType(v BillingCycleStatus)`

SetType sets Type field to given value.


### GetAmount

`func (o *CreditLedgerEntry) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreditLedgerEntry) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreditLedgerEntry) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetInvoiceId

`func (o *CreditLedgerEntry) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *CreditLedgerEntry) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *CreditLedgerEntry) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *CreditLedgerEntry) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetBalanceAfterEntry

`func (o *CreditLedgerEntry) GetBalanceAfterEntry() float32`

GetBalanceAfterEntry returns the BalanceAfterEntry field if non-nil, zero value otherwise.

### GetBalanceAfterEntryOk

`func (o *CreditLedgerEntry) GetBalanceAfterEntryOk() (*float32, bool)`

GetBalanceAfterEntryOk returns a tuple with the BalanceAfterEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfterEntry

`func (o *CreditLedgerEntry) SetBalanceAfterEntry(v float32)`

SetBalanceAfterEntry sets BalanceAfterEntry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



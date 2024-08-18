# ListCreditLedgerEntriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]CreditLedgerEntry**](CreditLedgerEntry.md) | The credit ledger entries for the customer. | 

## Methods

### NewListCreditLedgerEntriesResponse

`func NewListCreditLedgerEntriesResponse(entries []CreditLedgerEntry, ) *ListCreditLedgerEntriesResponse`

NewListCreditLedgerEntriesResponse instantiates a new ListCreditLedgerEntriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCreditLedgerEntriesResponseWithDefaults

`func NewListCreditLedgerEntriesResponseWithDefaults() *ListCreditLedgerEntriesResponse`

NewListCreditLedgerEntriesResponseWithDefaults instantiates a new ListCreditLedgerEntriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *ListCreditLedgerEntriesResponse) GetEntries() []CreditLedgerEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ListCreditLedgerEntriesResponse) GetEntriesOk() (*[]CreditLedgerEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ListCreditLedgerEntriesResponse) SetEntries(v []CreditLedgerEntry)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NetSuiteSyncInvoicesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **time.Time** |  | 
**EndDate** | **time.Time** |  | 
**TranDate** | **time.Time** |  | 
**Memo** | **string** |  | 
**ExternalId** | **string** |  | 
**Entity** | [**NetSuiteSyncInvoicesRequestDataEntity**](NetSuiteSyncInvoicesRequestDataEntity.md) |  | 
**Account** | Pointer to [**NetSuiteSyncInvoicesRequestDataEntity**](NetSuiteSyncInvoicesRequestDataEntity.md) |  | [optional] 
**Item** | [**NetSuiteSyncInvoicesRequestDataItem**](NetSuiteSyncInvoicesRequestDataItem.md) |  | 
**DiscountItem** | Pointer to [**NetSuiteSyncInvoicesRequestDataEntity**](NetSuiteSyncInvoicesRequestDataEntity.md) |  | [optional] 
**DiscountRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewNetSuiteSyncInvoicesRequestData

`func NewNetSuiteSyncInvoicesRequestData(startDate time.Time, endDate time.Time, tranDate time.Time, memo string, externalId string, entity NetSuiteSyncInvoicesRequestDataEntity, item NetSuiteSyncInvoicesRequestDataItem, ) *NetSuiteSyncInvoicesRequestData`

NewNetSuiteSyncInvoicesRequestData instantiates a new NetSuiteSyncInvoicesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetSuiteSyncInvoicesRequestDataWithDefaults

`func NewNetSuiteSyncInvoicesRequestDataWithDefaults() *NetSuiteSyncInvoicesRequestData`

NewNetSuiteSyncInvoicesRequestDataWithDefaults instantiates a new NetSuiteSyncInvoicesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *NetSuiteSyncInvoicesRequestData) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NetSuiteSyncInvoicesRequestData) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NetSuiteSyncInvoicesRequestData) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *NetSuiteSyncInvoicesRequestData) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *NetSuiteSyncInvoicesRequestData) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *NetSuiteSyncInvoicesRequestData) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetTranDate

`func (o *NetSuiteSyncInvoicesRequestData) GetTranDate() time.Time`

GetTranDate returns the TranDate field if non-nil, zero value otherwise.

### GetTranDateOk

`func (o *NetSuiteSyncInvoicesRequestData) GetTranDateOk() (*time.Time, bool)`

GetTranDateOk returns a tuple with the TranDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranDate

`func (o *NetSuiteSyncInvoicesRequestData) SetTranDate(v time.Time)`

SetTranDate sets TranDate field to given value.


### GetMemo

`func (o *NetSuiteSyncInvoicesRequestData) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *NetSuiteSyncInvoicesRequestData) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *NetSuiteSyncInvoicesRequestData) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetExternalId

`func (o *NetSuiteSyncInvoicesRequestData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetSuiteSyncInvoicesRequestData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetSuiteSyncInvoicesRequestData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetEntity

`func (o *NetSuiteSyncInvoicesRequestData) GetEntity() NetSuiteSyncInvoicesRequestDataEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *NetSuiteSyncInvoicesRequestData) GetEntityOk() (*NetSuiteSyncInvoicesRequestDataEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *NetSuiteSyncInvoicesRequestData) SetEntity(v NetSuiteSyncInvoicesRequestDataEntity)`

SetEntity sets Entity field to given value.


### GetAccount

`func (o *NetSuiteSyncInvoicesRequestData) GetAccount() NetSuiteSyncInvoicesRequestDataEntity`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetSuiteSyncInvoicesRequestData) GetAccountOk() (*NetSuiteSyncInvoicesRequestDataEntity, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetSuiteSyncInvoicesRequestData) SetAccount(v NetSuiteSyncInvoicesRequestDataEntity)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetSuiteSyncInvoicesRequestData) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetItem

`func (o *NetSuiteSyncInvoicesRequestData) GetItem() NetSuiteSyncInvoicesRequestDataItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *NetSuiteSyncInvoicesRequestData) GetItemOk() (*NetSuiteSyncInvoicesRequestDataItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *NetSuiteSyncInvoicesRequestData) SetItem(v NetSuiteSyncInvoicesRequestDataItem)`

SetItem sets Item field to given value.


### GetDiscountItem

`func (o *NetSuiteSyncInvoicesRequestData) GetDiscountItem() NetSuiteSyncInvoicesRequestDataEntity`

GetDiscountItem returns the DiscountItem field if non-nil, zero value otherwise.

### GetDiscountItemOk

`func (o *NetSuiteSyncInvoicesRequestData) GetDiscountItemOk() (*NetSuiteSyncInvoicesRequestDataEntity, bool)`

GetDiscountItemOk returns a tuple with the DiscountItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountItem

`func (o *NetSuiteSyncInvoicesRequestData) SetDiscountItem(v NetSuiteSyncInvoicesRequestDataEntity)`

SetDiscountItem sets DiscountItem field to given value.

### HasDiscountItem

`func (o *NetSuiteSyncInvoicesRequestData) HasDiscountItem() bool`

HasDiscountItem returns a boolean if a field has been set.

### GetDiscountRate

`func (o *NetSuiteSyncInvoicesRequestData) GetDiscountRate() float32`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *NetSuiteSyncInvoicesRequestData) GetDiscountRateOk() (*float32, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *NetSuiteSyncInvoicesRequestData) SetDiscountRate(v float32)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *NetSuiteSyncInvoicesRequestData) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



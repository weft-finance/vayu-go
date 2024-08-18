# EventsDryRunResponseInnerEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**CustomerAlias** | **string** |  | 
**AccountId** | **string** |  | 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Ref** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewEventsDryRunResponseInnerEvent

`func NewEventsDryRunResponseInnerEvent(timestamp time.Time, customerAlias string, accountId string, ref string, name string, ) *EventsDryRunResponseInnerEvent`

NewEventsDryRunResponseInnerEvent instantiates a new EventsDryRunResponseInnerEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsDryRunResponseInnerEventWithDefaults

`func NewEventsDryRunResponseInnerEventWithDefaults() *EventsDryRunResponseInnerEvent`

NewEventsDryRunResponseInnerEventWithDefaults instantiates a new EventsDryRunResponseInnerEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *EventsDryRunResponseInnerEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventsDryRunResponseInnerEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventsDryRunResponseInnerEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetCustomerAlias

`func (o *EventsDryRunResponseInnerEvent) GetCustomerAlias() string`

GetCustomerAlias returns the CustomerAlias field if non-nil, zero value otherwise.

### GetCustomerAliasOk

`func (o *EventsDryRunResponseInnerEvent) GetCustomerAliasOk() (*string, bool)`

GetCustomerAliasOk returns a tuple with the CustomerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAlias

`func (o *EventsDryRunResponseInnerEvent) SetCustomerAlias(v string)`

SetCustomerAlias sets CustomerAlias field to given value.


### GetAccountId

`func (o *EventsDryRunResponseInnerEvent) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EventsDryRunResponseInnerEvent) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EventsDryRunResponseInnerEvent) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetData

`func (o *EventsDryRunResponseInnerEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventsDryRunResponseInnerEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventsDryRunResponseInnerEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventsDryRunResponseInnerEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EventsDryRunResponseInnerEvent) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EventsDryRunResponseInnerEvent) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetRef

`func (o *EventsDryRunResponseInnerEvent) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *EventsDryRunResponseInnerEvent) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *EventsDryRunResponseInnerEvent) SetRef(v string)`

SetRef sets Ref field to given value.


### GetName

`func (o *EventsDryRunResponseInnerEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventsDryRunResponseInnerEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventsDryRunResponseInnerEvent) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



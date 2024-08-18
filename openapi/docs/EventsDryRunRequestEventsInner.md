# EventsDryRunRequestEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The distinctive label assigned to an event, serving as a critical identifier for categorizing and pricing events within the system&#39;s backend infrastructure. | 
**Timestamp** | **time.Time** | The temporal marker denoting the exact moment of event occurrence. The timestamp is formatted as an ISO 8601 string and is expressed in Coordinated Universal Time (UTC). i.e. YYYY-MM-DDTHH:MM:SS.SSSZ | 
**CustomerAlias** | **string** | A pseudonymous or otherwise obfuscated identifier uniquely assigned to each customer. | 
**Ref** | **string** | A universally unique identifier (UUID) or other form of high-entropy string serving as an immutable reference for each event entry. | 
**Data** | Pointer to **map[string]interface{}** | A schema-less JSON object encapsulating miscellaneous attributes and metrics associated with the event. | [optional] 

## Methods

### NewEventsDryRunRequestEventsInner

`func NewEventsDryRunRequestEventsInner(name string, timestamp time.Time, customerAlias string, ref string, ) *EventsDryRunRequestEventsInner`

NewEventsDryRunRequestEventsInner instantiates a new EventsDryRunRequestEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsDryRunRequestEventsInnerWithDefaults

`func NewEventsDryRunRequestEventsInnerWithDefaults() *EventsDryRunRequestEventsInner`

NewEventsDryRunRequestEventsInnerWithDefaults instantiates a new EventsDryRunRequestEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventsDryRunRequestEventsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventsDryRunRequestEventsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventsDryRunRequestEventsInner) SetName(v string)`

SetName sets Name field to given value.


### GetTimestamp

`func (o *EventsDryRunRequestEventsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EventsDryRunRequestEventsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EventsDryRunRequestEventsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetCustomerAlias

`func (o *EventsDryRunRequestEventsInner) GetCustomerAlias() string`

GetCustomerAlias returns the CustomerAlias field if non-nil, zero value otherwise.

### GetCustomerAliasOk

`func (o *EventsDryRunRequestEventsInner) GetCustomerAliasOk() (*string, bool)`

GetCustomerAliasOk returns a tuple with the CustomerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAlias

`func (o *EventsDryRunRequestEventsInner) SetCustomerAlias(v string)`

SetCustomerAlias sets CustomerAlias field to given value.


### GetRef

`func (o *EventsDryRunRequestEventsInner) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *EventsDryRunRequestEventsInner) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *EventsDryRunRequestEventsInner) SetRef(v string)`

SetRef sets Ref field to given value.


### GetData

`func (o *EventsDryRunRequestEventsInner) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventsDryRunRequestEventsInner) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventsDryRunRequestEventsInner) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventsDryRunRequestEventsInner) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *EventsDryRunRequestEventsInner) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *EventsDryRunRequestEventsInner) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DeleteEventByRefIdResponseEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The distinctive label assigned to an event, serving as a critical identifier for categorizing and pricing events within the system&#39;s backend infrastructure. | 
**Timestamp** | **time.Time** | The temporal marker denoting the exact moment of event occurrence. The timestamp is formatted as an ISO 8601 string and is expressed in Coordinated Universal Time (UTC). i.e. YYYY-MM-DDTHH:MM:SS.SSSZ | 
**CustomerAlias** | **string** | A pseudonymous or otherwise obfuscated identifier uniquely assigned to each customer. | 
**Ref** | **string** | A universally unique identifier (UUID) or other form of high-entropy string serving as an immutable reference for each event entry. | 
**Data** | Pointer to **map[string]interface{}** | A schema-less JSON object encapsulating miscellaneous attributes and metrics associated with the event. | [optional] 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **string** |  | 

## Methods

### NewDeleteEventByRefIdResponseEvent

`func NewDeleteEventByRefIdResponseEvent(name string, timestamp time.Time, customerAlias string, ref string, id string, createdAt time.Time, updatedAt time.Time, deletedAt string, ) *DeleteEventByRefIdResponseEvent`

NewDeleteEventByRefIdResponseEvent instantiates a new DeleteEventByRefIdResponseEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteEventByRefIdResponseEventWithDefaults

`func NewDeleteEventByRefIdResponseEventWithDefaults() *DeleteEventByRefIdResponseEvent`

NewDeleteEventByRefIdResponseEventWithDefaults instantiates a new DeleteEventByRefIdResponseEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeleteEventByRefIdResponseEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeleteEventByRefIdResponseEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeleteEventByRefIdResponseEvent) SetName(v string)`

SetName sets Name field to given value.


### GetTimestamp

`func (o *DeleteEventByRefIdResponseEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeleteEventByRefIdResponseEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeleteEventByRefIdResponseEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetCustomerAlias

`func (o *DeleteEventByRefIdResponseEvent) GetCustomerAlias() string`

GetCustomerAlias returns the CustomerAlias field if non-nil, zero value otherwise.

### GetCustomerAliasOk

`func (o *DeleteEventByRefIdResponseEvent) GetCustomerAliasOk() (*string, bool)`

GetCustomerAliasOk returns a tuple with the CustomerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAlias

`func (o *DeleteEventByRefIdResponseEvent) SetCustomerAlias(v string)`

SetCustomerAlias sets CustomerAlias field to given value.


### GetRef

`func (o *DeleteEventByRefIdResponseEvent) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DeleteEventByRefIdResponseEvent) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DeleteEventByRefIdResponseEvent) SetRef(v string)`

SetRef sets Ref field to given value.


### GetData

`func (o *DeleteEventByRefIdResponseEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteEventByRefIdResponseEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteEventByRefIdResponseEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *DeleteEventByRefIdResponseEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *DeleteEventByRefIdResponseEvent) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *DeleteEventByRefIdResponseEvent) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetId

`func (o *DeleteEventByRefIdResponseEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteEventByRefIdResponseEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteEventByRefIdResponseEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeleteEventByRefIdResponseEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeleteEventByRefIdResponseEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeleteEventByRefIdResponseEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeleteEventByRefIdResponseEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeleteEventByRefIdResponseEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeleteEventByRefIdResponseEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DeleteEventByRefIdResponseEvent) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DeleteEventByRefIdResponseEvent) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DeleteEventByRefIdResponseEvent) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



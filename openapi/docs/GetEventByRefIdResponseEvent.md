# GetEventByRefIdResponseEvent

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

## Methods

### NewGetEventByRefIdResponseEvent

`func NewGetEventByRefIdResponseEvent(name string, timestamp time.Time, customerAlias string, ref string, id string, createdAt time.Time, updatedAt time.Time, ) *GetEventByRefIdResponseEvent`

NewGetEventByRefIdResponseEvent instantiates a new GetEventByRefIdResponseEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventByRefIdResponseEventWithDefaults

`func NewGetEventByRefIdResponseEventWithDefaults() *GetEventByRefIdResponseEvent`

NewGetEventByRefIdResponseEventWithDefaults instantiates a new GetEventByRefIdResponseEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetEventByRefIdResponseEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEventByRefIdResponseEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEventByRefIdResponseEvent) SetName(v string)`

SetName sets Name field to given value.


### GetTimestamp

`func (o *GetEventByRefIdResponseEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetEventByRefIdResponseEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetEventByRefIdResponseEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetCustomerAlias

`func (o *GetEventByRefIdResponseEvent) GetCustomerAlias() string`

GetCustomerAlias returns the CustomerAlias field if non-nil, zero value otherwise.

### GetCustomerAliasOk

`func (o *GetEventByRefIdResponseEvent) GetCustomerAliasOk() (*string, bool)`

GetCustomerAliasOk returns a tuple with the CustomerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAlias

`func (o *GetEventByRefIdResponseEvent) SetCustomerAlias(v string)`

SetCustomerAlias sets CustomerAlias field to given value.


### GetRef

`func (o *GetEventByRefIdResponseEvent) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetEventByRefIdResponseEvent) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetEventByRefIdResponseEvent) SetRef(v string)`

SetRef sets Ref field to given value.


### GetData

`func (o *GetEventByRefIdResponseEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEventByRefIdResponseEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEventByRefIdResponseEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *GetEventByRefIdResponseEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GetEventByRefIdResponseEvent) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GetEventByRefIdResponseEvent) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetId

`func (o *GetEventByRefIdResponseEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEventByRefIdResponseEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEventByRefIdResponseEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetEventByRefIdResponseEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetEventByRefIdResponseEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetEventByRefIdResponseEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetEventByRefIdResponseEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetEventByRefIdResponseEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetEventByRefIdResponseEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



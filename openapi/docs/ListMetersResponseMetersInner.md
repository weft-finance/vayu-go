# ListMetersResponseMetersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the meter | 
**EventName** | **string** | The name of the event that the meter is tracking. | 
**AggregationMethod** | [**GetMeterResponseMeterAggregationMethod**](GetMeterResponseMeterAggregationMethod.md) |  | 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewListMetersResponseMetersInner

`func NewListMetersResponseMetersInner(name string, eventName string, aggregationMethod GetMeterResponseMeterAggregationMethod, id string, createdAt time.Time, updatedAt time.Time, ) *ListMetersResponseMetersInner`

NewListMetersResponseMetersInner instantiates a new ListMetersResponseMetersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetersResponseMetersInnerWithDefaults

`func NewListMetersResponseMetersInnerWithDefaults() *ListMetersResponseMetersInner`

NewListMetersResponseMetersInnerWithDefaults instantiates a new ListMetersResponseMetersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListMetersResponseMetersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMetersResponseMetersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMetersResponseMetersInner) SetName(v string)`

SetName sets Name field to given value.


### GetEventName

`func (o *ListMetersResponseMetersInner) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ListMetersResponseMetersInner) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ListMetersResponseMetersInner) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetAggregationMethod

`func (o *ListMetersResponseMetersInner) GetAggregationMethod() GetMeterResponseMeterAggregationMethod`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *ListMetersResponseMetersInner) GetAggregationMethodOk() (*GetMeterResponseMeterAggregationMethod, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *ListMetersResponseMetersInner) SetAggregationMethod(v GetMeterResponseMeterAggregationMethod)`

SetAggregationMethod sets AggregationMethod field to given value.


### GetFilter

`func (o *ListMetersResponseMetersInner) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListMetersResponseMetersInner) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListMetersResponseMetersInner) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListMetersResponseMetersInner) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *ListMetersResponseMetersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMetersResponseMetersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMetersResponseMetersInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListMetersResponseMetersInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListMetersResponseMetersInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListMetersResponseMetersInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListMetersResponseMetersInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListMetersResponseMetersInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListMetersResponseMetersInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



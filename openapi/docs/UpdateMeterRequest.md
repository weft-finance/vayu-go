# UpdateMeterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the meter | [optional] 
**EventName** | Pointer to **string** | The name of the event that the meter is tracking. | [optional] 
**AggregationMethod** | Pointer to [**GetMeterResponseMeterAggregationMethod**](GetMeterResponseMeterAggregationMethod.md) |  | [optional] 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 

## Methods

### NewUpdateMeterRequest

`func NewUpdateMeterRequest() *UpdateMeterRequest`

NewUpdateMeterRequest instantiates a new UpdateMeterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMeterRequestWithDefaults

`func NewUpdateMeterRequestWithDefaults() *UpdateMeterRequest`

NewUpdateMeterRequestWithDefaults instantiates a new UpdateMeterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMeterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMeterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMeterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateMeterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventName

`func (o *UpdateMeterRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *UpdateMeterRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *UpdateMeterRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *UpdateMeterRequest) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetAggregationMethod

`func (o *UpdateMeterRequest) GetAggregationMethod() GetMeterResponseMeterAggregationMethod`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *UpdateMeterRequest) GetAggregationMethodOk() (*GetMeterResponseMeterAggregationMethod, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *UpdateMeterRequest) SetAggregationMethod(v GetMeterResponseMeterAggregationMethod)`

SetAggregationMethod sets AggregationMethod field to given value.

### HasAggregationMethod

`func (o *UpdateMeterRequest) HasAggregationMethod() bool`

HasAggregationMethod returns a boolean if a field has been set.

### GetFilter

`func (o *UpdateMeterRequest) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateMeterRequest) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateMeterRequest) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UpdateMeterRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



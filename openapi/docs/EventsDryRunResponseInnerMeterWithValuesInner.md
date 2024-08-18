# EventsDryRunResponseInnerMeterWithValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the meter | 
**EventName** | **string** | The name of the event that the meter is tracking. | 
**AggregationMethod** | [**GetMeterResponseMeterAggregationMethod**](GetMeterResponseMeterAggregationMethod.md) |  | 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 
**Value** | **float32** |  | 
**InstanceValue** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewEventsDryRunResponseInnerMeterWithValuesInner

`func NewEventsDryRunResponseInnerMeterWithValuesInner(name string, eventName string, aggregationMethod GetMeterResponseMeterAggregationMethod, value float32, ) *EventsDryRunResponseInnerMeterWithValuesInner`

NewEventsDryRunResponseInnerMeterWithValuesInner instantiates a new EventsDryRunResponseInnerMeterWithValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsDryRunResponseInnerMeterWithValuesInnerWithDefaults

`func NewEventsDryRunResponseInnerMeterWithValuesInnerWithDefaults() *EventsDryRunResponseInnerMeterWithValuesInner`

NewEventsDryRunResponseInnerMeterWithValuesInnerWithDefaults instantiates a new EventsDryRunResponseInnerMeterWithValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) SetName(v string)`

SetName sets Name field to given value.


### GetEventName

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetAggregationMethod

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetAggregationMethod() GetMeterResponseMeterAggregationMethod`

GetAggregationMethod returns the AggregationMethod field if non-nil, zero value otherwise.

### GetAggregationMethodOk

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetAggregationMethodOk() (*GetMeterResponseMeterAggregationMethod, bool)`

GetAggregationMethodOk returns a tuple with the AggregationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationMethod

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) SetAggregationMethod(v GetMeterResponseMeterAggregationMethod)`

SetAggregationMethod sets AggregationMethod field to given value.


### GetFilter

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetValue

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) SetValue(v float32)`

SetValue sets Value field to given value.


### GetInstanceValue

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetInstanceValue() interface{}`

GetInstanceValue returns the InstanceValue field if non-nil, zero value otherwise.

### GetInstanceValueOk

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) GetInstanceValueOk() (*interface{}, bool)`

GetInstanceValueOk returns a tuple with the InstanceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceValue

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) SetInstanceValue(v interface{})`

SetInstanceValue sets InstanceValue field to given value.

### HasInstanceValue

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) HasInstanceValue() bool`

HasInstanceValue returns a boolean if a field has been set.

### SetInstanceValueNil

`func (o *EventsDryRunResponseInnerMeterWithValuesInner) SetInstanceValueNil(b bool)`

 SetInstanceValueNil sets the value for InstanceValue to be an explicit nil

### UnsetInstanceValue
`func (o *EventsDryRunResponseInnerMeterWithValuesInner) UnsetInstanceValue()`

UnsetInstanceValue ensures that no value is present for InstanceValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



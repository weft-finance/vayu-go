# GetMeterResponseMeterAggregationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | [**AggregationOperator**](AggregationOperator.md) |  | 
**Field** | Pointer to **string** | The field inside the \&quot;data\&quot; property to apply the aggregation method to. | [optional] 
**InstanceKey** | Pointer to **string** | The fields that is used to group different events together when measuring the usage. | [optional] 

## Methods

### NewGetMeterResponseMeterAggregationMethod

`func NewGetMeterResponseMeterAggregationMethod(operator AggregationOperator, ) *GetMeterResponseMeterAggregationMethod`

NewGetMeterResponseMeterAggregationMethod instantiates a new GetMeterResponseMeterAggregationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMeterResponseMeterAggregationMethodWithDefaults

`func NewGetMeterResponseMeterAggregationMethodWithDefaults() *GetMeterResponseMeterAggregationMethod`

NewGetMeterResponseMeterAggregationMethodWithDefaults instantiates a new GetMeterResponseMeterAggregationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *GetMeterResponseMeterAggregationMethod) GetOperator() AggregationOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GetMeterResponseMeterAggregationMethod) GetOperatorOk() (*AggregationOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GetMeterResponseMeterAggregationMethod) SetOperator(v AggregationOperator)`

SetOperator sets Operator field to given value.


### GetField

`func (o *GetMeterResponseMeterAggregationMethod) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GetMeterResponseMeterAggregationMethod) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GetMeterResponseMeterAggregationMethod) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *GetMeterResponseMeterAggregationMethod) HasField() bool`

HasField returns a boolean if a field has been set.

### GetInstanceKey

`func (o *GetMeterResponseMeterAggregationMethod) GetInstanceKey() string`

GetInstanceKey returns the InstanceKey field if non-nil, zero value otherwise.

### GetInstanceKeyOk

`func (o *GetMeterResponseMeterAggregationMethod) GetInstanceKeyOk() (*string, bool)`

GetInstanceKeyOk returns a tuple with the InstanceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceKey

`func (o *GetMeterResponseMeterAggregationMethod) SetInstanceKey(v string)`

SetInstanceKey sets InstanceKey field to given value.

### HasInstanceKey

`func (o *GetMeterResponseMeterAggregationMethod) HasInstanceKey() bool`

HasInstanceKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



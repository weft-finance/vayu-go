# Criterion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Operator** | [**CriterionOperator**](CriterionOperator.md) |  | 
**Value** | **string** |  | 

## Methods

### NewCriterion

`func NewCriterion(field string, operator CriterionOperator, value string, ) *Criterion`

NewCriterion instantiates a new Criterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriterionWithDefaults

`func NewCriterionWithDefaults() *Criterion`

NewCriterionWithDefaults instantiates a new Criterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *Criterion) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Criterion) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Criterion) SetField(v string)`

SetField sets Field field to given value.


### GetOperator

`func (o *Criterion) GetOperator() CriterionOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Criterion) GetOperatorOk() (*CriterionOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Criterion) SetOperator(v CriterionOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *Criterion) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Criterion) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Criterion) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



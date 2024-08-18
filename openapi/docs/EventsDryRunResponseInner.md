# EventsDryRunResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**EventsDryRunResponseInnerEvent**](EventsDryRunResponseInnerEvent.md) |  | 
**MatchedCustomer** | Pointer to **NullableString** |  | [optional] 
**MeterWithValues** | [**[]EventsDryRunResponseInnerMeterWithValuesInner**](EventsDryRunResponseInnerMeterWithValuesInner.md) |  | 

## Methods

### NewEventsDryRunResponseInner

`func NewEventsDryRunResponseInner(event EventsDryRunResponseInnerEvent, meterWithValues []EventsDryRunResponseInnerMeterWithValuesInner, ) *EventsDryRunResponseInner`

NewEventsDryRunResponseInner instantiates a new EventsDryRunResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsDryRunResponseInnerWithDefaults

`func NewEventsDryRunResponseInnerWithDefaults() *EventsDryRunResponseInner`

NewEventsDryRunResponseInnerWithDefaults instantiates a new EventsDryRunResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventsDryRunResponseInner) GetEvent() EventsDryRunResponseInnerEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventsDryRunResponseInner) GetEventOk() (*EventsDryRunResponseInnerEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventsDryRunResponseInner) SetEvent(v EventsDryRunResponseInnerEvent)`

SetEvent sets Event field to given value.


### GetMatchedCustomer

`func (o *EventsDryRunResponseInner) GetMatchedCustomer() string`

GetMatchedCustomer returns the MatchedCustomer field if non-nil, zero value otherwise.

### GetMatchedCustomerOk

`func (o *EventsDryRunResponseInner) GetMatchedCustomerOk() (*string, bool)`

GetMatchedCustomerOk returns a tuple with the MatchedCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedCustomer

`func (o *EventsDryRunResponseInner) SetMatchedCustomer(v string)`

SetMatchedCustomer sets MatchedCustomer field to given value.

### HasMatchedCustomer

`func (o *EventsDryRunResponseInner) HasMatchedCustomer() bool`

HasMatchedCustomer returns a boolean if a field has been set.

### SetMatchedCustomerNil

`func (o *EventsDryRunResponseInner) SetMatchedCustomerNil(b bool)`

 SetMatchedCustomerNil sets the value for MatchedCustomer to be an explicit nil

### UnsetMatchedCustomer
`func (o *EventsDryRunResponseInner) UnsetMatchedCustomer()`

UnsetMatchedCustomer ensures that no value is present for MatchedCustomer, not even an explicit nil
### GetMeterWithValues

`func (o *EventsDryRunResponseInner) GetMeterWithValues() []EventsDryRunResponseInnerMeterWithValuesInner`

GetMeterWithValues returns the MeterWithValues field if non-nil, zero value otherwise.

### GetMeterWithValuesOk

`func (o *EventsDryRunResponseInner) GetMeterWithValuesOk() (*[]EventsDryRunResponseInnerMeterWithValuesInner, bool)`

GetMeterWithValuesOk returns a tuple with the MeterWithValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterWithValues

`func (o *EventsDryRunResponseInner) SetMeterWithValues(v []EventsDryRunResponseInnerMeterWithValuesInner)`

SetMeterWithValues sets MeterWithValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



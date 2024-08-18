# SendEventsResponseInvalidEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**EventsDryRunRequestEventsInner**](EventsDryRunRequestEventsInner.md) |  | 
**Error** | **string** | The error message indicating the reason the event failed validation. | 

## Methods

### NewSendEventsResponseInvalidEventsInner

`func NewSendEventsResponseInvalidEventsInner(event EventsDryRunRequestEventsInner, error_ string, ) *SendEventsResponseInvalidEventsInner`

NewSendEventsResponseInvalidEventsInner instantiates a new SendEventsResponseInvalidEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEventsResponseInvalidEventsInnerWithDefaults

`func NewSendEventsResponseInvalidEventsInnerWithDefaults() *SendEventsResponseInvalidEventsInner`

NewSendEventsResponseInvalidEventsInnerWithDefaults instantiates a new SendEventsResponseInvalidEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *SendEventsResponseInvalidEventsInner) GetEvent() EventsDryRunRequestEventsInner`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SendEventsResponseInvalidEventsInner) GetEventOk() (*EventsDryRunRequestEventsInner, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SendEventsResponseInvalidEventsInner) SetEvent(v EventsDryRunRequestEventsInner)`

SetEvent sets Event field to given value.


### GetError

`func (o *SendEventsResponseInvalidEventsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SendEventsResponseInvalidEventsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SendEventsResponseInvalidEventsInner) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



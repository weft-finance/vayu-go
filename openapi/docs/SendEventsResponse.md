# SendEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidEvents** | [**[]EventsDryRunRequestEventsInner**](EventsDryRunRequestEventsInner.md) | An array of events that were successfully processed and sent to the queue. | 
**InvalidEvents** | [**[]SendEventsResponseInvalidEventsInner**](SendEventsResponseInvalidEventsInner.md) | An array of events that failed validation and were not sent to the queue. Each object contains the event and the error message. | 

## Methods

### NewSendEventsResponse

`func NewSendEventsResponse(validEvents []EventsDryRunRequestEventsInner, invalidEvents []SendEventsResponseInvalidEventsInner, ) *SendEventsResponse`

NewSendEventsResponse instantiates a new SendEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEventsResponseWithDefaults

`func NewSendEventsResponseWithDefaults() *SendEventsResponse`

NewSendEventsResponseWithDefaults instantiates a new SendEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidEvents

`func (o *SendEventsResponse) GetValidEvents() []EventsDryRunRequestEventsInner`

GetValidEvents returns the ValidEvents field if non-nil, zero value otherwise.

### GetValidEventsOk

`func (o *SendEventsResponse) GetValidEventsOk() (*[]EventsDryRunRequestEventsInner, bool)`

GetValidEventsOk returns a tuple with the ValidEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidEvents

`func (o *SendEventsResponse) SetValidEvents(v []EventsDryRunRequestEventsInner)`

SetValidEvents sets ValidEvents field to given value.


### GetInvalidEvents

`func (o *SendEventsResponse) GetInvalidEvents() []SendEventsResponseInvalidEventsInner`

GetInvalidEvents returns the InvalidEvents field if non-nil, zero value otherwise.

### GetInvalidEventsOk

`func (o *SendEventsResponse) GetInvalidEventsOk() (*[]SendEventsResponseInvalidEventsInner, bool)`

GetInvalidEventsOk returns a tuple with the InvalidEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidEvents

`func (o *SendEventsResponse) SetInvalidEvents(v []SendEventsResponseInvalidEventsInner)`

SetInvalidEvents sets InvalidEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SendEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]EventsDryRunRequestEventsInner**](EventsDryRunRequestEventsInner.md) |  | 

## Methods

### NewSendEventsRequest

`func NewSendEventsRequest(events []EventsDryRunRequestEventsInner, ) *SendEventsRequest`

NewSendEventsRequest instantiates a new SendEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEventsRequestWithDefaults

`func NewSendEventsRequestWithDefaults() *SendEventsRequest`

NewSendEventsRequestWithDefaults instantiates a new SendEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *SendEventsRequest) GetEvents() []EventsDryRunRequestEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SendEventsRequest) GetEventsOk() (*[]EventsDryRunRequestEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SendEventsRequest) SetEvents(v []EventsDryRunRequestEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# WebhookSubscribeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** |  | 
**EventType** | [**NotificationEventType**](NotificationEventType.md) |  | 

## Methods

### NewWebhookSubscribeRequest

`func NewWebhookSubscribeRequest(callbackUrl string, eventType NotificationEventType, ) *WebhookSubscribeRequest`

NewWebhookSubscribeRequest instantiates a new WebhookSubscribeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscribeRequestWithDefaults

`func NewWebhookSubscribeRequestWithDefaults() *WebhookSubscribeRequest`

NewWebhookSubscribeRequestWithDefaults instantiates a new WebhookSubscribeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *WebhookSubscribeRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *WebhookSubscribeRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *WebhookSubscribeRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetEventType

`func (o *WebhookSubscribeRequest) GetEventType() NotificationEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookSubscribeRequest) GetEventTypeOk() (*NotificationEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookSubscribeRequest) SetEventType(v NotificationEventType)`

SetEventType sets EventType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



/*
Vayu API

The Vayu API is a RESTful API that allows you to submit events for processing and storage & manage billing related entities.           The API is secured using the Bearer Authentication scheme with JWT tokens.           To obtain a JWT token, please contact Vayu at team@withvayu.com

API version: 1.0.0
Contact: dev@withvayu.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the SendEventsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendEventsResponse{}

// SendEventsResponse struct for SendEventsResponse
type SendEventsResponse struct {
	// An array of events that were successfully processed and sent to the queue.
	ValidEvents []EventsDryRunRequestEventsInner `json:"validEvents"`
	// An array of events that failed validation and were not sent to the queue. Each object contains the event and the error message.
	InvalidEvents []SendEventsResponseInvalidEventsInner `json:"invalidEvents"`
	AdditionalProperties map[string]interface{}
}

type _SendEventsResponse SendEventsResponse

// NewSendEventsResponse instantiates a new SendEventsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendEventsResponse(validEvents []EventsDryRunRequestEventsInner, invalidEvents []SendEventsResponseInvalidEventsInner) *SendEventsResponse {
	this := SendEventsResponse{}
	this.ValidEvents = validEvents
	this.InvalidEvents = invalidEvents
	return &this
}

// NewSendEventsResponseWithDefaults instantiates a new SendEventsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendEventsResponseWithDefaults() *SendEventsResponse {
	this := SendEventsResponse{}
	return &this
}

// GetValidEvents returns the ValidEvents field value
func (o *SendEventsResponse) GetValidEvents() []EventsDryRunRequestEventsInner {
	if o == nil {
		var ret []EventsDryRunRequestEventsInner
		return ret
	}

	return o.ValidEvents
}

// GetValidEventsOk returns a tuple with the ValidEvents field value
// and a boolean to check if the value has been set.
func (o *SendEventsResponse) GetValidEventsOk() ([]EventsDryRunRequestEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidEvents, true
}

// SetValidEvents sets field value
func (o *SendEventsResponse) SetValidEvents(v []EventsDryRunRequestEventsInner) {
	o.ValidEvents = v
}

// GetInvalidEvents returns the InvalidEvents field value
func (o *SendEventsResponse) GetInvalidEvents() []SendEventsResponseInvalidEventsInner {
	if o == nil {
		var ret []SendEventsResponseInvalidEventsInner
		return ret
	}

	return o.InvalidEvents
}

// GetInvalidEventsOk returns a tuple with the InvalidEvents field value
// and a boolean to check if the value has been set.
func (o *SendEventsResponse) GetInvalidEventsOk() ([]SendEventsResponseInvalidEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvalidEvents, true
}

// SetInvalidEvents sets field value
func (o *SendEventsResponse) SetInvalidEvents(v []SendEventsResponseInvalidEventsInner) {
	o.InvalidEvents = v
}

func (o SendEventsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendEventsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validEvents"] = o.ValidEvents
	toSerialize["invalidEvents"] = o.InvalidEvents

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SendEventsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"validEvents",
		"invalidEvents",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSendEventsResponse := _SendEventsResponse{}

	err = json.Unmarshal(data, &varSendEventsResponse)

	if err != nil {
		return err
	}

	*o = SendEventsResponse(varSendEventsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "validEvents")
		delete(additionalProperties, "invalidEvents")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSendEventsResponse struct {
	value *SendEventsResponse
	isSet bool
}

func (v NullableSendEventsResponse) Get() *SendEventsResponse {
	return v.value
}

func (v *NullableSendEventsResponse) Set(val *SendEventsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendEventsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendEventsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendEventsResponse(val *SendEventsResponse) *NullableSendEventsResponse {
	return &NullableSendEventsResponse{value: val, isSet: true}
}

func (v NullableSendEventsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendEventsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



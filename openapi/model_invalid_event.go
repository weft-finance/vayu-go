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

// checks if the InvalidEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidEvent{}

// InvalidEvent struct for InvalidEvent
type InvalidEvent struct {
	Event Event `json:"event"`
	// The error message indicating the reason the event failed validation.
	Error string `json:"error"`
	AdditionalProperties map[string]interface{}
}

type _InvalidEvent InvalidEvent

// NewInvalidEvent instantiates a new InvalidEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidEvent(event Event, error_ string) *InvalidEvent {
	this := InvalidEvent{}
	this.Event = event
	this.Error = error_
	return &this
}

// NewInvalidEventWithDefaults instantiates a new InvalidEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidEventWithDefaults() *InvalidEvent {
	this := InvalidEvent{}
	return &this
}

// GetEvent returns the Event field value
func (o *InvalidEvent) GetEvent() Event {
	if o == nil {
		var ret Event
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *InvalidEvent) GetEventOk() (*Event, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *InvalidEvent) SetEvent(v Event) {
	o.Event = v
}

// GetError returns the Error field value
func (o *InvalidEvent) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *InvalidEvent) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *InvalidEvent) SetError(v string) {
	o.Error = v
}

func (o InvalidEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvalidEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["error"] = o.Error

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvalidEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"error",
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

	varInvalidEvent := _InvalidEvent{}

	err = json.Unmarshal(data, &varInvalidEvent)

	if err != nil {
		return err
	}

	*o = InvalidEvent(varInvalidEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "event")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvalidEvent struct {
	value *InvalidEvent
	isSet bool
}

func (v NullableInvalidEvent) Get() *InvalidEvent {
	return v.value
}

func (v *NullableInvalidEvent) Set(val *InvalidEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidEvent(val *InvalidEvent) *NullableInvalidEvent {
	return &NullableInvalidEvent{value: val, isSet: true}
}

func (v NullableInvalidEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


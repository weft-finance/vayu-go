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
)

// checks if the EventsDryRunRequestEventsInnerTimestamp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsDryRunRequestEventsInnerTimestamp{}

// EventsDryRunRequestEventsInnerTimestamp struct for EventsDryRunRequestEventsInnerTimestamp
type EventsDryRunRequestEventsInnerTimestamp struct {
}

// NewEventsDryRunRequestEventsInnerTimestamp instantiates a new EventsDryRunRequestEventsInnerTimestamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsDryRunRequestEventsInnerTimestamp() *EventsDryRunRequestEventsInnerTimestamp {
	this := EventsDryRunRequestEventsInnerTimestamp{}
	return &this
}

// NewEventsDryRunRequestEventsInnerTimestampWithDefaults instantiates a new EventsDryRunRequestEventsInnerTimestamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsDryRunRequestEventsInnerTimestampWithDefaults() *EventsDryRunRequestEventsInnerTimestamp {
	this := EventsDryRunRequestEventsInnerTimestamp{}
	return &this
}

func (o EventsDryRunRequestEventsInnerTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsDryRunRequestEventsInnerTimestamp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableEventsDryRunRequestEventsInnerTimestamp struct {
	value *EventsDryRunRequestEventsInnerTimestamp
	isSet bool
}

func (v NullableEventsDryRunRequestEventsInnerTimestamp) Get() *EventsDryRunRequestEventsInnerTimestamp {
	return v.value
}

func (v *NullableEventsDryRunRequestEventsInnerTimestamp) Set(val *EventsDryRunRequestEventsInnerTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsDryRunRequestEventsInnerTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsDryRunRequestEventsInnerTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsDryRunRequestEventsInnerTimestamp(val *EventsDryRunRequestEventsInnerTimestamp) *NullableEventsDryRunRequestEventsInnerTimestamp {
	return &NullableEventsDryRunRequestEventsInnerTimestamp{value: val, isSet: true}
}

func (v NullableEventsDryRunRequestEventsInnerTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsDryRunRequestEventsInnerTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the EventsDryRunResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventsDryRunResponse{}

// EventsDryRunResponse struct for EventsDryRunResponse
type EventsDryRunResponse struct {
	Events []EventsDryRunResponseObject `json:"events"`
	AdditionalProperties map[string]interface{}
}

type _EventsDryRunResponse EventsDryRunResponse

// NewEventsDryRunResponse instantiates a new EventsDryRunResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventsDryRunResponse(events []EventsDryRunResponseObject) *EventsDryRunResponse {
	this := EventsDryRunResponse{}
	this.Events = events
	return &this
}

// NewEventsDryRunResponseWithDefaults instantiates a new EventsDryRunResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventsDryRunResponseWithDefaults() *EventsDryRunResponse {
	this := EventsDryRunResponse{}
	return &this
}

// GetEvents returns the Events field value
func (o *EventsDryRunResponse) GetEvents() []EventsDryRunResponseObject {
	if o == nil {
		var ret []EventsDryRunResponseObject
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EventsDryRunResponse) GetEventsOk() ([]EventsDryRunResponseObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EventsDryRunResponse) SetEvents(v []EventsDryRunResponseObject) {
	o.Events = v
}

func (o EventsDryRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventsDryRunResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventsDryRunResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varEventsDryRunResponse := _EventsDryRunResponse{}

	err = json.Unmarshal(data, &varEventsDryRunResponse)

	if err != nil {
		return err
	}

	*o = EventsDryRunResponse(varEventsDryRunResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "events")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventsDryRunResponse struct {
	value *EventsDryRunResponse
	isSet bool
}

func (v NullableEventsDryRunResponse) Get() *EventsDryRunResponse {
	return v.value
}

func (v *NullableEventsDryRunResponse) Set(val *EventsDryRunResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsDryRunResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsDryRunResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsDryRunResponse(val *EventsDryRunResponse) *NullableEventsDryRunResponse {
	return &NullableEventsDryRunResponse{value: val, isSet: true}
}

func (v NullableEventsDryRunResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsDryRunResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// UnlimitedDuration the model 'UnlimitedDuration'
type UnlimitedDuration float32

// List of UnlimitedDuration
const (
	_MINUS_1 UnlimitedDuration = -1
)

// All allowed values of UnlimitedDuration enum
var AllowedUnlimitedDurationEnumValues = []UnlimitedDuration{
	-1,
}

func (v *UnlimitedDuration) UnmarshalJSON(src []byte) error {
	var value float32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnlimitedDuration(value)
	for _, existing := range AllowedUnlimitedDurationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnlimitedDuration", value)
}

// NewUnlimitedDurationFromValue returns a pointer to a valid UnlimitedDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnlimitedDurationFromValue(v float32) (*UnlimitedDuration, error) {
	ev := UnlimitedDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnlimitedDuration: valid values are %v", v, AllowedUnlimitedDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnlimitedDuration) IsValid() bool {
	for _, existing := range AllowedUnlimitedDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnlimitedDuration value
func (v UnlimitedDuration) Ptr() *UnlimitedDuration {
	return &v
}

type NullableUnlimitedDuration struct {
	value *UnlimitedDuration
	isSet bool
}

func (v NullableUnlimitedDuration) Get() *UnlimitedDuration {
	return v.value
}

func (v *NullableUnlimitedDuration) Set(val *UnlimitedDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableUnlimitedDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableUnlimitedDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnlimitedDuration(val *UnlimitedDuration) *NullableUnlimitedDuration {
	return &NullableUnlimitedDuration{value: val, isSet: true}
}

func (v NullableUnlimitedDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnlimitedDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


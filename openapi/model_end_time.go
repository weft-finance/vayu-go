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
	"time"
	"fmt"
)

// EndTime struct for EndTime
type EndTime struct {
	String *string
	TimeTime *time.Time
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EndTime) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	// try to unmarshal JSON data into TimeTime
	err = json.Unmarshal(data, &dst.TimeTime);
	if err == nil {
		jsonTimeTime, _ := json.Marshal(dst.TimeTime)
		if string(jsonTimeTime) == "{}" { // empty struct
			dst.TimeTime = nil
		} else {
			return nil // data stored in dst.TimeTime, return on the first match
		}
	} else {
		dst.TimeTime = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EndTime)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EndTime) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	if src.TimeTime != nil {
		return json.Marshal(&src.TimeTime)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEndTime struct {
	value *EndTime
	isSet bool
}

func (v NullableEndTime) Get() *EndTime {
	return v.value
}

func (v *NullableEndTime) Set(val *EndTime) {
	v.value = val
	v.isSet = true
}

func (v NullableEndTime) IsSet() bool {
	return v.isSet
}

func (v *NullableEndTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndTime(val *EndTime) *NullableEndTime {
	return &NullableEndTime{value: val, isSet: true}
}

func (v NullableEndTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



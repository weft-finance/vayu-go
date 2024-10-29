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

// checks if the FullDayPeriod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FullDayPeriod{}

// FullDayPeriod The billing cycle of the invoice, consists of a start and end date
type FullDayPeriod struct {
	StartTime time.Time `json:"startTime"`
	EndTime time.Time `json:"endTime"`
	AdditionalProperties map[string]interface{}
}

type _FullDayPeriod FullDayPeriod

// NewFullDayPeriod instantiates a new FullDayPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDayPeriod(startTime time.Time, endTime time.Time) *FullDayPeriod {
	this := FullDayPeriod{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewFullDayPeriodWithDefaults instantiates a new FullDayPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDayPeriodWithDefaults() *FullDayPeriod {
	this := FullDayPeriod{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *FullDayPeriod) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *FullDayPeriod) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *FullDayPeriod) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *FullDayPeriod) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *FullDayPeriod) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *FullDayPeriod) SetEndTime(v time.Time) {
	o.EndTime = v
}

func (o FullDayPeriod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullDayPeriod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FullDayPeriod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startTime",
		"endTime",
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

	varFullDayPeriod := _FullDayPeriod{}

	err = json.Unmarshal(data, &varFullDayPeriod)

	if err != nil {
		return err
	}

	*o = FullDayPeriod(varFullDayPeriod)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullDayPeriod struct {
	value *FullDayPeriod
	isSet bool
}

func (v NullableFullDayPeriod) Get() *FullDayPeriod {
	return v.value
}

func (v *NullableFullDayPeriod) Set(val *FullDayPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDayPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDayPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDayPeriod(val *FullDayPeriod) *NullableFullDayPeriod {
	return &NullableFullDayPeriod{value: val, isSet: true}
}

func (v NullableFullDayPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDayPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


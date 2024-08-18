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
	"bytes"
	"fmt"
)

// checks if the Period type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Period{}

// Period The billing cycle of the invoice, consists of a start and end date
type Period struct {
	StartTime time.Time `json:"startTime"`
	EndTime time.Time `json:"endTime"`
}

type _Period Period

// NewPeriod instantiates a new Period object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriod(startTime time.Time, endTime time.Time) *Period {
	this := Period{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewPeriodWithDefaults instantiates a new Period object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodWithDefaults() *Period {
	this := Period{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *Period) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *Period) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *Period) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *Period) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *Period) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *Period) SetEndTime(v time.Time) {
	o.EndTime = v
}

func (o Period) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Period) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	return toSerialize, nil
}

func (o *Period) UnmarshalJSON(data []byte) (err error) {
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

	varPeriod := _Period{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPeriod)

	if err != nil {
		return err
	}

	*o = Period(varPeriod)

	return err
}

type NullablePeriod struct {
	value *Period
	isSet bool
}

func (v NullablePeriod) Get() *Period {
	return v.value
}

func (v *NullablePeriod) Set(val *Period) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriod) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriod(val *Period) *NullablePeriod {
	return &NullablePeriod{value: val, isSet: true}
}

func (v NullablePeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



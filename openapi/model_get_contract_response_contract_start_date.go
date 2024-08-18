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

// checks if the GetContractResponseContractStartDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContractResponseContractStartDate{}

// GetContractResponseContractStartDate struct for GetContractResponseContractStartDate
type GetContractResponseContractStartDate struct {
}

// NewGetContractResponseContractStartDate instantiates a new GetContractResponseContractStartDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContractResponseContractStartDate() *GetContractResponseContractStartDate {
	this := GetContractResponseContractStartDate{}
	return &this
}

// NewGetContractResponseContractStartDateWithDefaults instantiates a new GetContractResponseContractStartDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContractResponseContractStartDateWithDefaults() *GetContractResponseContractStartDate {
	this := GetContractResponseContractStartDate{}
	return &this
}

func (o GetContractResponseContractStartDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContractResponseContractStartDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableGetContractResponseContractStartDate struct {
	value *GetContractResponseContractStartDate
	isSet bool
}

func (v NullableGetContractResponseContractStartDate) Get() *GetContractResponseContractStartDate {
	return v.value
}

func (v *NullableGetContractResponseContractStartDate) Set(val *GetContractResponseContractStartDate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContractResponseContractStartDate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContractResponseContractStartDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContractResponseContractStartDate(val *GetContractResponseContractStartDate) *NullableGetContractResponseContractStartDate {
	return &NullableGetContractResponseContractStartDate{value: val, isSet: true}
}

func (v NullableGetContractResponseContractStartDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContractResponseContractStartDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



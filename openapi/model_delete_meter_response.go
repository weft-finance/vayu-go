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

// checks if the DeleteMeterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteMeterResponse{}

// DeleteMeterResponse struct for DeleteMeterResponse
type DeleteMeterResponse struct {
	Meter DeleteMeterResponseMeter `json:"meter"`
	AdditionalProperties map[string]interface{}
}

type _DeleteMeterResponse DeleteMeterResponse

// NewDeleteMeterResponse instantiates a new DeleteMeterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMeterResponse(meter DeleteMeterResponseMeter) *DeleteMeterResponse {
	this := DeleteMeterResponse{}
	this.Meter = meter
	return &this
}

// NewDeleteMeterResponseWithDefaults instantiates a new DeleteMeterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMeterResponseWithDefaults() *DeleteMeterResponse {
	this := DeleteMeterResponse{}
	return &this
}

// GetMeter returns the Meter field value
func (o *DeleteMeterResponse) GetMeter() DeleteMeterResponseMeter {
	if o == nil {
		var ret DeleteMeterResponseMeter
		return ret
	}

	return o.Meter
}

// GetMeterOk returns a tuple with the Meter field value
// and a boolean to check if the value has been set.
func (o *DeleteMeterResponse) GetMeterOk() (*DeleteMeterResponseMeter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meter, true
}

// SetMeter sets field value
func (o *DeleteMeterResponse) SetMeter(v DeleteMeterResponseMeter) {
	o.Meter = v
}

func (o DeleteMeterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMeterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meter"] = o.Meter

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteMeterResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meter",
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

	varDeleteMeterResponse := _DeleteMeterResponse{}

	err = json.Unmarshal(data, &varDeleteMeterResponse)

	if err != nil {
		return err
	}

	*o = DeleteMeterResponse(varDeleteMeterResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "meter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteMeterResponse struct {
	value *DeleteMeterResponse
	isSet bool
}

func (v NullableDeleteMeterResponse) Get() *DeleteMeterResponse {
	return v.value
}

func (v *NullableDeleteMeterResponse) Set(val *DeleteMeterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMeterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMeterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMeterResponse(val *DeleteMeterResponse) *NullableDeleteMeterResponse {
	return &NullableDeleteMeterResponse{value: val, isSet: true}
}

func (v NullableDeleteMeterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMeterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



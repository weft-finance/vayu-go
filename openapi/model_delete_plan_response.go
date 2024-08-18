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
	"bytes"
	"fmt"
)

// checks if the DeletePlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePlanResponse{}

// DeletePlanResponse struct for DeletePlanResponse
type DeletePlanResponse struct {
	Plan GetPlanResponsePlan `json:"plan"`
}

type _DeletePlanResponse DeletePlanResponse

// NewDeletePlanResponse instantiates a new DeletePlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePlanResponse(plan GetPlanResponsePlan) *DeletePlanResponse {
	this := DeletePlanResponse{}
	this.Plan = plan
	return &this
}

// NewDeletePlanResponseWithDefaults instantiates a new DeletePlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePlanResponseWithDefaults() *DeletePlanResponse {
	this := DeletePlanResponse{}
	return &this
}

// GetPlan returns the Plan field value
func (o *DeletePlanResponse) GetPlan() GetPlanResponsePlan {
	if o == nil {
		var ret GetPlanResponsePlan
		return ret
	}

	return o.Plan
}

// GetPlanOk returns a tuple with the Plan field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponse) GetPlanOk() (*GetPlanResponsePlan, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plan, true
}

// SetPlan sets field value
func (o *DeletePlanResponse) SetPlan(v GetPlanResponsePlan) {
	o.Plan = v
}

func (o DeletePlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan"] = o.Plan
	return toSerialize, nil
}

func (o *DeletePlanResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plan",
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

	varDeletePlanResponse := _DeletePlanResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeletePlanResponse)

	if err != nil {
		return err
	}

	*o = DeletePlanResponse(varDeletePlanResponse)

	return err
}

type NullableDeletePlanResponse struct {
	value *DeletePlanResponse
	isSet bool
}

func (v NullableDeletePlanResponse) Get() *DeletePlanResponse {
	return v.value
}

func (v *NullableDeletePlanResponse) Set(val *DeletePlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePlanResponse(val *DeletePlanResponse) *NullableDeletePlanResponse {
	return &NullableDeletePlanResponse{value: val, isSet: true}
}

func (v NullableDeletePlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



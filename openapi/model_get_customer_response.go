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

// checks if the GetCustomerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomerResponse{}

// GetCustomerResponse struct for GetCustomerResponse
type GetCustomerResponse struct {
	Customer CreateCustomerResponseCustomer `json:"customer"`
	AdditionalProperties map[string]interface{}
}

type _GetCustomerResponse GetCustomerResponse

// NewGetCustomerResponse instantiates a new GetCustomerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomerResponse(customer CreateCustomerResponseCustomer) *GetCustomerResponse {
	this := GetCustomerResponse{}
	this.Customer = customer
	return &this
}

// NewGetCustomerResponseWithDefaults instantiates a new GetCustomerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomerResponseWithDefaults() *GetCustomerResponse {
	this := GetCustomerResponse{}
	return &this
}

// GetCustomer returns the Customer field value
func (o *GetCustomerResponse) GetCustomer() CreateCustomerResponseCustomer {
	if o == nil {
		var ret CreateCustomerResponseCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *GetCustomerResponse) GetCustomerOk() (*CreateCustomerResponseCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *GetCustomerResponse) SetCustomer(v CreateCustomerResponseCustomer) {
	o.Customer = v
}

func (o GetCustomerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customer"] = o.Customer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCustomerResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customer",
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

	varGetCustomerResponse := _GetCustomerResponse{}

	err = json.Unmarshal(data, &varGetCustomerResponse)

	if err != nil {
		return err
	}

	*o = GetCustomerResponse(varGetCustomerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCustomerResponse struct {
	value *GetCustomerResponse
	isSet bool
}

func (v NullableGetCustomerResponse) Get() *GetCustomerResponse {
	return v.value
}

func (v *NullableGetCustomerResponse) Set(val *GetCustomerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomerResponse(val *GetCustomerResponse) *NullableGetCustomerResponse {
	return &NullableGetCustomerResponse{value: val, isSet: true}
}

func (v NullableGetCustomerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



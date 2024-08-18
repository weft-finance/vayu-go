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

// checks if the DeductCreditsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeductCreditsRequest{}

// DeductCreditsRequest struct for DeductCreditsRequest
type DeductCreditsRequest struct {
	// The amount of credits to be deducted from the user.
	CreditAmount float32 `json:"creditAmount"`
	// The ID of the customer to whom the credits will be deducted from.
	CustomerId string `json:"customerId" validate:"regexp=^[0-9a-fA-F]{24}$"`
}

type _DeductCreditsRequest DeductCreditsRequest

// NewDeductCreditsRequest instantiates a new DeductCreditsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeductCreditsRequest(creditAmount float32, customerId string) *DeductCreditsRequest {
	this := DeductCreditsRequest{}
	this.CreditAmount = creditAmount
	this.CustomerId = customerId
	return &this
}

// NewDeductCreditsRequestWithDefaults instantiates a new DeductCreditsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeductCreditsRequestWithDefaults() *DeductCreditsRequest {
	this := DeductCreditsRequest{}
	return &this
}

// GetCreditAmount returns the CreditAmount field value
func (o *DeductCreditsRequest) GetCreditAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreditAmount
}

// GetCreditAmountOk returns a tuple with the CreditAmount field value
// and a boolean to check if the value has been set.
func (o *DeductCreditsRequest) GetCreditAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditAmount, true
}

// SetCreditAmount sets field value
func (o *DeductCreditsRequest) SetCreditAmount(v float32) {
	o.CreditAmount = v
}

// GetCustomerId returns the CustomerId field value
func (o *DeductCreditsRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *DeductCreditsRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *DeductCreditsRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

func (o DeductCreditsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeductCreditsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creditAmount"] = o.CreditAmount
	toSerialize["customerId"] = o.CustomerId
	return toSerialize, nil
}

func (o *DeductCreditsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"creditAmount",
		"customerId",
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

	varDeductCreditsRequest := _DeductCreditsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeductCreditsRequest)

	if err != nil {
		return err
	}

	*o = DeductCreditsRequest(varDeductCreditsRequest)

	return err
}

type NullableDeductCreditsRequest struct {
	value *DeductCreditsRequest
	isSet bool
}

func (v NullableDeductCreditsRequest) Get() *DeductCreditsRequest {
	return v.value
}

func (v *NullableDeductCreditsRequest) Set(val *DeductCreditsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeductCreditsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeductCreditsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeductCreditsRequest(val *DeductCreditsRequest) *NullableDeductCreditsRequest {
	return &NullableDeductCreditsRequest{value: val, isSet: true}
}

func (v NullableDeductCreditsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeductCreditsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// PaymentTerm flag to indicate if the payment is postpayment or prepayment
type PaymentTerm string

// List of PaymentTerm
const (
	PAYMENTTERM_PREPAYMENT PaymentTerm = "Prepayment"
	PAYMENTTERM_POSTPAYMENT PaymentTerm = "Postpayment"
)

// All allowed values of PaymentTerm enum
var AllowedPaymentTermEnumValues = []PaymentTerm{
	"Prepayment",
	"Postpayment",
}

func (v *PaymentTerm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentTerm(value)
	for _, existing := range AllowedPaymentTermEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentTerm", value)
}

// NewPaymentTermFromValue returns a pointer to a valid PaymentTerm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentTermFromValue(v string) (*PaymentTerm, error) {
	ev := PaymentTerm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentTerm: valid values are %v", v, AllowedPaymentTermEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentTerm) IsValid() bool {
	for _, existing := range AllowedPaymentTermEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentTerm value
func (v PaymentTerm) Ptr() *PaymentTerm {
	return &v
}

type NullablePaymentTerm struct {
	value *PaymentTerm
	isSet bool
}

func (v NullablePaymentTerm) Get() *PaymentTerm {
	return v.value
}

func (v *NullablePaymentTerm) Set(val *PaymentTerm) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTerm) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTerm(val *PaymentTerm) *NullablePaymentTerm {
	return &NullablePaymentTerm{value: val, isSet: true}
}

func (v NullablePaymentTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


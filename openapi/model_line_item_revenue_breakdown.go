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

// checks if the LineItemRevenueBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemRevenueBreakdown{}

// LineItemRevenueBreakdown The revenue breakdown of the line item
type LineItemRevenueBreakdown struct {
	Total float32 `json:"total"`
	Subtotal float32 `json:"subtotal"`
	Overage float32 `json:"overage"`
	Discount float32 `json:"discount"`
	CreditsUsed float32 `json:"creditsUsed"`
	Tax float32 `json:"tax"`
	AdditionalProperties map[string]interface{}
}

type _LineItemRevenueBreakdown LineItemRevenueBreakdown

// NewLineItemRevenueBreakdown instantiates a new LineItemRevenueBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemRevenueBreakdown(total float32, subtotal float32, overage float32, discount float32, creditsUsed float32, tax float32) *LineItemRevenueBreakdown {
	this := LineItemRevenueBreakdown{}
	this.Total = total
	this.Subtotal = subtotal
	this.Overage = overage
	this.Discount = discount
	this.CreditsUsed = creditsUsed
	this.Tax = tax
	return &this
}

// NewLineItemRevenueBreakdownWithDefaults instantiates a new LineItemRevenueBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemRevenueBreakdownWithDefaults() *LineItemRevenueBreakdown {
	this := LineItemRevenueBreakdown{}
	return &this
}

// GetTotal returns the Total field value
func (o *LineItemRevenueBreakdown) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *LineItemRevenueBreakdown) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *LineItemRevenueBreakdown) SetTotal(v float32) {
	o.Total = v
}

// GetSubtotal returns the Subtotal field value
func (o *LineItemRevenueBreakdown) GetSubtotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *LineItemRevenueBreakdown) GetSubtotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *LineItemRevenueBreakdown) SetSubtotal(v float32) {
	o.Subtotal = v
}

// GetOverage returns the Overage field value
func (o *LineItemRevenueBreakdown) GetOverage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Overage
}

// GetOverageOk returns a tuple with the Overage field value
// and a boolean to check if the value has been set.
func (o *LineItemRevenueBreakdown) GetOverageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Overage, true
}

// SetOverage sets field value
func (o *LineItemRevenueBreakdown) SetOverage(v float32) {
	o.Overage = v
}

// GetDiscount returns the Discount field value
func (o *LineItemRevenueBreakdown) GetDiscount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value
// and a boolean to check if the value has been set.
func (o *LineItemRevenueBreakdown) GetDiscountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Discount, true
}

// SetDiscount sets field value
func (o *LineItemRevenueBreakdown) SetDiscount(v float32) {
	o.Discount = v
}

// GetCreditsUsed returns the CreditsUsed field value
func (o *LineItemRevenueBreakdown) GetCreditsUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreditsUsed
}

// GetCreditsUsedOk returns a tuple with the CreditsUsed field value
// and a boolean to check if the value has been set.
func (o *LineItemRevenueBreakdown) GetCreditsUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsUsed, true
}

// SetCreditsUsed sets field value
func (o *LineItemRevenueBreakdown) SetCreditsUsed(v float32) {
	o.CreditsUsed = v
}

// GetTax returns the Tax field value
func (o *LineItemRevenueBreakdown) GetTax() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Tax
}

// GetTaxOk returns a tuple with the Tax field value
// and a boolean to check if the value has been set.
func (o *LineItemRevenueBreakdown) GetTaxOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tax, true
}

// SetTax sets field value
func (o *LineItemRevenueBreakdown) SetTax(v float32) {
	o.Tax = v
}

func (o LineItemRevenueBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemRevenueBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["subtotal"] = o.Subtotal
	toSerialize["overage"] = o.Overage
	toSerialize["discount"] = o.Discount
	toSerialize["creditsUsed"] = o.CreditsUsed
	toSerialize["tax"] = o.Tax

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LineItemRevenueBreakdown) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"subtotal",
		"overage",
		"discount",
		"creditsUsed",
		"tax",
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

	varLineItemRevenueBreakdown := _LineItemRevenueBreakdown{}

	err = json.Unmarshal(data, &varLineItemRevenueBreakdown)

	if err != nil {
		return err
	}

	*o = LineItemRevenueBreakdown(varLineItemRevenueBreakdown)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "subtotal")
		delete(additionalProperties, "overage")
		delete(additionalProperties, "discount")
		delete(additionalProperties, "creditsUsed")
		delete(additionalProperties, "tax")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLineItemRevenueBreakdown struct {
	value *LineItemRevenueBreakdown
	isSet bool
}

func (v NullableLineItemRevenueBreakdown) Get() *LineItemRevenueBreakdown {
	return v.value
}

func (v *NullableLineItemRevenueBreakdown) Set(val *LineItemRevenueBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemRevenueBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemRevenueBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemRevenueBreakdown(val *LineItemRevenueBreakdown) *NullableLineItemRevenueBreakdown {
	return &NullableLineItemRevenueBreakdown{value: val, isSet: true}
}

func (v NullableLineItemRevenueBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemRevenueBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



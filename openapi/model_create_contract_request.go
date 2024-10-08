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

// checks if the CreateContractRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContractRequest{}

// CreateContractRequest struct for CreateContractRequest
type CreateContractRequest struct {
	// The start date of the contract
	StartDate time.Time `json:"startDate"`
	// The end date of the contract
	EndDate *time.Time `json:"endDate,omitempty"`
	// The id of the customer that the contract is associated with
	CustomerId string `json:"customerId" validate:"regexp=^[0-9a-fA-F]{24}$"`
	// The id of the plan that the contract is associated with
	PlanId string `json:"planId" validate:"regexp=^[0-9a-fA-F]{24}$"`
	AdditionalProperties map[string]interface{}
}

type _CreateContractRequest CreateContractRequest

// NewCreateContractRequest instantiates a new CreateContractRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContractRequest(startDate time.Time, customerId string, planId string) *CreateContractRequest {
	this := CreateContractRequest{}
	this.StartDate = startDate
	this.CustomerId = customerId
	this.PlanId = planId
	return &this
}

// NewCreateContractRequestWithDefaults instantiates a new CreateContractRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContractRequestWithDefaults() *CreateContractRequest {
	this := CreateContractRequest{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *CreateContractRequest) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CreateContractRequest) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreateContractRequest) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreateContractRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *CreateContractRequest) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetCustomerId returns the CustomerId field value
func (o *CreateContractRequest) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CreateContractRequest) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetPlanId returns the PlanId field value
func (o *CreateContractRequest) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *CreateContractRequest) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *CreateContractRequest) SetPlanId(v string) {
	o.PlanId = v
}

func (o CreateContractRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContractRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startDate"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["customerId"] = o.CustomerId
	toSerialize["planId"] = o.PlanId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateContractRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startDate",
		"customerId",
		"planId",
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

	varCreateContractRequest := _CreateContractRequest{}

	err = json.Unmarshal(data, &varCreateContractRequest)

	if err != nil {
		return err
	}

	*o = CreateContractRequest(varCreateContractRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "planId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateContractRequest struct {
	value *CreateContractRequest
	isSet bool
}

func (v NullableCreateContractRequest) Get() *CreateContractRequest {
	return v.value
}

func (v *NullableCreateContractRequest) Set(val *CreateContractRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContractRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContractRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContractRequest(val *CreateContractRequest) *NullableCreateContractRequest {
	return &NullableCreateContractRequest{value: val, isSet: true}
}

func (v NullableCreateContractRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContractRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



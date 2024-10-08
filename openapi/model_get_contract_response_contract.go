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

// checks if the GetContractResponseContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContractResponseContract{}

// GetContractResponseContract struct for GetContractResponseContract
type GetContractResponseContract struct {
	// The start date of the contract
	StartDate time.Time `json:"startDate"`
	// The end date of the contract
	EndDate *time.Time `json:"endDate,omitempty"`
	// The id of the customer that the contract is associated with
	CustomerId string `json:"customerId" validate:"regexp=^[0-9a-fA-F]{24}$"`
	// The id of the plan that the contract is associated with
	PlanId string `json:"planId" validate:"regexp=^[0-9a-fA-F]{24}$"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _GetContractResponseContract GetContractResponseContract

// NewGetContractResponseContract instantiates a new GetContractResponseContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContractResponseContract(startDate time.Time, customerId string, planId string, id string, createdAt time.Time, updatedAt time.Time) *GetContractResponseContract {
	this := GetContractResponseContract{}
	this.StartDate = startDate
	this.CustomerId = customerId
	this.PlanId = planId
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewGetContractResponseContractWithDefaults instantiates a new GetContractResponseContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContractResponseContractWithDefaults() *GetContractResponseContract {
	this := GetContractResponseContract{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *GetContractResponseContract) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *GetContractResponseContract) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *GetContractResponseContract) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetContractResponseContract) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContractResponseContract) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetContractResponseContract) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *GetContractResponseContract) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetCustomerId returns the CustomerId field value
func (o *GetContractResponseContract) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *GetContractResponseContract) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *GetContractResponseContract) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetPlanId returns the PlanId field value
func (o *GetContractResponseContract) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *GetContractResponseContract) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *GetContractResponseContract) SetPlanId(v string) {
	o.PlanId = v
}

// GetId returns the Id field value
func (o *GetContractResponseContract) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetContractResponseContract) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetContractResponseContract) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetContractResponseContract) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetContractResponseContract) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetContractResponseContract) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetContractResponseContract) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetContractResponseContract) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetContractResponseContract) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o GetContractResponseContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContractResponseContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startDate"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["customerId"] = o.CustomerId
	toSerialize["planId"] = o.PlanId
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetContractResponseContract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startDate",
		"customerId",
		"planId",
		"id",
		"createdAt",
		"updatedAt",
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

	varGetContractResponseContract := _GetContractResponseContract{}

	err = json.Unmarshal(data, &varGetContractResponseContract)

	if err != nil {
		return err
	}

	*o = GetContractResponseContract(varGetContractResponseContract)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "planId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetContractResponseContract struct {
	value *GetContractResponseContract
	isSet bool
}

func (v NullableGetContractResponseContract) Get() *GetContractResponseContract {
	return v.value
}

func (v *NullableGetContractResponseContract) Set(val *GetContractResponseContract) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContractResponseContract) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContractResponseContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContractResponseContract(val *GetContractResponseContract) *NullableGetContractResponseContract {
	return &NullableGetContractResponseContract{value: val, isSet: true}
}

func (v NullableGetContractResponseContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContractResponseContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the DeletePlanResponsePlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePlanResponsePlan{}

// DeletePlanResponsePlan struct for DeletePlanResponsePlan
type DeletePlanResponsePlan struct {
	// The name of the plan
	Name string `json:"name"`
	Status PlanStatus `json:"status"`
	BillingData PlanBillingData `json:"billingData"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
	AdditionalProperties map[string]interface{}
}

type _DeletePlanResponsePlan DeletePlanResponsePlan

// NewDeletePlanResponsePlan instantiates a new DeletePlanResponsePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePlanResponsePlan(name string, status PlanStatus, billingData PlanBillingData, id string, createdAt time.Time, updatedAt time.Time, deletedAt string) *DeletePlanResponsePlan {
	this := DeletePlanResponsePlan{}
	this.Name = name
	this.Status = status
	this.BillingData = billingData
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DeletedAt = deletedAt
	return &this
}

// NewDeletePlanResponsePlanWithDefaults instantiates a new DeletePlanResponsePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePlanResponsePlanWithDefaults() *DeletePlanResponsePlan {
	this := DeletePlanResponsePlan{}
	return &this
}

// GetName returns the Name field value
func (o *DeletePlanResponsePlan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponsePlan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeletePlanResponsePlan) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *DeletePlanResponsePlan) GetStatus() PlanStatus {
	if o == nil {
		var ret PlanStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponsePlan) GetStatusOk() (*PlanStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeletePlanResponsePlan) SetStatus(v PlanStatus) {
	o.Status = v
}

// GetBillingData returns the BillingData field value
func (o *DeletePlanResponsePlan) GetBillingData() PlanBillingData {
	if o == nil {
		var ret PlanBillingData
		return ret
	}

	return o.BillingData
}

// GetBillingDataOk returns a tuple with the BillingData field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponsePlan) GetBillingDataOk() (*PlanBillingData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingData, true
}

// SetBillingData sets field value
func (o *DeletePlanResponsePlan) SetBillingData(v PlanBillingData) {
	o.BillingData = v
}

// GetId returns the Id field value
func (o *DeletePlanResponsePlan) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponsePlan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeletePlanResponsePlan) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeletePlanResponsePlan) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponsePlan) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeletePlanResponsePlan) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DeletePlanResponsePlan) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponsePlan) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DeletePlanResponsePlan) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *DeletePlanResponsePlan) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *DeletePlanResponsePlan) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *DeletePlanResponsePlan) SetDeletedAt(v string) {
	o.DeletedAt = v
}

func (o DeletePlanResponsePlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePlanResponsePlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["billingData"] = o.BillingData
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["deletedAt"] = o.DeletedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeletePlanResponsePlan) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"status",
		"billingData",
		"id",
		"createdAt",
		"updatedAt",
		"deletedAt",
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

	varDeletePlanResponsePlan := _DeletePlanResponsePlan{}

	err = json.Unmarshal(data, &varDeletePlanResponsePlan)

	if err != nil {
		return err
	}

	*o = DeletePlanResponsePlan(varDeletePlanResponsePlan)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "billingData")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "deletedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeletePlanResponsePlan struct {
	value *DeletePlanResponsePlan
	isSet bool
}

func (v NullableDeletePlanResponsePlan) Get() *DeletePlanResponsePlan {
	return v.value
}

func (v *NullableDeletePlanResponsePlan) Set(val *DeletePlanResponsePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePlanResponsePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePlanResponsePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePlanResponsePlan(val *DeletePlanResponsePlan) *NullableDeletePlanResponsePlan {
	return &NullableDeletePlanResponsePlan{value: val, isSet: true}
}

func (v NullableDeletePlanResponsePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePlanResponsePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the CreateCustomerResponseCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCustomerResponseCustomer{}

// CreateCustomerResponseCustomer struct for CreateCustomerResponseCustomer
type CreateCustomerResponseCustomer struct {
	// The name of the customer
	Name string `json:"name"`
	// The alias of the customer used to match events to the customer.
	Alias *string `json:"alias,omitempty"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
	AdditionalProperties map[string]interface{}
}

type _CreateCustomerResponseCustomer CreateCustomerResponseCustomer

// NewCreateCustomerResponseCustomer instantiates a new CreateCustomerResponseCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCustomerResponseCustomer(name string, id string, createdAt time.Time, updatedAt time.Time, deletedAt string) *CreateCustomerResponseCustomer {
	this := CreateCustomerResponseCustomer{}
	this.Name = name
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DeletedAt = deletedAt
	return &this
}

// NewCreateCustomerResponseCustomerWithDefaults instantiates a new CreateCustomerResponseCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCustomerResponseCustomerWithDefaults() *CreateCustomerResponseCustomer {
	this := CreateCustomerResponseCustomer{}
	return &this
}

// GetName returns the Name field value
func (o *CreateCustomerResponseCustomer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerResponseCustomer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateCustomerResponseCustomer) SetName(v string) {
	o.Name = v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *CreateCustomerResponseCustomer) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCustomerResponseCustomer) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *CreateCustomerResponseCustomer) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *CreateCustomerResponseCustomer) SetAlias(v string) {
	o.Alias = &v
}

// GetId returns the Id field value
func (o *CreateCustomerResponseCustomer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerResponseCustomer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateCustomerResponseCustomer) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CreateCustomerResponseCustomer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerResponseCustomer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CreateCustomerResponseCustomer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CreateCustomerResponseCustomer) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerResponseCustomer) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CreateCustomerResponseCustomer) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *CreateCustomerResponseCustomer) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *CreateCustomerResponseCustomer) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *CreateCustomerResponseCustomer) SetDeletedAt(v string) {
	o.DeletedAt = v
}

func (o CreateCustomerResponseCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCustomerResponseCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["deletedAt"] = o.DeletedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateCustomerResponseCustomer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCreateCustomerResponseCustomer := _CreateCustomerResponseCustomer{}

	err = json.Unmarshal(data, &varCreateCustomerResponseCustomer)

	if err != nil {
		return err
	}

	*o = CreateCustomerResponseCustomer(varCreateCustomerResponseCustomer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "alias")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "deletedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateCustomerResponseCustomer struct {
	value *CreateCustomerResponseCustomer
	isSet bool
}

func (v NullableCreateCustomerResponseCustomer) Get() *CreateCustomerResponseCustomer {
	return v.value
}

func (v *NullableCreateCustomerResponseCustomer) Set(val *CreateCustomerResponseCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCustomerResponseCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCustomerResponseCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCustomerResponseCustomer(val *CreateCustomerResponseCustomer) *NullableCreateCustomerResponseCustomer {
	return &NullableCreateCustomerResponseCustomer{value: val, isSet: true}
}

func (v NullableCreateCustomerResponseCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCustomerResponseCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ListCustomersResponseCustomersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCustomersResponseCustomersInner{}

// ListCustomersResponseCustomersInner struct for ListCustomersResponseCustomersInner
type ListCustomersResponseCustomersInner struct {
	// The name of the customer
	Name string `json:"name"`
	// The alias of the customer used to match events to the customer.
	Alias *string `json:"alias,omitempty"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	AdditionalProperties map[string]interface{}
}

type _ListCustomersResponseCustomersInner ListCustomersResponseCustomersInner

// NewListCustomersResponseCustomersInner instantiates a new ListCustomersResponseCustomersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCustomersResponseCustomersInner(name string, id string, createdAt time.Time, updatedAt time.Time) *ListCustomersResponseCustomersInner {
	this := ListCustomersResponseCustomersInner{}
	this.Name = name
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewListCustomersResponseCustomersInnerWithDefaults instantiates a new ListCustomersResponseCustomersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCustomersResponseCustomersInnerWithDefaults() *ListCustomersResponseCustomersInner {
	this := ListCustomersResponseCustomersInner{}
	return &this
}

// GetName returns the Name field value
func (o *ListCustomersResponseCustomersInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponseCustomersInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ListCustomersResponseCustomersInner) SetName(v string) {
	o.Name = v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ListCustomersResponseCustomersInner) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomersResponseCustomersInner) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ListCustomersResponseCustomersInner) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ListCustomersResponseCustomersInner) SetAlias(v string) {
	o.Alias = &v
}

// GetId returns the Id field value
func (o *ListCustomersResponseCustomersInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponseCustomersInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListCustomersResponseCustomersInner) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ListCustomersResponseCustomersInner) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponseCustomersInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ListCustomersResponseCustomersInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ListCustomersResponseCustomersInner) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponseCustomersInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ListCustomersResponseCustomersInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o ListCustomersResponseCustomersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCustomersResponseCustomersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCustomersResponseCustomersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varListCustomersResponseCustomersInner := _ListCustomersResponseCustomersInner{}

	err = json.Unmarshal(data, &varListCustomersResponseCustomersInner)

	if err != nil {
		return err
	}

	*o = ListCustomersResponseCustomersInner(varListCustomersResponseCustomersInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "alias")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCustomersResponseCustomersInner struct {
	value *ListCustomersResponseCustomersInner
	isSet bool
}

func (v NullableListCustomersResponseCustomersInner) Get() *ListCustomersResponseCustomersInner {
	return v.value
}

func (v *NullableListCustomersResponseCustomersInner) Set(val *ListCustomersResponseCustomersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListCustomersResponseCustomersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListCustomersResponseCustomersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCustomersResponseCustomersInner(val *ListCustomersResponseCustomersInner) *NullableListCustomersResponseCustomersInner {
	return &NullableListCustomersResponseCustomersInner{value: val, isSet: true}
}

func (v NullableListCustomersResponseCustomersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCustomersResponseCustomersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



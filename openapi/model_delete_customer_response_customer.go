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

// checks if the DeleteCustomerResponseCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCustomerResponseCustomer{}

// DeleteCustomerResponseCustomer struct for DeleteCustomerResponseCustomer
type DeleteCustomerResponseCustomer struct {
	// The name of the customer
	Name string `json:"name"`
	// The aliases of the customer used to match events to the customer.
	Aliases []string `json:"aliases,omitempty"`
	Address *Address `json:"address,omitempty"`
	// The contacts of the customer. Contact marked as primary is the target for invoice sharing.
	Contacts []Contact `json:"contacts,omitempty"`
	// The external ID of the customer
	ExternalId *string `json:"externalId,omitempty"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
	AdditionalProperties map[string]interface{}
}

type _DeleteCustomerResponseCustomer DeleteCustomerResponseCustomer

// NewDeleteCustomerResponseCustomer instantiates a new DeleteCustomerResponseCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCustomerResponseCustomer(name string, id string, createdAt time.Time, updatedAt time.Time, deletedAt string) *DeleteCustomerResponseCustomer {
	this := DeleteCustomerResponseCustomer{}
	this.Name = name
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DeletedAt = deletedAt
	return &this
}

// NewDeleteCustomerResponseCustomerWithDefaults instantiates a new DeleteCustomerResponseCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCustomerResponseCustomerWithDefaults() *DeleteCustomerResponseCustomer {
	this := DeleteCustomerResponseCustomer{}
	return &this
}

// GetName returns the Name field value
func (o *DeleteCustomerResponseCustomer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeleteCustomerResponseCustomer) SetName(v string) {
	o.Name = v
}

// GetAliases returns the Aliases field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteCustomerResponseCustomer) GetAliases() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteCustomerResponseCustomer) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *DeleteCustomerResponseCustomer) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *DeleteCustomerResponseCustomer) SetAliases(v []string) {
	o.Aliases = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DeleteCustomerResponseCustomer) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DeleteCustomerResponseCustomer) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *DeleteCustomerResponseCustomer) SetAddress(v Address) {
	o.Address = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *DeleteCustomerResponseCustomer) GetContacts() []Contact {
	if o == nil || IsNil(o.Contacts) {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *DeleteCustomerResponseCustomer) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *DeleteCustomerResponseCustomer) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *DeleteCustomerResponseCustomer) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *DeleteCustomerResponseCustomer) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *DeleteCustomerResponseCustomer) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value
func (o *DeleteCustomerResponseCustomer) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteCustomerResponseCustomer) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeleteCustomerResponseCustomer) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeleteCustomerResponseCustomer) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DeleteCustomerResponseCustomer) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DeleteCustomerResponseCustomer) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *DeleteCustomerResponseCustomer) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomerResponseCustomer) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *DeleteCustomerResponseCustomer) SetDeletedAt(v string) {
	o.DeletedAt = v
}

func (o DeleteCustomerResponseCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCustomerResponseCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Aliases != nil {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
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

func (o *DeleteCustomerResponseCustomer) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteCustomerResponseCustomer := _DeleteCustomerResponseCustomer{}

	err = json.Unmarshal(data, &varDeleteCustomerResponseCustomer)

	if err != nil {
		return err
	}

	*o = DeleteCustomerResponseCustomer(varDeleteCustomerResponseCustomer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "aliases")
		delete(additionalProperties, "address")
		delete(additionalProperties, "contacts")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "deletedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteCustomerResponseCustomer struct {
	value *DeleteCustomerResponseCustomer
	isSet bool
}

func (v NullableDeleteCustomerResponseCustomer) Get() *DeleteCustomerResponseCustomer {
	return v.value
}

func (v *NullableDeleteCustomerResponseCustomer) Set(val *DeleteCustomerResponseCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCustomerResponseCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCustomerResponseCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCustomerResponseCustomer(val *DeleteCustomerResponseCustomer) *NullableDeleteCustomerResponseCustomer {
	return &NullableDeleteCustomerResponseCustomer{value: val, isSet: true}
}

func (v NullableDeleteCustomerResponseCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCustomerResponseCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



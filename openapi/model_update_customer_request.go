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
)

// checks if the UpdateCustomerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCustomerRequest{}

// UpdateCustomerRequest struct for UpdateCustomerRequest
type UpdateCustomerRequest struct {
	// The name of the customer
	Name *string `json:"name,omitempty"`
	// The external ID of the customer
	ExternalId *string `json:"externalId,omitempty"`
	// The aliases of the customer used to match events to the customer.
	Aliases []string `json:"aliases,omitempty"`
	Address *Address `json:"address,omitempty"`
}

// NewUpdateCustomerRequest instantiates a new UpdateCustomerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCustomerRequest() *UpdateCustomerRequest {
	this := UpdateCustomerRequest{}
	return &this
}

// NewUpdateCustomerRequestWithDefaults instantiates a new UpdateCustomerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCustomerRequestWithDefaults() *UpdateCustomerRequest {
	this := UpdateCustomerRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCustomerRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomerRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateCustomerRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateCustomerRequest) SetName(v string) {
	o.Name = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UpdateCustomerRequest) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomerRequest) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UpdateCustomerRequest) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UpdateCustomerRequest) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetAliases returns the Aliases field value if set, zero value otherwise.
func (o *UpdateCustomerRequest) GetAliases() []string {
	if o == nil || IsNil(o.Aliases) {
		var ret []string
		return ret
	}
	return o.Aliases
}

// GetAliasesOk returns a tuple with the Aliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomerRequest) GetAliasesOk() ([]string, bool) {
	if o == nil || IsNil(o.Aliases) {
		return nil, false
	}
	return o.Aliases, true
}

// HasAliases returns a boolean if a field has been set.
func (o *UpdateCustomerRequest) HasAliases() bool {
	if o != nil && !IsNil(o.Aliases) {
		return true
	}

	return false
}

// SetAliases gets a reference to the given []string and assigns it to the Aliases field.
func (o *UpdateCustomerRequest) SetAliases(v []string) {
	o.Aliases = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateCustomerRequest) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCustomerRequest) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateCustomerRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *UpdateCustomerRequest) SetAddress(v Address) {
	o.Address = &v
}

func (o UpdateCustomerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCustomerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Aliases) {
		toSerialize["aliases"] = o.Aliases
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableUpdateCustomerRequest struct {
	value *UpdateCustomerRequest
	isSet bool
}

func (v NullableUpdateCustomerRequest) Get() *UpdateCustomerRequest {
	return v.value
}

func (v *NullableUpdateCustomerRequest) Set(val *UpdateCustomerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCustomerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCustomerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCustomerRequest(val *UpdateCustomerRequest) *NullableUpdateCustomerRequest {
	return &NullableUpdateCustomerRequest{value: val, isSet: true}
}

func (v NullableUpdateCustomerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCustomerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



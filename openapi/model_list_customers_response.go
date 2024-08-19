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

// checks if the ListCustomersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCustomersResponse{}

// ListCustomersResponse struct for ListCustomersResponse
type ListCustomersResponse struct {
	Customers []ListCustomersResponseCustomersInner `json:"customers"`
	Total float32 `json:"total"`
	HasMore bool `json:"hasMore"`
	NextCursor *string `json:"nextCursor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListCustomersResponse ListCustomersResponse

// NewListCustomersResponse instantiates a new ListCustomersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCustomersResponse(customers []ListCustomersResponseCustomersInner, total float32, hasMore bool) *ListCustomersResponse {
	this := ListCustomersResponse{}
	this.Customers = customers
	this.Total = total
	this.HasMore = hasMore
	return &this
}

// NewListCustomersResponseWithDefaults instantiates a new ListCustomersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCustomersResponseWithDefaults() *ListCustomersResponse {
	this := ListCustomersResponse{}
	return &this
}

// GetCustomers returns the Customers field value
func (o *ListCustomersResponse) GetCustomers() []ListCustomersResponseCustomersInner {
	if o == nil {
		var ret []ListCustomersResponseCustomersInner
		return ret
	}

	return o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponse) GetCustomersOk() ([]ListCustomersResponseCustomersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Customers, true
}

// SetCustomers sets field value
func (o *ListCustomersResponse) SetCustomers(v []ListCustomersResponseCustomersInner) {
	o.Customers = v
}

// GetTotal returns the Total field value
func (o *ListCustomersResponse) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponse) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListCustomersResponse) SetTotal(v float32) {
	o.Total = v
}

// GetHasMore returns the HasMore field value
func (o *ListCustomersResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ListCustomersResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ListCustomersResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *ListCustomersResponse) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor) {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomersResponse) GetNextCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ListCustomersResponse) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *ListCustomersResponse) SetNextCursor(v string) {
	o.NextCursor = &v
}

func (o ListCustomersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCustomersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customers"] = o.Customers
	toSerialize["total"] = o.Total
	toSerialize["hasMore"] = o.HasMore
	if !IsNil(o.NextCursor) {
		toSerialize["nextCursor"] = o.NextCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListCustomersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customers",
		"total",
		"hasMore",
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

	varListCustomersResponse := _ListCustomersResponse{}

	err = json.Unmarshal(data, &varListCustomersResponse)

	if err != nil {
		return err
	}

	*o = ListCustomersResponse(varListCustomersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customers")
		delete(additionalProperties, "total")
		delete(additionalProperties, "hasMore")
		delete(additionalProperties, "nextCursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListCustomersResponse struct {
	value *ListCustomersResponse
	isSet bool
}

func (v NullableListCustomersResponse) Get() *ListCustomersResponse {
	return v.value
}

func (v *NullableListCustomersResponse) Set(val *ListCustomersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCustomersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCustomersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCustomersResponse(val *ListCustomersResponse) *NullableListCustomersResponse {
	return &NullableListCustomersResponse{value: val, isSet: true}
}

func (v NullableListCustomersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCustomersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



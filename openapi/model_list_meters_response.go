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
	"bytes"
	"fmt"
)

// checks if the ListMetersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMetersResponse{}

// ListMetersResponse struct for ListMetersResponse
type ListMetersResponse struct {
	Meters []ListMetersResponseMetersInner `json:"meters"`
	Total float32 `json:"total"`
	HasMore bool `json:"hasMore"`
	NextCursor *string `json:"nextCursor,omitempty"`
}

type _ListMetersResponse ListMetersResponse

// NewListMetersResponse instantiates a new ListMetersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMetersResponse(meters []ListMetersResponseMetersInner, total float32, hasMore bool) *ListMetersResponse {
	this := ListMetersResponse{}
	this.Meters = meters
	this.Total = total
	this.HasMore = hasMore
	return &this
}

// NewListMetersResponseWithDefaults instantiates a new ListMetersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMetersResponseWithDefaults() *ListMetersResponse {
	this := ListMetersResponse{}
	return &this
}

// GetMeters returns the Meters field value
func (o *ListMetersResponse) GetMeters() []ListMetersResponseMetersInner {
	if o == nil {
		var ret []ListMetersResponseMetersInner
		return ret
	}

	return o.Meters
}

// GetMetersOk returns a tuple with the Meters field value
// and a boolean to check if the value has been set.
func (o *ListMetersResponse) GetMetersOk() ([]ListMetersResponseMetersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meters, true
}

// SetMeters sets field value
func (o *ListMetersResponse) SetMeters(v []ListMetersResponseMetersInner) {
	o.Meters = v
}

// GetTotal returns the Total field value
func (o *ListMetersResponse) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ListMetersResponse) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ListMetersResponse) SetTotal(v float32) {
	o.Total = v
}

// GetHasMore returns the HasMore field value
func (o *ListMetersResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *ListMetersResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *ListMetersResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *ListMetersResponse) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor) {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMetersResponse) GetNextCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ListMetersResponse) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *ListMetersResponse) SetNextCursor(v string) {
	o.NextCursor = &v
}

func (o ListMetersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMetersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["meters"] = o.Meters
	toSerialize["total"] = o.Total
	toSerialize["hasMore"] = o.HasMore
	if !IsNil(o.NextCursor) {
		toSerialize["nextCursor"] = o.NextCursor
	}
	return toSerialize, nil
}

func (o *ListMetersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"meters",
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

	varListMetersResponse := _ListMetersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListMetersResponse)

	if err != nil {
		return err
	}

	*o = ListMetersResponse(varListMetersResponse)

	return err
}

type NullableListMetersResponse struct {
	value *ListMetersResponse
	isSet bool
}

func (v NullableListMetersResponse) Get() *ListMetersResponse {
	return v.value
}

func (v *NullableListMetersResponse) Set(val *ListMetersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListMetersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListMetersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMetersResponse(val *ListMetersResponse) *NullableListMetersResponse {
	return &NullableListMetersResponse{value: val, isSet: true}
}

func (v NullableListMetersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMetersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



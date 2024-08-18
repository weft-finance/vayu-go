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
	"bytes"
	"fmt"
)

// checks if the DeleteEventByRefIdResponseEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteEventByRefIdResponseEvent{}

// DeleteEventByRefIdResponseEvent The event matching the provided refId
type DeleteEventByRefIdResponseEvent struct {
	// The distinctive label assigned to an event, serving as a critical identifier for categorizing and pricing events within the system's backend infrastructure.
	Name string `json:"name"`
	// The temporal marker denoting the exact moment of event occurrence. The timestamp is formatted as an ISO 8601 string and is expressed in Coordinated Universal Time (UTC). i.e. YYYY-MM-DDTHH:MM:SS.SSSZ
	Timestamp time.Time `json:"timestamp"`
	// A pseudonymous or otherwise obfuscated identifier uniquely assigned to each customer.
	CustomerAlias string `json:"customerAlias"`
	// A universally unique identifier (UUID) or other form of high-entropy string serving as an immutable reference for each event entry.
	Ref string `json:"ref"`
	// A schema-less JSON object encapsulating miscellaneous attributes and metrics associated with the event.
	Data map[string]interface{} `json:"data,omitempty"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}

type _DeleteEventByRefIdResponseEvent DeleteEventByRefIdResponseEvent

// NewDeleteEventByRefIdResponseEvent instantiates a new DeleteEventByRefIdResponseEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteEventByRefIdResponseEvent(name string, timestamp time.Time, customerAlias string, ref string, id string, createdAt time.Time, updatedAt time.Time, deletedAt string) *DeleteEventByRefIdResponseEvent {
	this := DeleteEventByRefIdResponseEvent{}
	this.Name = name
	this.Timestamp = timestamp
	this.CustomerAlias = customerAlias
	this.Ref = ref
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DeletedAt = deletedAt
	return &this
}

// NewDeleteEventByRefIdResponseEventWithDefaults instantiates a new DeleteEventByRefIdResponseEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteEventByRefIdResponseEventWithDefaults() *DeleteEventByRefIdResponseEvent {
	this := DeleteEventByRefIdResponseEvent{}
	return &this
}

// GetName returns the Name field value
func (o *DeleteEventByRefIdResponseEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeleteEventByRefIdResponseEvent) SetName(v string) {
	o.Name = v
}

// GetTimestamp returns the Timestamp field value
func (o *DeleteEventByRefIdResponseEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *DeleteEventByRefIdResponseEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetCustomerAlias returns the CustomerAlias field value
func (o *DeleteEventByRefIdResponseEvent) GetCustomerAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerAlias
}

// GetCustomerAliasOk returns a tuple with the CustomerAlias field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetCustomerAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAlias, true
}

// SetCustomerAlias sets field value
func (o *DeleteEventByRefIdResponseEvent) SetCustomerAlias(v string) {
	o.CustomerAlias = v
}

// GetRef returns the Ref field value
func (o *DeleteEventByRefIdResponseEvent) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *DeleteEventByRefIdResponseEvent) SetRef(v string) {
	o.Ref = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeleteEventByRefIdResponseEvent) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeleteEventByRefIdResponseEvent) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteEventByRefIdResponseEvent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *DeleteEventByRefIdResponseEvent) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetId returns the Id field value
func (o *DeleteEventByRefIdResponseEvent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeleteEventByRefIdResponseEvent) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DeleteEventByRefIdResponseEvent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DeleteEventByRefIdResponseEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DeleteEventByRefIdResponseEvent) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DeleteEventByRefIdResponseEvent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *DeleteEventByRefIdResponseEvent) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *DeleteEventByRefIdResponseEvent) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *DeleteEventByRefIdResponseEvent) SetDeletedAt(v string) {
	o.DeletedAt = v
}

func (o DeleteEventByRefIdResponseEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteEventByRefIdResponseEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["customerAlias"] = o.CustomerAlias
	toSerialize["ref"] = o.Ref
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["deletedAt"] = o.DeletedAt
	return toSerialize, nil
}

func (o *DeleteEventByRefIdResponseEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"timestamp",
		"customerAlias",
		"ref",
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

	varDeleteEventByRefIdResponseEvent := _DeleteEventByRefIdResponseEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteEventByRefIdResponseEvent)

	if err != nil {
		return err
	}

	*o = DeleteEventByRefIdResponseEvent(varDeleteEventByRefIdResponseEvent)

	return err
}

type NullableDeleteEventByRefIdResponseEvent struct {
	value *DeleteEventByRefIdResponseEvent
	isSet bool
}

func (v NullableDeleteEventByRefIdResponseEvent) Get() *DeleteEventByRefIdResponseEvent {
	return v.value
}

func (v *NullableDeleteEventByRefIdResponseEvent) Set(val *DeleteEventByRefIdResponseEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteEventByRefIdResponseEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteEventByRefIdResponseEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteEventByRefIdResponseEvent(val *DeleteEventByRefIdResponseEvent) *NullableDeleteEventByRefIdResponseEvent {
	return &NullableDeleteEventByRefIdResponseEvent{value: val, isSet: true}
}

func (v NullableDeleteEventByRefIdResponseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteEventByRefIdResponseEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



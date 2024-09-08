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

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event struct for Event
type Event struct {
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
	AdditionalProperties map[string]interface{}
}

type _Event Event

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent(name string, timestamp time.Time, customerAlias string, ref string) *Event {
	this := Event{}
	this.Name = name
	this.Timestamp = timestamp
	this.CustomerAlias = customerAlias
	this.Ref = ref
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetName returns the Name field value
func (o *Event) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Event) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Event) SetName(v string) {
	o.Name = v
}

// GetTimestamp returns the Timestamp field value
func (o *Event) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Event) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Event) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetCustomerAlias returns the CustomerAlias field value
func (o *Event) GetCustomerAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerAlias
}

// GetCustomerAliasOk returns a tuple with the CustomerAlias field value
// and a boolean to check if the value has been set.
func (o *Event) GetCustomerAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAlias, true
}

// SetCustomerAlias sets field value
func (o *Event) SetCustomerAlias(v string) {
	o.CustomerAlias = v
}

// GetRef returns the Ref field value
func (o *Event) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *Event) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *Event) SetRef(v string) {
	o.Ref = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Event) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Event) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Event) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Event) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["customerAlias"] = o.CustomerAlias
	toSerialize["ref"] = o.Ref
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Event) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"timestamp",
		"customerAlias",
		"ref",
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

	varEvent := _Event{}

	err = json.Unmarshal(data, &varEvent)

	if err != nil {
		return err
	}

	*o = Event(varEvent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "customerAlias")
		delete(additionalProperties, "ref")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



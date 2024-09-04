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

// checks if the GetInvoiceResponseInvoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoiceResponseInvoice{}

// GetInvoiceResponseInvoice struct for GetInvoiceResponseInvoice
type GetInvoiceResponseInvoice struct {
	// The id of the customer that the invoice is associated with
	CustomerId string `json:"customerId" validate:"regexp=^[0-9a-fA-F]{24}$"`
	// The id of the contract that the invoice is associated with
	ContractId *string `json:"contractId,omitempty" validate:"regexp=^[0-9a-fA-F]{24}$"`
	// The name of the invoice, usually a description of the billing period
	Name string `json:"name"`
	BillingCycle Period `json:"billingCycle"`
	LineItems []LineItem `json:"lineItems"`
	// The total amount of the invoice
	Amount float32 `json:"amount"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
	AdditionalProperties map[string]interface{}
}

type _GetInvoiceResponseInvoice GetInvoiceResponseInvoice

// NewGetInvoiceResponseInvoice instantiates a new GetInvoiceResponseInvoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoiceResponseInvoice(customerId string, name string, billingCycle Period, lineItems []LineItem, amount float32, id string, createdAt time.Time, updatedAt time.Time, deletedAt string) *GetInvoiceResponseInvoice {
	this := GetInvoiceResponseInvoice{}
	this.CustomerId = customerId
	this.Name = name
	this.BillingCycle = billingCycle
	this.LineItems = lineItems
	this.Amount = amount
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DeletedAt = deletedAt
	return &this
}

// NewGetInvoiceResponseInvoiceWithDefaults instantiates a new GetInvoiceResponseInvoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoiceResponseInvoiceWithDefaults() *GetInvoiceResponseInvoice {
	this := GetInvoiceResponseInvoice{}
	return &this
}

// GetCustomerId returns the CustomerId field value
func (o *GetInvoiceResponseInvoice) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *GetInvoiceResponseInvoice) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *GetInvoiceResponseInvoice) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *GetInvoiceResponseInvoice) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *GetInvoiceResponseInvoice) SetContractId(v string) {
	o.ContractId = &v
}

// GetName returns the Name field value
func (o *GetInvoiceResponseInvoice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetInvoiceResponseInvoice) SetName(v string) {
	o.Name = v
}

// GetBillingCycle returns the BillingCycle field value
func (o *GetInvoiceResponseInvoice) GetBillingCycle() Period {
	if o == nil {
		var ret Period
		return ret
	}

	return o.BillingCycle
}

// GetBillingCycleOk returns a tuple with the BillingCycle field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetBillingCycleOk() (*Period, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingCycle, true
}

// SetBillingCycle sets field value
func (o *GetInvoiceResponseInvoice) SetBillingCycle(v Period) {
	o.BillingCycle = v
}

// GetLineItems returns the LineItems field value
func (o *GetInvoiceResponseInvoice) GetLineItems() []LineItem {
	if o == nil {
		var ret []LineItem
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetLineItemsOk() ([]LineItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *GetInvoiceResponseInvoice) SetLineItems(v []LineItem) {
	o.LineItems = v
}

// GetAmount returns the Amount field value
func (o *GetInvoiceResponseInvoice) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetInvoiceResponseInvoice) SetAmount(v float32) {
	o.Amount = v
}

// GetId returns the Id field value
func (o *GetInvoiceResponseInvoice) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetInvoiceResponseInvoice) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetInvoiceResponseInvoice) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetInvoiceResponseInvoice) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetInvoiceResponseInvoice) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetInvoiceResponseInvoice) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDeletedAt returns the DeletedAt field value
func (o *GetInvoiceResponseInvoice) GetDeletedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *GetInvoiceResponseInvoice) GetDeletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *GetInvoiceResponseInvoice) SetDeletedAt(v string) {
	o.DeletedAt = v
}

func (o GetInvoiceResponseInvoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInvoiceResponseInvoice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerId"] = o.CustomerId
	if !IsNil(o.ContractId) {
		toSerialize["contractId"] = o.ContractId
	}
	toSerialize["name"] = o.Name
	toSerialize["billingCycle"] = o.BillingCycle
	toSerialize["lineItems"] = o.LineItems
	toSerialize["amount"] = o.Amount
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["deletedAt"] = o.DeletedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetInvoiceResponseInvoice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerId",
		"name",
		"billingCycle",
		"lineItems",
		"amount",
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

	varGetInvoiceResponseInvoice := _GetInvoiceResponseInvoice{}

	err = json.Unmarshal(data, &varGetInvoiceResponseInvoice)

	if err != nil {
		return err
	}

	*o = GetInvoiceResponseInvoice(varGetInvoiceResponseInvoice)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "contractId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "billingCycle")
		delete(additionalProperties, "lineItems")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "deletedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetInvoiceResponseInvoice struct {
	value *GetInvoiceResponseInvoice
	isSet bool
}

func (v NullableGetInvoiceResponseInvoice) Get() *GetInvoiceResponseInvoice {
	return v.value
}

func (v *NullableGetInvoiceResponseInvoice) Set(val *GetInvoiceResponseInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoiceResponseInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoiceResponseInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoiceResponseInvoice(val *GetInvoiceResponseInvoice) *NullableGetInvoiceResponseInvoice {
	return &NullableGetInvoiceResponseInvoice{value: val, isSet: true}
}

func (v NullableGetInvoiceResponseInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvoiceResponseInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the GetCommitmentReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCommitmentReportResponse{}

// GetCommitmentReportResponse struct for GetCommitmentReportResponse
type GetCommitmentReportResponse struct {
	ContractStatus ContractStatus `json:"contractStatus"`
	CustomerName string `json:"customerName"`
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
	ContractStartDate time.Time `json:"contractStartDate"`
	DaysToContractEnd *float32 `json:"daysToContractEnd,omitempty"`
	ProductVariantName string `json:"productVariantName"`
	CommitmentConsumptionCurrencyAmount float32 `json:"commitmentConsumptionCurrencyAmount"`
	CommitmentConsumptionUnitsAmount float32 `json:"commitmentConsumptionUnitsAmount"`
	CommitmentConsumptionPercentage float32 `json:"commitmentConsumptionPercentage"`
	TotalCommitmentCurrencyAmount float32 `json:"totalCommitmentCurrencyAmount"`
	TotalCommitmentUnitsAmount float32 `json:"totalCommitmentUnitsAmount"`
	CommercialTermsAmount *float32 `json:"commercialTermsAmount,omitempty"`
	ErpId *string `json:"erpId,omitempty"`
	Currency Currency `json:"currency"`
	AdditionalProperties map[string]interface{}
}

type _GetCommitmentReportResponse GetCommitmentReportResponse

// NewGetCommitmentReportResponse instantiates a new GetCommitmentReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCommitmentReportResponse(contractStatus ContractStatus, customerName string, startDate time.Time, endDate time.Time, contractStartDate time.Time, productVariantName string, commitmentConsumptionCurrencyAmount float32, commitmentConsumptionUnitsAmount float32, commitmentConsumptionPercentage float32, totalCommitmentCurrencyAmount float32, totalCommitmentUnitsAmount float32, currency Currency) *GetCommitmentReportResponse {
	this := GetCommitmentReportResponse{}
	this.ContractStatus = contractStatus
	this.CustomerName = customerName
	this.StartDate = startDate
	this.EndDate = endDate
	this.ContractStartDate = contractStartDate
	this.ProductVariantName = productVariantName
	this.CommitmentConsumptionCurrencyAmount = commitmentConsumptionCurrencyAmount
	this.CommitmentConsumptionUnitsAmount = commitmentConsumptionUnitsAmount
	this.CommitmentConsumptionPercentage = commitmentConsumptionPercentage
	this.TotalCommitmentCurrencyAmount = totalCommitmentCurrencyAmount
	this.TotalCommitmentUnitsAmount = totalCommitmentUnitsAmount
	this.Currency = currency
	return &this
}

// NewGetCommitmentReportResponseWithDefaults instantiates a new GetCommitmentReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCommitmentReportResponseWithDefaults() *GetCommitmentReportResponse {
	this := GetCommitmentReportResponse{}
	return &this
}

// GetContractStatus returns the ContractStatus field value
func (o *GetCommitmentReportResponse) GetContractStatus() ContractStatus {
	if o == nil {
		var ret ContractStatus
		return ret
	}

	return o.ContractStatus
}

// GetContractStatusOk returns a tuple with the ContractStatus field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetContractStatusOk() (*ContractStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractStatus, true
}

// SetContractStatus sets field value
func (o *GetCommitmentReportResponse) SetContractStatus(v ContractStatus) {
	o.ContractStatus = v
}

// GetCustomerName returns the CustomerName field value
func (o *GetCommitmentReportResponse) GetCustomerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetCustomerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerName, true
}

// SetCustomerName sets field value
func (o *GetCommitmentReportResponse) SetCustomerName(v string) {
	o.CustomerName = v
}

// GetStartDate returns the StartDate field value
func (o *GetCommitmentReportResponse) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *GetCommitmentReportResponse) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *GetCommitmentReportResponse) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *GetCommitmentReportResponse) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetContractStartDate returns the ContractStartDate field value
func (o *GetCommitmentReportResponse) GetContractStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ContractStartDate
}

// GetContractStartDateOk returns a tuple with the ContractStartDate field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetContractStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractStartDate, true
}

// SetContractStartDate sets field value
func (o *GetCommitmentReportResponse) SetContractStartDate(v time.Time) {
	o.ContractStartDate = v
}

// GetDaysToContractEnd returns the DaysToContractEnd field value if set, zero value otherwise.
func (o *GetCommitmentReportResponse) GetDaysToContractEnd() float32 {
	if o == nil || IsNil(o.DaysToContractEnd) {
		var ret float32
		return ret
	}
	return *o.DaysToContractEnd
}

// GetDaysToContractEndOk returns a tuple with the DaysToContractEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetDaysToContractEndOk() (*float32, bool) {
	if o == nil || IsNil(o.DaysToContractEnd) {
		return nil, false
	}
	return o.DaysToContractEnd, true
}

// HasDaysToContractEnd returns a boolean if a field has been set.
func (o *GetCommitmentReportResponse) HasDaysToContractEnd() bool {
	if o != nil && !IsNil(o.DaysToContractEnd) {
		return true
	}

	return false
}

// SetDaysToContractEnd gets a reference to the given float32 and assigns it to the DaysToContractEnd field.
func (o *GetCommitmentReportResponse) SetDaysToContractEnd(v float32) {
	o.DaysToContractEnd = &v
}

// GetProductVariantName returns the ProductVariantName field value
func (o *GetCommitmentReportResponse) GetProductVariantName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductVariantName
}

// GetProductVariantNameOk returns a tuple with the ProductVariantName field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetProductVariantNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductVariantName, true
}

// SetProductVariantName sets field value
func (o *GetCommitmentReportResponse) SetProductVariantName(v string) {
	o.ProductVariantName = v
}

// GetCommitmentConsumptionCurrencyAmount returns the CommitmentConsumptionCurrencyAmount field value
func (o *GetCommitmentReportResponse) GetCommitmentConsumptionCurrencyAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CommitmentConsumptionCurrencyAmount
}

// GetCommitmentConsumptionCurrencyAmountOk returns a tuple with the CommitmentConsumptionCurrencyAmount field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetCommitmentConsumptionCurrencyAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentConsumptionCurrencyAmount, true
}

// SetCommitmentConsumptionCurrencyAmount sets field value
func (o *GetCommitmentReportResponse) SetCommitmentConsumptionCurrencyAmount(v float32) {
	o.CommitmentConsumptionCurrencyAmount = v
}

// GetCommitmentConsumptionUnitsAmount returns the CommitmentConsumptionUnitsAmount field value
func (o *GetCommitmentReportResponse) GetCommitmentConsumptionUnitsAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CommitmentConsumptionUnitsAmount
}

// GetCommitmentConsumptionUnitsAmountOk returns a tuple with the CommitmentConsumptionUnitsAmount field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetCommitmentConsumptionUnitsAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentConsumptionUnitsAmount, true
}

// SetCommitmentConsumptionUnitsAmount sets field value
func (o *GetCommitmentReportResponse) SetCommitmentConsumptionUnitsAmount(v float32) {
	o.CommitmentConsumptionUnitsAmount = v
}

// GetCommitmentConsumptionPercentage returns the CommitmentConsumptionPercentage field value
func (o *GetCommitmentReportResponse) GetCommitmentConsumptionPercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CommitmentConsumptionPercentage
}

// GetCommitmentConsumptionPercentageOk returns a tuple with the CommitmentConsumptionPercentage field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetCommitmentConsumptionPercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentConsumptionPercentage, true
}

// SetCommitmentConsumptionPercentage sets field value
func (o *GetCommitmentReportResponse) SetCommitmentConsumptionPercentage(v float32) {
	o.CommitmentConsumptionPercentage = v
}

// GetTotalCommitmentCurrencyAmount returns the TotalCommitmentCurrencyAmount field value
func (o *GetCommitmentReportResponse) GetTotalCommitmentCurrencyAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCommitmentCurrencyAmount
}

// GetTotalCommitmentCurrencyAmountOk returns a tuple with the TotalCommitmentCurrencyAmount field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetTotalCommitmentCurrencyAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCommitmentCurrencyAmount, true
}

// SetTotalCommitmentCurrencyAmount sets field value
func (o *GetCommitmentReportResponse) SetTotalCommitmentCurrencyAmount(v float32) {
	o.TotalCommitmentCurrencyAmount = v
}

// GetTotalCommitmentUnitsAmount returns the TotalCommitmentUnitsAmount field value
func (o *GetCommitmentReportResponse) GetTotalCommitmentUnitsAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCommitmentUnitsAmount
}

// GetTotalCommitmentUnitsAmountOk returns a tuple with the TotalCommitmentUnitsAmount field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetTotalCommitmentUnitsAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCommitmentUnitsAmount, true
}

// SetTotalCommitmentUnitsAmount sets field value
func (o *GetCommitmentReportResponse) SetTotalCommitmentUnitsAmount(v float32) {
	o.TotalCommitmentUnitsAmount = v
}

// GetCommercialTermsAmount returns the CommercialTermsAmount field value if set, zero value otherwise.
func (o *GetCommitmentReportResponse) GetCommercialTermsAmount() float32 {
	if o == nil || IsNil(o.CommercialTermsAmount) {
		var ret float32
		return ret
	}
	return *o.CommercialTermsAmount
}

// GetCommercialTermsAmountOk returns a tuple with the CommercialTermsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetCommercialTermsAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CommercialTermsAmount) {
		return nil, false
	}
	return o.CommercialTermsAmount, true
}

// HasCommercialTermsAmount returns a boolean if a field has been set.
func (o *GetCommitmentReportResponse) HasCommercialTermsAmount() bool {
	if o != nil && !IsNil(o.CommercialTermsAmount) {
		return true
	}

	return false
}

// SetCommercialTermsAmount gets a reference to the given float32 and assigns it to the CommercialTermsAmount field.
func (o *GetCommitmentReportResponse) SetCommercialTermsAmount(v float32) {
	o.CommercialTermsAmount = &v
}

// GetErpId returns the ErpId field value if set, zero value otherwise.
func (o *GetCommitmentReportResponse) GetErpId() string {
	if o == nil || IsNil(o.ErpId) {
		var ret string
		return ret
	}
	return *o.ErpId
}

// GetErpIdOk returns a tuple with the ErpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetErpIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErpId) {
		return nil, false
	}
	return o.ErpId, true
}

// HasErpId returns a boolean if a field has been set.
func (o *GetCommitmentReportResponse) HasErpId() bool {
	if o != nil && !IsNil(o.ErpId) {
		return true
	}

	return false
}

// SetErpId gets a reference to the given string and assigns it to the ErpId field.
func (o *GetCommitmentReportResponse) SetErpId(v string) {
	o.ErpId = &v
}

// GetCurrency returns the Currency field value
func (o *GetCommitmentReportResponse) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetCommitmentReportResponse) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetCommitmentReportResponse) SetCurrency(v Currency) {
	o.Currency = v
}

func (o GetCommitmentReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCommitmentReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contractStatus"] = o.ContractStatus
	toSerialize["customerName"] = o.CustomerName
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	toSerialize["contractStartDate"] = o.ContractStartDate
	if !IsNil(o.DaysToContractEnd) {
		toSerialize["daysToContractEnd"] = o.DaysToContractEnd
	}
	toSerialize["productVariantName"] = o.ProductVariantName
	toSerialize["commitmentConsumptionCurrencyAmount"] = o.CommitmentConsumptionCurrencyAmount
	toSerialize["commitmentConsumptionUnitsAmount"] = o.CommitmentConsumptionUnitsAmount
	toSerialize["commitmentConsumptionPercentage"] = o.CommitmentConsumptionPercentage
	toSerialize["totalCommitmentCurrencyAmount"] = o.TotalCommitmentCurrencyAmount
	toSerialize["totalCommitmentUnitsAmount"] = o.TotalCommitmentUnitsAmount
	if !IsNil(o.CommercialTermsAmount) {
		toSerialize["commercialTermsAmount"] = o.CommercialTermsAmount
	}
	if !IsNil(o.ErpId) {
		toSerialize["erpId"] = o.ErpId
	}
	toSerialize["currency"] = o.Currency

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetCommitmentReportResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contractStatus",
		"customerName",
		"startDate",
		"endDate",
		"contractStartDate",
		"productVariantName",
		"commitmentConsumptionCurrencyAmount",
		"commitmentConsumptionUnitsAmount",
		"commitmentConsumptionPercentage",
		"totalCommitmentCurrencyAmount",
		"totalCommitmentUnitsAmount",
		"currency",
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

	varGetCommitmentReportResponse := _GetCommitmentReportResponse{}

	err = json.Unmarshal(data, &varGetCommitmentReportResponse)

	if err != nil {
		return err
	}

	*o = GetCommitmentReportResponse(varGetCommitmentReportResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contractStatus")
		delete(additionalProperties, "customerName")
		delete(additionalProperties, "startDate")
		delete(additionalProperties, "endDate")
		delete(additionalProperties, "contractStartDate")
		delete(additionalProperties, "daysToContractEnd")
		delete(additionalProperties, "productVariantName")
		delete(additionalProperties, "commitmentConsumptionCurrencyAmount")
		delete(additionalProperties, "commitmentConsumptionUnitsAmount")
		delete(additionalProperties, "commitmentConsumptionPercentage")
		delete(additionalProperties, "totalCommitmentCurrencyAmount")
		delete(additionalProperties, "totalCommitmentUnitsAmount")
		delete(additionalProperties, "commercialTermsAmount")
		delete(additionalProperties, "erpId")
		delete(additionalProperties, "currency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetCommitmentReportResponse struct {
	value *GetCommitmentReportResponse
	isSet bool
}

func (v NullableGetCommitmentReportResponse) Get() *GetCommitmentReportResponse {
	return v.value
}

func (v *NullableGetCommitmentReportResponse) Set(val *GetCommitmentReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCommitmentReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCommitmentReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCommitmentReportResponse(val *GetCommitmentReportResponse) *NullableGetCommitmentReportResponse {
	return &NullableGetCommitmentReportResponse{value: val, isSet: true}
}

func (v NullableGetCommitmentReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCommitmentReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



# GetProductsUsageReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractStatus** | [**ContractStatus**](ContractStatus.md) |  | 
**CustomerName** | **string** |  | 
**StartDate** | **time.Time** |  | 
**EndDate** | **time.Time** |  | 
**ContractStartDate** | **time.Time** |  | 
**DaysToContractEnd** | Pointer to **float32** |  | [optional] 
**ProductVariantName** | **string** |  | 
**CommitmentConsumptionPercentage** | Pointer to **float32** |  | [optional] 
**TotalCommitmentCurrencyAmount** | Pointer to **float32** |  | [optional] 
**TotalCommitmentUnitsAmount** | Pointer to **float32** |  | [optional] 
**CommercialTermsAmount** | Pointer to **float32** |  | [optional] 
**ErpId** | Pointer to **string** |  | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 
**UsageConsumptionCurrencyAmount** | Pointer to **float32** |  | [optional] 
**UsageConsumptionUnitsAmount** | Pointer to **float32** |  | [optional] 

## Methods

### NewGetProductsUsageReportResponse

`func NewGetProductsUsageReportResponse(contractStatus ContractStatus, customerName string, startDate time.Time, endDate time.Time, contractStartDate time.Time, productVariantName string, currency Currency, ) *GetProductsUsageReportResponse`

NewGetProductsUsageReportResponse instantiates a new GetProductsUsageReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProductsUsageReportResponseWithDefaults

`func NewGetProductsUsageReportResponseWithDefaults() *GetProductsUsageReportResponse`

NewGetProductsUsageReportResponseWithDefaults instantiates a new GetProductsUsageReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractStatus

`func (o *GetProductsUsageReportResponse) GetContractStatus() ContractStatus`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *GetProductsUsageReportResponse) GetContractStatusOk() (*ContractStatus, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *GetProductsUsageReportResponse) SetContractStatus(v ContractStatus)`

SetContractStatus sets ContractStatus field to given value.


### GetCustomerName

`func (o *GetProductsUsageReportResponse) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *GetProductsUsageReportResponse) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *GetProductsUsageReportResponse) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetStartDate

`func (o *GetProductsUsageReportResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetProductsUsageReportResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetProductsUsageReportResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *GetProductsUsageReportResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetProductsUsageReportResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetProductsUsageReportResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetContractStartDate

`func (o *GetProductsUsageReportResponse) GetContractStartDate() time.Time`

GetContractStartDate returns the ContractStartDate field if non-nil, zero value otherwise.

### GetContractStartDateOk

`func (o *GetProductsUsageReportResponse) GetContractStartDateOk() (*time.Time, bool)`

GetContractStartDateOk returns a tuple with the ContractStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStartDate

`func (o *GetProductsUsageReportResponse) SetContractStartDate(v time.Time)`

SetContractStartDate sets ContractStartDate field to given value.


### GetDaysToContractEnd

`func (o *GetProductsUsageReportResponse) GetDaysToContractEnd() float32`

GetDaysToContractEnd returns the DaysToContractEnd field if non-nil, zero value otherwise.

### GetDaysToContractEndOk

`func (o *GetProductsUsageReportResponse) GetDaysToContractEndOk() (*float32, bool)`

GetDaysToContractEndOk returns a tuple with the DaysToContractEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToContractEnd

`func (o *GetProductsUsageReportResponse) SetDaysToContractEnd(v float32)`

SetDaysToContractEnd sets DaysToContractEnd field to given value.

### HasDaysToContractEnd

`func (o *GetProductsUsageReportResponse) HasDaysToContractEnd() bool`

HasDaysToContractEnd returns a boolean if a field has been set.

### GetProductVariantName

`func (o *GetProductsUsageReportResponse) GetProductVariantName() string`

GetProductVariantName returns the ProductVariantName field if non-nil, zero value otherwise.

### GetProductVariantNameOk

`func (o *GetProductsUsageReportResponse) GetProductVariantNameOk() (*string, bool)`

GetProductVariantNameOk returns a tuple with the ProductVariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariantName

`func (o *GetProductsUsageReportResponse) SetProductVariantName(v string)`

SetProductVariantName sets ProductVariantName field to given value.


### GetCommitmentConsumptionPercentage

`func (o *GetProductsUsageReportResponse) GetCommitmentConsumptionPercentage() float32`

GetCommitmentConsumptionPercentage returns the CommitmentConsumptionPercentage field if non-nil, zero value otherwise.

### GetCommitmentConsumptionPercentageOk

`func (o *GetProductsUsageReportResponse) GetCommitmentConsumptionPercentageOk() (*float32, bool)`

GetCommitmentConsumptionPercentageOk returns a tuple with the CommitmentConsumptionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentConsumptionPercentage

`func (o *GetProductsUsageReportResponse) SetCommitmentConsumptionPercentage(v float32)`

SetCommitmentConsumptionPercentage sets CommitmentConsumptionPercentage field to given value.

### HasCommitmentConsumptionPercentage

`func (o *GetProductsUsageReportResponse) HasCommitmentConsumptionPercentage() bool`

HasCommitmentConsumptionPercentage returns a boolean if a field has been set.

### GetTotalCommitmentCurrencyAmount

`func (o *GetProductsUsageReportResponse) GetTotalCommitmentCurrencyAmount() float32`

GetTotalCommitmentCurrencyAmount returns the TotalCommitmentCurrencyAmount field if non-nil, zero value otherwise.

### GetTotalCommitmentCurrencyAmountOk

`func (o *GetProductsUsageReportResponse) GetTotalCommitmentCurrencyAmountOk() (*float32, bool)`

GetTotalCommitmentCurrencyAmountOk returns a tuple with the TotalCommitmentCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommitmentCurrencyAmount

`func (o *GetProductsUsageReportResponse) SetTotalCommitmentCurrencyAmount(v float32)`

SetTotalCommitmentCurrencyAmount sets TotalCommitmentCurrencyAmount field to given value.

### HasTotalCommitmentCurrencyAmount

`func (o *GetProductsUsageReportResponse) HasTotalCommitmentCurrencyAmount() bool`

HasTotalCommitmentCurrencyAmount returns a boolean if a field has been set.

### GetTotalCommitmentUnitsAmount

`func (o *GetProductsUsageReportResponse) GetTotalCommitmentUnitsAmount() float32`

GetTotalCommitmentUnitsAmount returns the TotalCommitmentUnitsAmount field if non-nil, zero value otherwise.

### GetTotalCommitmentUnitsAmountOk

`func (o *GetProductsUsageReportResponse) GetTotalCommitmentUnitsAmountOk() (*float32, bool)`

GetTotalCommitmentUnitsAmountOk returns a tuple with the TotalCommitmentUnitsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommitmentUnitsAmount

`func (o *GetProductsUsageReportResponse) SetTotalCommitmentUnitsAmount(v float32)`

SetTotalCommitmentUnitsAmount sets TotalCommitmentUnitsAmount field to given value.

### HasTotalCommitmentUnitsAmount

`func (o *GetProductsUsageReportResponse) HasTotalCommitmentUnitsAmount() bool`

HasTotalCommitmentUnitsAmount returns a boolean if a field has been set.

### GetCommercialTermsAmount

`func (o *GetProductsUsageReportResponse) GetCommercialTermsAmount() float32`

GetCommercialTermsAmount returns the CommercialTermsAmount field if non-nil, zero value otherwise.

### GetCommercialTermsAmountOk

`func (o *GetProductsUsageReportResponse) GetCommercialTermsAmountOk() (*float32, bool)`

GetCommercialTermsAmountOk returns a tuple with the CommercialTermsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialTermsAmount

`func (o *GetProductsUsageReportResponse) SetCommercialTermsAmount(v float32)`

SetCommercialTermsAmount sets CommercialTermsAmount field to given value.

### HasCommercialTermsAmount

`func (o *GetProductsUsageReportResponse) HasCommercialTermsAmount() bool`

HasCommercialTermsAmount returns a boolean if a field has been set.

### GetErpId

`func (o *GetProductsUsageReportResponse) GetErpId() string`

GetErpId returns the ErpId field if non-nil, zero value otherwise.

### GetErpIdOk

`func (o *GetProductsUsageReportResponse) GetErpIdOk() (*string, bool)`

GetErpIdOk returns a tuple with the ErpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErpId

`func (o *GetProductsUsageReportResponse) SetErpId(v string)`

SetErpId sets ErpId field to given value.

### HasErpId

`func (o *GetProductsUsageReportResponse) HasErpId() bool`

HasErpId returns a boolean if a field has been set.

### GetCurrency

`func (o *GetProductsUsageReportResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetProductsUsageReportResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetProductsUsageReportResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### GetUsageConsumptionCurrencyAmount

`func (o *GetProductsUsageReportResponse) GetUsageConsumptionCurrencyAmount() float32`

GetUsageConsumptionCurrencyAmount returns the UsageConsumptionCurrencyAmount field if non-nil, zero value otherwise.

### GetUsageConsumptionCurrencyAmountOk

`func (o *GetProductsUsageReportResponse) GetUsageConsumptionCurrencyAmountOk() (*float32, bool)`

GetUsageConsumptionCurrencyAmountOk returns a tuple with the UsageConsumptionCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageConsumptionCurrencyAmount

`func (o *GetProductsUsageReportResponse) SetUsageConsumptionCurrencyAmount(v float32)`

SetUsageConsumptionCurrencyAmount sets UsageConsumptionCurrencyAmount field to given value.

### HasUsageConsumptionCurrencyAmount

`func (o *GetProductsUsageReportResponse) HasUsageConsumptionCurrencyAmount() bool`

HasUsageConsumptionCurrencyAmount returns a boolean if a field has been set.

### GetUsageConsumptionUnitsAmount

`func (o *GetProductsUsageReportResponse) GetUsageConsumptionUnitsAmount() float32`

GetUsageConsumptionUnitsAmount returns the UsageConsumptionUnitsAmount field if non-nil, zero value otherwise.

### GetUsageConsumptionUnitsAmountOk

`func (o *GetProductsUsageReportResponse) GetUsageConsumptionUnitsAmountOk() (*float32, bool)`

GetUsageConsumptionUnitsAmountOk returns a tuple with the UsageConsumptionUnitsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageConsumptionUnitsAmount

`func (o *GetProductsUsageReportResponse) SetUsageConsumptionUnitsAmount(v float32)`

SetUsageConsumptionUnitsAmount sets UsageConsumptionUnitsAmount field to given value.

### HasUsageConsumptionUnitsAmount

`func (o *GetProductsUsageReportResponse) HasUsageConsumptionUnitsAmount() bool`

HasUsageConsumptionUnitsAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



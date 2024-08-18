# GetCommitmentReportResponse

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
**CommitmentConsumptionCurrencyAmount** | **float32** |  | 
**CommitmentConsumptionUnitsAmount** | **float32** |  | 
**CommitmentConsumptionPercentage** | **float32** |  | 
**TotalCommitmentCurrencyAmount** | **float32** |  | 
**TotalCommitmentUnitsAmount** | **float32** |  | 
**CommercialTermsAmount** | Pointer to **float32** |  | [optional] 
**ErpId** | Pointer to **string** |  | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 

## Methods

### NewGetCommitmentReportResponse

`func NewGetCommitmentReportResponse(contractStatus ContractStatus, customerName string, startDate time.Time, endDate time.Time, contractStartDate time.Time, productVariantName string, commitmentConsumptionCurrencyAmount float32, commitmentConsumptionUnitsAmount float32, commitmentConsumptionPercentage float32, totalCommitmentCurrencyAmount float32, totalCommitmentUnitsAmount float32, currency Currency, ) *GetCommitmentReportResponse`

NewGetCommitmentReportResponse instantiates a new GetCommitmentReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommitmentReportResponseWithDefaults

`func NewGetCommitmentReportResponseWithDefaults() *GetCommitmentReportResponse`

NewGetCommitmentReportResponseWithDefaults instantiates a new GetCommitmentReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractStatus

`func (o *GetCommitmentReportResponse) GetContractStatus() ContractStatus`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *GetCommitmentReportResponse) GetContractStatusOk() (*ContractStatus, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *GetCommitmentReportResponse) SetContractStatus(v ContractStatus)`

SetContractStatus sets ContractStatus field to given value.


### GetCustomerName

`func (o *GetCommitmentReportResponse) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *GetCommitmentReportResponse) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *GetCommitmentReportResponse) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetStartDate

`func (o *GetCommitmentReportResponse) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetCommitmentReportResponse) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetCommitmentReportResponse) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *GetCommitmentReportResponse) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetCommitmentReportResponse) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetCommitmentReportResponse) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetContractStartDate

`func (o *GetCommitmentReportResponse) GetContractStartDate() time.Time`

GetContractStartDate returns the ContractStartDate field if non-nil, zero value otherwise.

### GetContractStartDateOk

`func (o *GetCommitmentReportResponse) GetContractStartDateOk() (*time.Time, bool)`

GetContractStartDateOk returns a tuple with the ContractStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStartDate

`func (o *GetCommitmentReportResponse) SetContractStartDate(v time.Time)`

SetContractStartDate sets ContractStartDate field to given value.


### GetDaysToContractEnd

`func (o *GetCommitmentReportResponse) GetDaysToContractEnd() float32`

GetDaysToContractEnd returns the DaysToContractEnd field if non-nil, zero value otherwise.

### GetDaysToContractEndOk

`func (o *GetCommitmentReportResponse) GetDaysToContractEndOk() (*float32, bool)`

GetDaysToContractEndOk returns a tuple with the DaysToContractEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToContractEnd

`func (o *GetCommitmentReportResponse) SetDaysToContractEnd(v float32)`

SetDaysToContractEnd sets DaysToContractEnd field to given value.

### HasDaysToContractEnd

`func (o *GetCommitmentReportResponse) HasDaysToContractEnd() bool`

HasDaysToContractEnd returns a boolean if a field has been set.

### GetProductVariantName

`func (o *GetCommitmentReportResponse) GetProductVariantName() string`

GetProductVariantName returns the ProductVariantName field if non-nil, zero value otherwise.

### GetProductVariantNameOk

`func (o *GetCommitmentReportResponse) GetProductVariantNameOk() (*string, bool)`

GetProductVariantNameOk returns a tuple with the ProductVariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariantName

`func (o *GetCommitmentReportResponse) SetProductVariantName(v string)`

SetProductVariantName sets ProductVariantName field to given value.


### GetCommitmentConsumptionCurrencyAmount

`func (o *GetCommitmentReportResponse) GetCommitmentConsumptionCurrencyAmount() float32`

GetCommitmentConsumptionCurrencyAmount returns the CommitmentConsumptionCurrencyAmount field if non-nil, zero value otherwise.

### GetCommitmentConsumptionCurrencyAmountOk

`func (o *GetCommitmentReportResponse) GetCommitmentConsumptionCurrencyAmountOk() (*float32, bool)`

GetCommitmentConsumptionCurrencyAmountOk returns a tuple with the CommitmentConsumptionCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentConsumptionCurrencyAmount

`func (o *GetCommitmentReportResponse) SetCommitmentConsumptionCurrencyAmount(v float32)`

SetCommitmentConsumptionCurrencyAmount sets CommitmentConsumptionCurrencyAmount field to given value.


### GetCommitmentConsumptionUnitsAmount

`func (o *GetCommitmentReportResponse) GetCommitmentConsumptionUnitsAmount() float32`

GetCommitmentConsumptionUnitsAmount returns the CommitmentConsumptionUnitsAmount field if non-nil, zero value otherwise.

### GetCommitmentConsumptionUnitsAmountOk

`func (o *GetCommitmentReportResponse) GetCommitmentConsumptionUnitsAmountOk() (*float32, bool)`

GetCommitmentConsumptionUnitsAmountOk returns a tuple with the CommitmentConsumptionUnitsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentConsumptionUnitsAmount

`func (o *GetCommitmentReportResponse) SetCommitmentConsumptionUnitsAmount(v float32)`

SetCommitmentConsumptionUnitsAmount sets CommitmentConsumptionUnitsAmount field to given value.


### GetCommitmentConsumptionPercentage

`func (o *GetCommitmentReportResponse) GetCommitmentConsumptionPercentage() float32`

GetCommitmentConsumptionPercentage returns the CommitmentConsumptionPercentage field if non-nil, zero value otherwise.

### GetCommitmentConsumptionPercentageOk

`func (o *GetCommitmentReportResponse) GetCommitmentConsumptionPercentageOk() (*float32, bool)`

GetCommitmentConsumptionPercentageOk returns a tuple with the CommitmentConsumptionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitmentConsumptionPercentage

`func (o *GetCommitmentReportResponse) SetCommitmentConsumptionPercentage(v float32)`

SetCommitmentConsumptionPercentage sets CommitmentConsumptionPercentage field to given value.


### GetTotalCommitmentCurrencyAmount

`func (o *GetCommitmentReportResponse) GetTotalCommitmentCurrencyAmount() float32`

GetTotalCommitmentCurrencyAmount returns the TotalCommitmentCurrencyAmount field if non-nil, zero value otherwise.

### GetTotalCommitmentCurrencyAmountOk

`func (o *GetCommitmentReportResponse) GetTotalCommitmentCurrencyAmountOk() (*float32, bool)`

GetTotalCommitmentCurrencyAmountOk returns a tuple with the TotalCommitmentCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommitmentCurrencyAmount

`func (o *GetCommitmentReportResponse) SetTotalCommitmentCurrencyAmount(v float32)`

SetTotalCommitmentCurrencyAmount sets TotalCommitmentCurrencyAmount field to given value.


### GetTotalCommitmentUnitsAmount

`func (o *GetCommitmentReportResponse) GetTotalCommitmentUnitsAmount() float32`

GetTotalCommitmentUnitsAmount returns the TotalCommitmentUnitsAmount field if non-nil, zero value otherwise.

### GetTotalCommitmentUnitsAmountOk

`func (o *GetCommitmentReportResponse) GetTotalCommitmentUnitsAmountOk() (*float32, bool)`

GetTotalCommitmentUnitsAmountOk returns a tuple with the TotalCommitmentUnitsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommitmentUnitsAmount

`func (o *GetCommitmentReportResponse) SetTotalCommitmentUnitsAmount(v float32)`

SetTotalCommitmentUnitsAmount sets TotalCommitmentUnitsAmount field to given value.


### GetCommercialTermsAmount

`func (o *GetCommitmentReportResponse) GetCommercialTermsAmount() float32`

GetCommercialTermsAmount returns the CommercialTermsAmount field if non-nil, zero value otherwise.

### GetCommercialTermsAmountOk

`func (o *GetCommitmentReportResponse) GetCommercialTermsAmountOk() (*float32, bool)`

GetCommercialTermsAmountOk returns a tuple with the CommercialTermsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialTermsAmount

`func (o *GetCommitmentReportResponse) SetCommercialTermsAmount(v float32)`

SetCommercialTermsAmount sets CommercialTermsAmount field to given value.

### HasCommercialTermsAmount

`func (o *GetCommitmentReportResponse) HasCommercialTermsAmount() bool`

HasCommercialTermsAmount returns a boolean if a field has been set.

### GetErpId

`func (o *GetCommitmentReportResponse) GetErpId() string`

GetErpId returns the ErpId field if non-nil, zero value otherwise.

### GetErpIdOk

`func (o *GetCommitmentReportResponse) GetErpIdOk() (*string, bool)`

GetErpIdOk returns a tuple with the ErpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErpId

`func (o *GetCommitmentReportResponse) SetErpId(v string)`

SetErpId sets ErpId field to given value.

### HasErpId

`func (o *GetCommitmentReportResponse) HasErpId() bool`

HasErpId returns a boolean if a field has been set.

### GetCurrency

`func (o *GetCommitmentReportResponse) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetCommitmentReportResponse) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetCommitmentReportResponse) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



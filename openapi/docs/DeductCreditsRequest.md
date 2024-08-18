# DeductCreditsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAmount** | **float32** | The amount of credits to be deducted from the user. | 
**CustomerId** | **string** | The ID of the customer to whom the credits will be deducted from. | 

## Methods

### NewDeductCreditsRequest

`func NewDeductCreditsRequest(creditAmount float32, customerId string, ) *DeductCreditsRequest`

NewDeductCreditsRequest instantiates a new DeductCreditsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeductCreditsRequestWithDefaults

`func NewDeductCreditsRequestWithDefaults() *DeductCreditsRequest`

NewDeductCreditsRequestWithDefaults instantiates a new DeductCreditsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAmount

`func (o *DeductCreditsRequest) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *DeductCreditsRequest) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *DeductCreditsRequest) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.


### GetCustomerId

`func (o *DeductCreditsRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DeductCreditsRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DeductCreditsRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



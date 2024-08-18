# GetCustomerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | [**CreateCustomerResponseCustomer**](CreateCustomerResponseCustomer.md) |  | 

## Methods

### NewGetCustomerResponse

`func NewGetCustomerResponse(customer CreateCustomerResponseCustomer, ) *GetCustomerResponse`

NewGetCustomerResponse instantiates a new GetCustomerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerResponseWithDefaults

`func NewGetCustomerResponseWithDefaults() *GetCustomerResponse`

NewGetCustomerResponseWithDefaults instantiates a new GetCustomerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *GetCustomerResponse) GetCustomer() CreateCustomerResponseCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *GetCustomerResponse) GetCustomerOk() (*CreateCustomerResponseCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *GetCustomerResponse) SetCustomer(v CreateCustomerResponseCustomer)`

SetCustomer sets Customer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



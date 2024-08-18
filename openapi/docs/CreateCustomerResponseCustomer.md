# CreateCustomerResponseCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the customer | 
**Alias** | Pointer to **string** | The alias of the customer used to match events to the customer. | [optional] 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **string** |  | 

## Methods

### NewCreateCustomerResponseCustomer

`func NewCreateCustomerResponseCustomer(name string, id string, createdAt time.Time, updatedAt time.Time, deletedAt string, ) *CreateCustomerResponseCustomer`

NewCreateCustomerResponseCustomer instantiates a new CreateCustomerResponseCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerResponseCustomerWithDefaults

`func NewCreateCustomerResponseCustomerWithDefaults() *CreateCustomerResponseCustomer`

NewCreateCustomerResponseCustomerWithDefaults instantiates a new CreateCustomerResponseCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustomerResponseCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomerResponseCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomerResponseCustomer) SetName(v string)`

SetName sets Name field to given value.


### GetAlias

`func (o *CreateCustomerResponseCustomer) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateCustomerResponseCustomer) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateCustomerResponseCustomer) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CreateCustomerResponseCustomer) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetId

`func (o *CreateCustomerResponseCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCustomerResponseCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCustomerResponseCustomer) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreateCustomerResponseCustomer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateCustomerResponseCustomer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateCustomerResponseCustomer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateCustomerResponseCustomer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateCustomerResponseCustomer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateCustomerResponseCustomer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *CreateCustomerResponseCustomer) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *CreateCustomerResponseCustomer) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *CreateCustomerResponseCustomer) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListCustomersResponseCustomersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the customer | 
**Alias** | Pointer to **string** | The alias of the customer used to match events to the customer. | [optional] 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewListCustomersResponseCustomersInner

`func NewListCustomersResponseCustomersInner(name string, id string, createdAt time.Time, updatedAt time.Time, ) *ListCustomersResponseCustomersInner`

NewListCustomersResponseCustomersInner instantiates a new ListCustomersResponseCustomersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomersResponseCustomersInnerWithDefaults

`func NewListCustomersResponseCustomersInnerWithDefaults() *ListCustomersResponseCustomersInner`

NewListCustomersResponseCustomersInnerWithDefaults instantiates a new ListCustomersResponseCustomersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListCustomersResponseCustomersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCustomersResponseCustomersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCustomersResponseCustomersInner) SetName(v string)`

SetName sets Name field to given value.


### GetAlias

`func (o *ListCustomersResponseCustomersInner) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ListCustomersResponseCustomersInner) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ListCustomersResponseCustomersInner) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ListCustomersResponseCustomersInner) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetId

`func (o *ListCustomersResponseCustomersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCustomersResponseCustomersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCustomersResponseCustomersInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListCustomersResponseCustomersInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListCustomersResponseCustomersInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListCustomersResponseCustomersInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListCustomersResponseCustomersInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListCustomersResponseCustomersInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListCustomersResponseCustomersInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



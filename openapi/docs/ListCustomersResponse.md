# ListCustomersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | [**[]ListCustomersResponseCustomersInner**](ListCustomersResponseCustomersInner.md) |  | 
**Total** | **float32** |  | 
**HasMore** | **bool** |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListCustomersResponse

`func NewListCustomersResponse(customers []ListCustomersResponseCustomersInner, total float32, hasMore bool, ) *ListCustomersResponse`

NewListCustomersResponse instantiates a new ListCustomersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomersResponseWithDefaults

`func NewListCustomersResponseWithDefaults() *ListCustomersResponse`

NewListCustomersResponseWithDefaults instantiates a new ListCustomersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomers

`func (o *ListCustomersResponse) GetCustomers() []ListCustomersResponseCustomersInner`

GetCustomers returns the Customers field if non-nil, zero value otherwise.

### GetCustomersOk

`func (o *ListCustomersResponse) GetCustomersOk() (*[]ListCustomersResponseCustomersInner, bool)`

GetCustomersOk returns a tuple with the Customers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomers

`func (o *ListCustomersResponse) SetCustomers(v []ListCustomersResponseCustomersInner)`

SetCustomers sets Customers field to given value.


### GetTotal

`func (o *ListCustomersResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListCustomersResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListCustomersResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetHasMore

`func (o *ListCustomersResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListCustomersResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListCustomersResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *ListCustomersResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListCustomersResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListCustomersResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListCustomersResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



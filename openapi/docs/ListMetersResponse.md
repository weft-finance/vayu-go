# ListMetersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meters** | [**[]ListMetersResponseMetersInner**](ListMetersResponseMetersInner.md) |  | 
**Total** | **float32** |  | 
**HasMore** | **bool** |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListMetersResponse

`func NewListMetersResponse(meters []ListMetersResponseMetersInner, total float32, hasMore bool, ) *ListMetersResponse`

NewListMetersResponse instantiates a new ListMetersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMetersResponseWithDefaults

`func NewListMetersResponseWithDefaults() *ListMetersResponse`

NewListMetersResponseWithDefaults instantiates a new ListMetersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeters

`func (o *ListMetersResponse) GetMeters() []ListMetersResponseMetersInner`

GetMeters returns the Meters field if non-nil, zero value otherwise.

### GetMetersOk

`func (o *ListMetersResponse) GetMetersOk() (*[]ListMetersResponseMetersInner, bool)`

GetMetersOk returns a tuple with the Meters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeters

`func (o *ListMetersResponse) SetMeters(v []ListMetersResponseMetersInner)`

SetMeters sets Meters field to given value.


### GetTotal

`func (o *ListMetersResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListMetersResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListMetersResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetHasMore

`func (o *ListMetersResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListMetersResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListMetersResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *ListMetersResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListMetersResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListMetersResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListMetersResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



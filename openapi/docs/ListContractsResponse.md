# ListContractsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contracts** | [**[]ListContractsResponseContractsInner**](ListContractsResponseContractsInner.md) |  | 
**Total** | **float32** |  | 
**HasMore** | **bool** |  | 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewListContractsResponse

`func NewListContractsResponse(contracts []ListContractsResponseContractsInner, total float32, hasMore bool, ) *ListContractsResponse`

NewListContractsResponse instantiates a new ListContractsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContractsResponseWithDefaults

`func NewListContractsResponseWithDefaults() *ListContractsResponse`

NewListContractsResponseWithDefaults instantiates a new ListContractsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContracts

`func (o *ListContractsResponse) GetContracts() []ListContractsResponseContractsInner`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *ListContractsResponse) GetContractsOk() (*[]ListContractsResponseContractsInner, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *ListContractsResponse) SetContracts(v []ListContractsResponseContractsInner)`

SetContracts sets Contracts field to given value.


### GetTotal

`func (o *ListContractsResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListContractsResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListContractsResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetHasMore

`func (o *ListContractsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListContractsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListContractsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetNextCursor

`func (o *ListContractsResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListContractsResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListContractsResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListContractsResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



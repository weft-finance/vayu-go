# NetSuiteSyncInvoicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationType** | **string** |  | 
**Uid** | **string** |  | 
**Data** | [**NetSuiteSyncInvoicesRequestData**](NetSuiteSyncInvoicesRequestData.md) |  | 

## Methods

### NewNetSuiteSyncInvoicesRequest

`func NewNetSuiteSyncInvoicesRequest(integrationType string, uid string, data NetSuiteSyncInvoicesRequestData, ) *NetSuiteSyncInvoicesRequest`

NewNetSuiteSyncInvoicesRequest instantiates a new NetSuiteSyncInvoicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetSuiteSyncInvoicesRequestWithDefaults

`func NewNetSuiteSyncInvoicesRequestWithDefaults() *NetSuiteSyncInvoicesRequest`

NewNetSuiteSyncInvoicesRequestWithDefaults instantiates a new NetSuiteSyncInvoicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationType

`func (o *NetSuiteSyncInvoicesRequest) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *NetSuiteSyncInvoicesRequest) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *NetSuiteSyncInvoicesRequest) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.


### GetUid

`func (o *NetSuiteSyncInvoicesRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NetSuiteSyncInvoicesRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NetSuiteSyncInvoicesRequest) SetUid(v string)`

SetUid sets Uid field to given value.


### GetData

`func (o *NetSuiteSyncInvoicesRequest) GetData() NetSuiteSyncInvoicesRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetSuiteSyncInvoicesRequest) GetDataOk() (*NetSuiteSyncInvoicesRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetSuiteSyncInvoicesRequest) SetData(v NetSuiteSyncInvoicesRequestData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



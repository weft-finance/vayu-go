# GetPlanResponsePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the plan | 
**Status** | [**PlanStatus**](PlanStatus.md) |  | 
**BillingData** | [**PlanBillingData**](PlanBillingData.md) |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **string** |  | 

## Methods

### NewGetPlanResponsePlan

`func NewGetPlanResponsePlan(name string, status PlanStatus, billingData PlanBillingData, id string, createdAt time.Time, updatedAt time.Time, deletedAt string, ) *GetPlanResponsePlan`

NewGetPlanResponsePlan instantiates a new GetPlanResponsePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlanResponsePlanWithDefaults

`func NewGetPlanResponsePlanWithDefaults() *GetPlanResponsePlan`

NewGetPlanResponsePlanWithDefaults instantiates a new GetPlanResponsePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetPlanResponsePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPlanResponsePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPlanResponsePlan) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *GetPlanResponsePlan) GetStatus() PlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPlanResponsePlan) GetStatusOk() (*PlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPlanResponsePlan) SetStatus(v PlanStatus)`

SetStatus sets Status field to given value.


### GetBillingData

`func (o *GetPlanResponsePlan) GetBillingData() PlanBillingData`

GetBillingData returns the BillingData field if non-nil, zero value otherwise.

### GetBillingDataOk

`func (o *GetPlanResponsePlan) GetBillingDataOk() (*PlanBillingData, bool)`

GetBillingDataOk returns a tuple with the BillingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingData

`func (o *GetPlanResponsePlan) SetBillingData(v PlanBillingData)`

SetBillingData sets BillingData field to given value.


### GetId

`func (o *GetPlanResponsePlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPlanResponsePlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPlanResponsePlan) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetPlanResponsePlan) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetPlanResponsePlan) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetPlanResponsePlan) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetPlanResponsePlan) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetPlanResponsePlan) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetPlanResponsePlan) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *GetPlanResponsePlan) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GetPlanResponsePlan) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GetPlanResponsePlan) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



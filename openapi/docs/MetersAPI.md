# \MetersAPI

All URIs are relative to *https://connect.withvayu.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMeter**](MetersAPI.md#DeleteMeter) | **Delete** /meters/{meterId} | Delete Meter
[**GetMeter**](MetersAPI.md#GetMeter) | **Get** /meters/{meterId} | Get Meter
[**ListMeters**](MetersAPI.md#ListMeters) | **Get** /meters | List Meters
[**MetersMeterIdOptions**](MetersAPI.md#MetersMeterIdOptions) | **Options** /meters/{meterId} | 
[**MetersOptions**](MetersAPI.md#MetersOptions) | **Options** /meters | 
[**UpdateMeter**](MetersAPI.md#UpdateMeter) | **Put** /meters/{meterId} | Update Meter



## DeleteMeter

> DeleteMeterResponse DeleteMeter(ctx, meterId).Execute()

Delete Meter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	meterId := "meterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.DeleteMeter(context.Background(), meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.DeleteMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMeter`: DeleteMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.DeleteMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMeterResponse**](DeleteMeterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeter

> GetMeterResponse GetMeter(ctx, meterId).Execute()

Get Meter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	meterId := "meterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.GetMeter(context.Background(), meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.GetMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMeter`: GetMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.GetMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMeterResponse**](GetMeterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMeters

> ListMetersResponse ListMeters(ctx).Limit(limit).Cursor(cursor).Execute()

List Meters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := float32(8.14) // float32 |  (optional) (default to 10)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.ListMeters(context.Background()).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.ListMeters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMeters`: ListMetersResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.ListMeters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** |  | [default to 10]
 **cursor** | **string** |  | 

### Return type

[**ListMetersResponse**](ListMetersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersMeterIdOptions

> MetersMeterIdOptions(ctx, meterId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	meterId := "meterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetersAPI.MetersMeterIdOptions(context.Background(), meterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersMeterIdOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMetersMeterIdOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetersOptions

> MetersOptions(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MetersAPI.MetersOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.MetersOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMetersOptionsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeter

> UpdateMeterResponse UpdateMeter(ctx, meterId).UpdateMeterRequest(updateMeterRequest).Execute()

Update Meter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateMeterRequest := *openapiclient.NewUpdateMeterRequest() // UpdateMeterRequest | 
	meterId := "meterId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetersAPI.UpdateMeter(context.Background(), meterId).UpdateMeterRequest(updateMeterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetersAPI.UpdateMeter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMeter`: UpdateMeterResponse
	fmt.Fprintf(os.Stdout, "Response from `MetersAPI.UpdateMeter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMeterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMeterRequest** | [**UpdateMeterRequest**](UpdateMeterRequest.md) |  | 


### Return type

[**UpdateMeterResponse**](UpdateMeterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


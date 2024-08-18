# \EventsAPI

All URIs are relative to *https://connect.withvayu.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEventByRefId**](EventsAPI.md#DeleteEventByRefId) | **Delete** /events/{refId} | Delete an event by refId
[**EventsDryRunOptions**](EventsAPI.md#EventsDryRunOptions) | **Options** /events/dry-run | 
[**EventsOptions**](EventsAPI.md#EventsOptions) | **Options** /events | 
[**EventsRefIdOptions**](EventsAPI.md#EventsRefIdOptions) | **Options** /events/{refId} | 
[**GetEventByRefId**](EventsAPI.md#GetEventByRefId) | **Get** /events/{refId} | Get event by refId
[**QueryEvents**](EventsAPI.md#QueryEvents) | **Get** /events | Query events by timestamp period and optional event name
[**SendEvents**](EventsAPI.md#SendEvents) | **Post** /events | Submit a batch of events for ingestion
[**SendEventsDryRun**](EventsAPI.md#SendEventsDryRun) | **Post** /events/dry-run | Submit a batch of events for testing



## DeleteEventByRefId

> DeleteEventByRefIdResponse DeleteEventByRefId(ctx, refId).Execute()

Delete an event by refId



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
	refId := "refId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.DeleteEventByRefId(context.Background(), refId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.DeleteEventByRefId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEventByRefId`: DeleteEventByRefIdResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.DeleteEventByRefId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventByRefIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteEventByRefIdResponse**](DeleteEventByRefIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsDryRunOptions

> EventsDryRunOptions(ctx).Execute()



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
	r, err := apiClient.EventsAPI.EventsDryRunOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsDryRunOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventsDryRunOptionsRequest struct via the builder pattern


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


## EventsOptions

> EventsOptions(ctx).Execute()



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
	r, err := apiClient.EventsAPI.EventsOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventsOptionsRequest struct via the builder pattern


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


## EventsRefIdOptions

> EventsRefIdOptions(ctx, refId).Execute()



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
	refId := "refId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.EventsRefIdOptions(context.Background(), refId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRefIdOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRefIdOptionsRequest struct via the builder pattern


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


## GetEventByRefId

> GetEventByRefIdResponse GetEventByRefId(ctx, refId).Execute()

Get event by refId



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
	refId := "refId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetEventByRefId(context.Background(), refId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventByRefId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventByRefId`: GetEventByRefIdResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventByRefId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventByRefIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetEventByRefIdResponse**](GetEventByRefIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryEvents

> QueryEventsResponse QueryEvents(ctx).StartTime(startTime).EndTime(endTime).EventName(eventName).Limit(limit).Cursor(cursor).Execute()

Query events by timestamp period and optional event name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	startTime := time.Now() // time.Time | 
	endTime := time.Now() // time.Time | 
	eventName := "eventName_example" // string |  (optional)
	limit := float32(8.14) // float32 |  (optional) (default to 10)
	cursor := "cursor_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.QueryEvents(context.Background()).StartTime(startTime).EndTime(endTime).EventName(eventName).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.QueryEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryEvents`: QueryEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.QueryEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startTime** | **time.Time** |  | 
 **endTime** | **time.Time** |  | 
 **eventName** | **string** |  | 
 **limit** | **float32** |  | [default to 10]
 **cursor** | **string** |  | 

### Return type

[**QueryEventsResponse**](QueryEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEvents

> SendEventsResponse SendEvents(ctx).SendEventsRequest(sendEventsRequest).Execute()

Submit a batch of events for ingestion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sendEventsRequest := *openapiclient.NewSendEventsRequest([]openapiclient.EventsDryRunRequestEventsInner{*openapiclient.NewEventsDryRunRequestEventsInner("api_call", time.Now(), "customer_12345", "4f6cf35x-2c4y-483z-a0a9-158621f77a21")}) // SendEventsRequest | An array of events following the EventInput schema. This request body should be included in the PUT request to '/events'       Up to 1000 events or a total payload max size of 256KB

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.SendEvents(context.Background()).SendEventsRequest(sendEventsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.SendEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEvents`: SendEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.SendEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendEventsRequest** | [**SendEventsRequest**](SendEventsRequest.md) | An array of events following the EventInput schema. This request body should be included in the PUT request to &#39;/events&#39;       Up to 1000 events or a total payload max size of 256KB | 

### Return type

[**SendEventsResponse**](SendEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEventsDryRun

> []EventsDryRunResponseInner SendEventsDryRun(ctx).EventsDryRunRequest(eventsDryRunRequest).Execute()

Submit a batch of events for testing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	eventsDryRunRequest := *openapiclient.NewEventsDryRunRequest([]openapiclient.EventsDryRunRequestEventsInner{*openapiclient.NewEventsDryRunRequestEventsInner("api_call", time.Now(), "customer_12345", "4f6cf35x-2c4y-483z-a0a9-158621f77a21")}) // EventsDryRunRequest | An array of events following the EventInput schema. This request body should be included in the PUT request to '/events'       Up to 1000 events or a total payload max size of 256KB

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.SendEventsDryRun(context.Background()).EventsDryRunRequest(eventsDryRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.SendEventsDryRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEventsDryRun`: []EventsDryRunResponseInner
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.SendEventsDryRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEventsDryRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventsDryRunRequest** | [**EventsDryRunRequest**](EventsDryRunRequest.md) | An array of events following the EventInput schema. This request body should be included in the PUT request to &#39;/events&#39;       Up to 1000 events or a total payload max size of 256KB | 

### Return type

[**[]EventsDryRunResponseInner**](EventsDryRunResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


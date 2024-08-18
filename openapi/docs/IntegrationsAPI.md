# \IntegrationsAPI

All URIs are relative to *https://connect.withvayu.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportNetSuiteSalesOrder**](IntegrationsAPI.md#ExportNetSuiteSalesOrder) | **Post** /integrations/netsuite/export-sales-order | Export sales order to NetSuite
[**IntegrationsNetsuiteExportSalesOrderOptions**](IntegrationsAPI.md#IntegrationsNetsuiteExportSalesOrderOptions) | **Options** /integrations/netsuite/export-sales-order | 
[**IntegrationsNetsuiteSyncInvoicesOptions**](IntegrationsAPI.md#IntegrationsNetsuiteSyncInvoicesOptions) | **Options** /integrations/netsuite/sync-invoices | 
[**NetSuiteSyncInvoices**](IntegrationsAPI.md#NetSuiteSyncInvoices) | **Post** /integrations/netsuite/sync-invoices | Sync invoices to NetSuite



## ExportNetSuiteSalesOrder

> ExportNetSuiteSalesOrder(ctx).NetSuiteExportSalesOrderRequest(netSuiteExportSalesOrderRequest).Execute()

Export sales order to NetSuite



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
	netSuiteExportSalesOrderRequest := *openapiclient.NewNetSuiteExportSalesOrderRequest("ContractId_example", []string{"ProductsIds_example"}, "SubsidiaryId_example") // NetSuiteExportSalesOrderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationsAPI.ExportNetSuiteSalesOrder(context.Background()).NetSuiteExportSalesOrderRequest(netSuiteExportSalesOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.ExportNetSuiteSalesOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportNetSuiteSalesOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteExportSalesOrderRequest** | [**NetSuiteExportSalesOrderRequest**](NetSuiteExportSalesOrderRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsNetsuiteExportSalesOrderOptions

> IntegrationsNetsuiteExportSalesOrderOptions(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.IntegrationsNetsuiteExportSalesOrderOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsNetsuiteExportSalesOrderOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsNetsuiteExportSalesOrderOptionsRequest struct via the builder pattern


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


## IntegrationsNetsuiteSyncInvoicesOptions

> IntegrationsNetsuiteSyncInvoicesOptions(ctx).Execute()



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
	r, err := apiClient.IntegrationsAPI.IntegrationsNetsuiteSyncInvoicesOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.IntegrationsNetsuiteSyncInvoicesOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsNetsuiteSyncInvoicesOptionsRequest struct via the builder pattern


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


## NetSuiteSyncInvoices

> NetSuiteSyncInvoicesResponse NetSuiteSyncInvoices(ctx).NetSuiteSyncInvoicesRequest(netSuiteSyncInvoicesRequest).Execute()

Sync invoices to NetSuite



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
	netSuiteSyncInvoicesRequest := *openapiclient.NewNetSuiteSyncInvoicesRequest("IntegrationType_example", "Uid_example", *openapiclient.NewNetSuiteSyncInvoicesRequestData(time.Now(), time.Now(), time.Now(), "Memo_example", "ExternalId_example", *openapiclient.NewNetSuiteSyncInvoicesRequestDataEntity("Id_example"), *openapiclient.NewNetSuiteSyncInvoicesRequestDataItem([]openapiclient.NetSuiteSyncInvoicesRequestDataItemItemsInner{*openapiclient.NewNetSuiteSyncInvoicesRequestDataItemItemsInner(*openapiclient.NewNetSuiteSyncInvoicesRequestDataEntity("Id_example"), float32(123), float32(123))}))) // NetSuiteSyncInvoicesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationsAPI.NetSuiteSyncInvoices(context.Background()).NetSuiteSyncInvoicesRequest(netSuiteSyncInvoicesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationsAPI.NetSuiteSyncInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetSuiteSyncInvoices`: NetSuiteSyncInvoicesResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationsAPI.NetSuiteSyncInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetSuiteSyncInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **netSuiteSyncInvoicesRequest** | [**NetSuiteSyncInvoicesRequest**](NetSuiteSyncInvoicesRequest.md) |  | 

### Return type

[**NetSuiteSyncInvoicesResponse**](NetSuiteSyncInvoicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


package api

import (
	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type CustomersAPI struct {
	vayuClient *client.VayuClient
}

type Customer = openapi.CreateCustomerResponseCustomer
type ListCustomersResponse = openapi.ListCustomersResponse
type GetCustomerResponse = openapi.GetCustomerResponse
type CreateCustomerRequest = openapi.CreateCustomerRequest
type CreateCustomerResponse = openapi.CreateCustomerResponse
type UpdateCustomerRequest = openapi.UpdateCustomerRequest
type UpdateCustomerResponse = openapi.UpdateCustomerResponse
type DeleteCustomerResponse = openapi.DeleteCustomerResponse

func NewCustomersAPI(client *client.VayuClient) *CustomersAPI {
	return &CustomersAPI{
		vayuClient: client,
	}
}

func NewCreateCustomerRequest(name string, alias *string) *CreateCustomerRequest {
	return &openapi.CreateCustomerRequest{Name: name, Alias: alias}
}
func NewUpdateCustomerRequest(name *string, alias *string) *UpdateCustomerRequest {
	return &openapi.UpdateCustomerRequest{Name: name, Alias: alias}
}

func (api *CustomersAPI) ListCustomers(limit *float32, cursor *string) (*ListCustomersResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.ListCustomers(ctx)

	if limit != nil {
		request = request.Limit(*limit)
	}

	if cursor != nil {
		request = request.Cursor(*cursor)
	}

	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *CustomersAPI) GetCustomer(customerId string) (*GetCustomerResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.GetCustomer(ctx, customerId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *CustomersAPI) CreateCustomer(payload CreateCustomerRequest) (*CreateCustomerResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.CreateCustomer(ctx)
	request = request.CreateCustomerRequest(payload)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *CustomersAPI) UpdateCustomer(customerId string, payload UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.UpdateCustomer(ctx, customerId)
	request = request.UpdateCustomerRequest(payload)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *CustomersAPI) DeleteCustomer(customerId string) (*DeleteCustomerResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.DeleteCustomer(ctx, customerId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

package api

import (
	"context"
	"time"

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

func (api *CustomersAPI) ListCustomers(limit *float32, cursor *string) (*ListCustomersResponse, *client.VayuError) {
	if invalidLoggedInStatus := api.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

func (api *CustomersAPI) GetCustomer(customerId string) (*GetCustomerResponse, *client.VayuError) {
	if invalidLoggedInStatus := api.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.GetCustomer(ctx, customerId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *CustomersAPI) CreateCustomer(payload CreateCustomerRequest) (*CreateCustomerResponse, *client.VayuError) {
	if invalidLoggedInStatus := api.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.CreateCustomer(ctx)
	request = request.CreateCustomerRequest(payload)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *CustomersAPI) UpdateCustomer(customerId string, payload UpdateCustomerRequest) (*UpdateCustomerResponse, *client.VayuError) {
	if invalidLoggedInStatus := api.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.UpdateCustomer(ctx, customerId)
	request = request.UpdateCustomerRequest(payload)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *CustomersAPI) DeleteCustomer(customerId string) (*DeleteCustomerResponse, *client.VayuError) {
	if invalidLoggedInStatus := api.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := api.vayuClient.Client.CustomersAPI.DeleteCustomer(ctx, customerId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

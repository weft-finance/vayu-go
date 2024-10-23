package api

import (
	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type CustomersAPI struct {
	vayuClient *client.VayuClient
}

type Customer = openapi.CreateCustomerResponseCustomer
type Address = openapi.Address
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

func NewAddress(country *string, city *string, addressText *string, state *string, postalCode *string) *Address {
	return &openapi.Address{
		Country:     country,
		City:        city,
		AddressText: addressText,
		State:       state,
		PostalCode:  postalCode,
	}
}

func NewCreateCustomerRequest(name string, externalId *string, aliases []string, address *Address) *CreateCustomerRequest {
	return &openapi.CreateCustomerRequest{Name: name, ExternalId: externalId, Aliases: aliases, Address: address}
}

func NewUpdateCustomerRequest(name *string, externalId *string, aliases []string, address *Address) *UpdateCustomerRequest {
	return &openapi.UpdateCustomerRequest{Name: name, ExternalId: externalId, Aliases: aliases, Address: address}
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

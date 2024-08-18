package api

import (
	"context"
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type CustomersAPI struct {
	vayuClient *client.VayuClient
}
type ListCustomersResponse = openapi.ListCustomersResponse
type Customer = openapi.CreateCustomerResponseCustomer
type CreateCustomerRequest = openapi.CreateCustomerRequest
type UpdateCustomerRequest = openapi.UpdateCustomerRequest

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

func (c *CustomersAPI) ListCustomers(limit *float32, cursor *string) (*ListCustomersResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("Vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.CustomersAPI.ListCustomers(ctx)

	if limit != nil {
		request = request.Limit(*limit)
	}

	if cursor != nil {
		request = request.Cursor(*cursor)
	}

	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *CustomersAPI) GetCustomer(customerId string) (*Customer, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("Vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.CustomersAPI.GetCustomer(ctx, customerId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return &response.Customer, nil
}

func (c *CustomersAPI) CreateCustomer(payload CreateCustomerRequest) (*Customer, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("Vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.CustomersAPI.CreateCustomer(ctx)
	request = request.CreateCustomerRequest(payload)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return &response.Customer, nil
}

func (c *CustomersAPI) UpdateCustomer(customerId string, payload UpdateCustomerRequest) (*Customer, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("Vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.CustomersAPI.UpdateCustomer(ctx, customerId)
	request = request.UpdateCustomerRequest(payload)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return &response.Customer, nil
}

func (c *CustomersAPI) DeleteCustomer(customerId string) (*Customer, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("Vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.CustomersAPI.DeleteCustomer(ctx, customerId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return &response.Customer, nil
}

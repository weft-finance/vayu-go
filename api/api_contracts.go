package api

import (
	"context"
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type ContractsAPI struct {
	vayuClient *client.VayuClient
}

type Contract = openapi.GetContractResponseContract
type ListContractsResponse = openapi.ListContractsResponse
type CreateContractRequest = openapi.CreateContractRequest
type GetContractResponse = openapi.GetContractResponse
type CreateContractResponse = openapi.CreateContractResponse
type DeleteContractResponse = openapi.DeleteContractResponse

func NewContractsAPI(client *client.VayuClient) *ContractsAPI {
	return &ContractsAPI{
		vayuClient: client,
	}
}
func NewCreateContractRequest(startDate time.Time, endDate *time.Time, customerId string, planId string) *CreateContractRequest {
	return &openapi.CreateContractRequest{StartDate: startDate, EndDate: endDate, CustomerId: customerId, PlanId: planId}
}

func (c *ContractsAPI) ListContracts(limit *float32, cursor *string) (*ListContractsResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.ContractsAPI.ListContracts(ctx)

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

func (c *ContractsAPI) GetContract(contractId string) (*GetContractResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.ContractsAPI.GetContract(ctx, contractId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ContractsAPI) CreateContract(input CreateContractRequest) (*CreateContractResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.ContractsAPI.CreateContract(ctx)
	request = request.CreateContractRequest((openapi.CreateContractRequest)(input))
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ContractsAPI) DeleteContract(contractId string) (*DeleteContractResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.ContractsAPI.DeleteContract(ctx, contractId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

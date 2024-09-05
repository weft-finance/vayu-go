package api

import (
	"context"
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type PlansAPI struct {
	vayuClient *client.VayuClient
}

type Plan = openapi.GetPlanResponsePlan
type ListPlansResponse = openapi.ListPlansResponse
type GetPlanResponse = openapi.GetPlanResponse
type DeletePlanResponse = openapi.DeletePlanResponse

func NewPlansAPI(client *client.VayuClient) *PlansAPI {
	return &PlansAPI{
		vayuClient: client,
	}
}

func (c *PlansAPI) ListPlans(limit *float32, cursor *string) (*ListPlansResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.PlansAPI.ListPlans(ctx)
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

func (c *PlansAPI) GetPlan(planId string) (*GetPlanResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.PlansAPI.GetPlan(ctx, planId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *PlansAPI) DeletePlan(planId string) (*DeletePlanResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.PlansAPI.DeletePlan(ctx, planId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

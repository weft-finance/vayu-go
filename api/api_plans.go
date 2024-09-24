package api

import (
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

func (api *PlansAPI) ListPlans(limit *float32, cursor *string) (*ListPlansResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.PlansAPI.ListPlans(ctx)
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

func (api *PlansAPI) GetPlan(planId string) (*GetPlanResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.PlansAPI.GetPlan(ctx, planId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (api *PlansAPI) DeletePlan(planId string) (*DeletePlanResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.PlansAPI.DeletePlan(ctx, planId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

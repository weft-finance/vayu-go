package api

import (
	"context"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type MetersAPI struct {
	vayuClient *client.VayuClient
}

type Meter = openapi.Meter
type ListMetersResponse = openapi.ListMetersResponse
type GetMeterResponse = openapi.GetMeterResponse

func NewMetersAPI(client *client.VayuClient) *MetersAPI {
	return &MetersAPI{
		vayuClient: client,
	}
}

func (api *MetersAPI) ListMeters(limit *float32, cursor *string) (*ListMetersResponse, *client.VayuError) {
	if invalidLoggedInStatus := api.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := api.vayuClient.Client.MetersAPI.ListMeters(ctx)
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

func (api *MetersAPI) GetMeter(meterId string) (*GetMeterResponse, *client.VayuError) {
	if invalidLoggedInStatus := api.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := api.vayuClient.Client.MetersAPI.GetMeter(ctx, meterId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

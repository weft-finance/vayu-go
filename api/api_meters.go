package api

import (
	"context"
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type MetersAPI struct {
	vayuClient *client.VayuClient
}

type Meter = openapi.GetMeterResponseMeter
type ListMetersResponse = openapi.ListMetersResponse

func NewMetersAPI(client *client.VayuClient) *MetersAPI {
	return &MetersAPI{
		vayuClient: client,
	}
}

func (c *MetersAPI) ListMeters(limit *float32, cursor *string) (*ListMetersResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.MetersAPI.ListMeters(ctx)
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

func (c *MetersAPI) GetMeter(meterId string) (*Meter, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.MetersAPI.GetMeter(ctx, meterId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return &response.Meter, nil
}

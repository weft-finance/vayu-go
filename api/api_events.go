package api

import (
	"context"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type EventsAPI struct {
	vayuClient *client.VayuClient
}

type Event = openapi.Event
type GetEventResponse = openapi.GetEventResponse
type DeleteEventResponse = openapi.DeleteEventResponse
type SendEventsResponse = openapi.SendEventsResponse
type EventsDryRunResponse = openapi.EventsDryRunResponse
type QueryEventsResponse = openapi.QueryEventsResponse

type QueryEventsRequest struct {
	StartTime time.Time
	EndTime   time.Time
	Name      string
	Limit     *float32
	Cursor    *string
}

func NewEventsAPI(client *client.VayuClient) *EventsAPI {
	return &EventsAPI{
		vayuClient: client,
	}
}

func NewQueryEventsRequest(startTime time.Time, endTime time.Time, name string, limit *float32, cursor *string) *QueryEventsRequest {
	return &QueryEventsRequest{
		StartTime: startTime,
		EndTime:   endTime,
		Name:      name,
		Limit:     limit,
		Cursor:    cursor,
	}
}

func (e *EventsAPI) GetEvent(refId string) (*GetEventResponse, *client.VayuError) {
	if invalidLoggedInStatus := e.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.GetEventByRefId(ctx, refId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (e *EventsAPI) DeleteEvent(refId string) (*DeleteEventResponse, *client.VayuError) {
	if invalidLoggedInStatus := e.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.DeleteEventByRefId(ctx, refId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (e *EventsAPI) QueryEvents(payload QueryEventsRequest) (*QueryEventsResponse, *client.VayuError) {
	if invalidLoggedInStatus := e.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.QueryEvents(ctx)

	request = request.StartTime(payload.StartTime)
	request = request.EndTime(payload.EndTime)
	request = request.EventName(payload.Name)

	if payload.Limit != nil {
		request = request.Limit(*payload.Limit)
	}
	if payload.Cursor != nil {
		request = request.Cursor(*payload.Cursor)
	}

	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (e *EventsAPI) SendEvents(events []Event) (*SendEventsResponse, *client.VayuError) {
	if invalidLoggedInStatus := e.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 11115*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.SendEvents(ctx)
	request = request.SendEventsRequest(openapi.SendEventsRequest{Events: events})
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

func (e *EventsAPI) SendEventsDryRun(events []Event) (*EventsDryRunResponse, *client.VayuError) {
	if invalidLoggedInStatus := e.vayuClient.ValidateLoggedIn(); invalidLoggedInStatus != nil {
		return nil, invalidLoggedInStatus
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.SendEventsDryRun(ctx)
	request = request.EventsDryRunRequest(openapi.EventsDryRunRequest{Events: events})
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

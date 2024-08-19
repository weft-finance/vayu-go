package api

import (
	"context"
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type EventsAPI struct {
	vayuClient *client.VayuClient
}

type Event = openapi.QueryEventsResponseEventsInner
type SendEventsResponse = openapi.SendEventsResponse
type SendEventsDryRunResponse = openapi.EventsDryRunResponseInner

func convertGetToEvent(event *openapi.GetEventByRefIdResponseEvent) *Event {
	return &Event{
		Name:          event.Name,
		Timestamp:     event.Timestamp,
		CustomerAlias: event.CustomerAlias,
		Ref:           event.Ref,
		Data:          event.Data,
		Id:            event.Id,
		CreatedAt:     event.CreatedAt,
		UpdatedAt:     event.UpdatedAt,
	}
}

func convertDeleteToEvent(event *openapi.DeleteEventByRefIdResponseEvent) *Event {
	return &Event{
		Name:          event.Name,
		Timestamp:     event.Timestamp,
		CustomerAlias: event.CustomerAlias,
		Ref:           event.Ref,
		Data:          event.Data,
		Id:            event.Id,
		CreatedAt:     event.CreatedAt,
		UpdatedAt:     event.UpdatedAt,
	}
}

type QueryEventsPayload struct {
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

func NewQueryEventsPayload(startTime time.Time, endTime time.Time, name string, limit *float32, cursor *string) *QueryEventsPayload {
	return &QueryEventsPayload{
		StartTime: startTime,
		EndTime:   endTime,
		Name:      name,
		Limit:     limit,
		Cursor:    cursor,
	}
}

func (e *EventsAPI) GetEvent(refId string) (*Event, error) {
	if !e.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.GetEventByRefId(ctx, refId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return convertGetToEvent(&response.Event), nil
}

func (e *EventsAPI) DeleteEvent(refId string) (*Event, error) {
	if !e.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.DeleteEventByRefId(ctx, refId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return convertDeleteToEvent(&response.Event), nil
}

func (e *EventsAPI) QueryEvents(payload QueryEventsPayload) ([]Event, error) {
	if !e.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
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
		return nil, err
	}

	return response.Events, nil
}

func convertEventsToSendEventsPayload(events []Event) []openapi.EventsDryRunRequestEventsInner {
	convertedEvents := make([]openapi.EventsDryRunRequestEventsInner, len(events))
	for i, event := range events {
		convertedEvents[i] = openapi.EventsDryRunRequestEventsInner{
			Name:          event.Name,
			Timestamp:     event.Timestamp,
			CustomerAlias: event.CustomerAlias,
			Ref:           event.Ref,
			Data:          event.Data,
		}
	}
	return convertedEvents
}

func (e *EventsAPI) SendEvents(events []Event) (*SendEventsResponse, error) {
	if !e.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.SendEvents(ctx)
	request = request.SendEventsRequest(openapi.SendEventsRequest{Events: convertEventsToSendEventsPayload(events)})
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (e *EventsAPI) SendEventsDryRun(events []Event) ([]SendEventsDryRunResponse, error) {
	if !e.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := e.vayuClient.Client.EventsAPI.SendEventsDryRun(ctx)
	request = request.EventsDryRunRequest(openapi.EventsDryRunRequest{Events: convertEventsToSendEventsPayload(events)})
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return response, nil
}

package api

import (
	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type WebhooksAPI struct {
	vayuClient *client.VayuClient
}

type Webhook = openapi.ApiWebhookSubscribeRequest
type WebhookSubscribeRequest = openapi.WebhookSubscribeRequest
type NotificationEventType = openapi.NotificationEventType

func NewWebhooksAPI(client *client.VayuClient) *WebhooksAPI {
	return &WebhooksAPI{
		vayuClient: client,
	}
}

func NewWebhookSubscribeRequest(callbackUrl string, eventType NotificationEventType) *WebhookSubscribeRequest {
	return &openapi.WebhookSubscribeRequest{
		CallbackUrl: callbackUrl,
		EventType:   eventType,
	}
}

func (api *PlansAPI) Subscribe(payload WebhookSubscribeRequest) error {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.WebhooksAPI.WebhookSubscribe(ctx)
	request.WebhookSubscribeRequest(payload)

	_, err := request.Execute()

	if err != nil {
		return client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return nil
}

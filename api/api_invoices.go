package api

import (
	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type InvoicesAPI struct {
	vayuClient *client.VayuClient
}

type Invoice = openapi.GetInvoiceResponseInvoice
type GetInvoiceResponse = openapi.GetInvoiceResponse
type ListInvoicesResponse = openapi.ListInvoicesResponse

func NewInvoicesAPI(client *client.VayuClient) *InvoicesAPI {
	return &InvoicesAPI{
		vayuClient: client,
	}
}

func (api *InvoicesAPI) ListInvoices(limit *float32, cursor *string) (*ListInvoicesResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.InvoicesAPI.ListInvoices(ctx)
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

func (api *InvoicesAPI) GetInvoice(invoiceId string) (*GetInvoiceResponse, error) {
	ctx, cancel := client.GenerateContextWithTimeout()
	defer cancel()

	request := api.vayuClient.Client.InvoicesAPI.GetInvoice(ctx, invoiceId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, client.BuildVayuErrorFromGenericOpenAPIError(err)
	}

	return response, nil
}

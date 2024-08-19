package api

import (
	"context"
	"fmt"
	"time"

	"github.com/weft-finance/vayu-go/client"
	"github.com/weft-finance/vayu-go/openapi"
)

type InvoicesAPI struct {
	vayuClient *client.VayuClient
}

type Invoice = openapi.GetInvoiceResponseInvoice
type ListInvoicesResponse = openapi.ListInvoicesResponse

func NewInvoicesAPI(client *client.VayuClient) *InvoicesAPI {
	return &InvoicesAPI{
		vayuClient: client,
	}
}

func (c *InvoicesAPI) ListInvoices(limit *float32, cursor *string) (*ListInvoicesResponse, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.InvoicesAPI.ListInvoices(ctx)
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

func (c *InvoicesAPI) GetInvoice(invoiceId string) (*Invoice, error) {
	if !c.vayuClient.IsLoggedIn() {
		return nil, fmt.Errorf("vayu client is not logged in. please call `vayu.login()` before calling this method")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := c.vayuClient.Client.InvoicesAPI.GetInvoice(ctx, invoiceId)
	response, _, err := request.Execute()

	if err != nil {
		return nil, err
	}

	return &response.Invoice, nil
}

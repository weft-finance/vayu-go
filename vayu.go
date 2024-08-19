package VayuSDK

import (
	"fmt"

	"github.com/weft-finance/vayu-go/api"
	"github.com/weft-finance/vayu-go/client"
)

type Vayu struct {
	client    *client.VayuClient
	Contracts *api.ContractsAPI
	Customers *api.CustomersAPI
	Events    *api.EventsAPI
	Invoices  *api.InvoicesAPI
	Meters    *api.MetersAPI
	Plans     *api.PlansAPI
}

func NewVayu(APIKey string) *Vayu {
	vayuClient := client.NewVayuClient(APIKey)
	return &Vayu{
		client:    vayuClient,
		Contracts: api.NewContractsAPI(vayuClient),
		Customers: api.NewCustomersAPI(vayuClient),
		Events:    api.NewEventsAPI(vayuClient),
		Invoices:  api.NewInvoicesAPI(vayuClient),
		Meters:    api.NewMetersAPI(vayuClient),
		Plans:     api.NewPlansAPI(vayuClient),
	}
}

func (v *Vayu) Login() error {
	return v.client.Login()
}

func (v *Vayu) IsLoggedIn() bool {
	loggedIn := v.client.IsLoggedIn()

	if !loggedIn {
		fmt.Println("vayu client is not logged in. please call `vayu.login()` before calling this method'")
	}

	return loggedIn
}

// Types

type (
	Contract              = api.Contract
	ListContractsResponse = api.ListContractsResponse
	CreateContractRequest = api.CreateContractRequest
)

type (
	Customer              = api.Customer
	ListCustomersResponse = api.ListCustomersResponse
	CreateCustomerRequest = api.CreateCustomerRequest
	UpdateCustomerRequest = api.UpdateCustomerRequest
)

type (
	Event                    = api.Event
	SendEventsResponse       = api.SendEventsResponse
	SendEventsDryRunResponse = api.SendEventsDryRunResponse
	QueryEventsPayload       = api.QueryEventsPayload
)

type (
	Invoice              = api.Invoice
	ListInvoicesResponse = api.ListInvoicesResponse
)

type (
	Meter              = api.Meter
	ListMetersResponse = api.ListMetersResponse
)

type (
	Plan              = api.Plan
	ListPlansResponse = api.ListPlansResponse
)

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
	Webhooks  *api.WebhooksAPI
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
		Webhooks:  api.NewWebhooksAPI(vayuClient),
	}
}

func (v *Vayu) SetCustomHost(host string) {
	v.client.SetCustomHost(host)
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
	VayuError = client.VayuError
)

type (
	Contract               = api.Contract
	ListContractsResponse  = api.ListContractsResponse
	GetContractResponse    = api.GetContractResponse
	CreateContractRequest  = api.CreateContractRequest
	CreateContractResponse = api.CreateContractResponse
	DeleteContractResponse = api.DeleteContractResponse
)

type (
	Customer               = api.Customer
	Address                = api.Address
	ListCustomersResponse  = api.ListCustomersResponse
	GetCustomerResponse    = api.GetCustomerResponse
	CreateCustomerRequest  = api.CreateCustomerRequest
	CreateCustomerResponse = api.CreateCustomerResponse
	UpdateCustomerRequest  = api.UpdateCustomerRequest
	UpdateCustomerResponse = api.UpdateCustomerResponse
	DeleteCustomerResponse = api.DeleteCustomerResponse
)

type (
	Event                = api.Event
	GetEventResponse     = api.GetEventResponse
	DeleteEventResponse  = api.DeleteEventResponse
	SendEventsResponse   = api.SendEventsResponse
	EventsDryRunResponse = api.EventsDryRunResponse
	QueryEventsResponse  = api.QueryEventsResponse
)

type (
	Invoice              = api.Invoice
	GetInvoiceResponse   = api.GetInvoiceResponse
	ListInvoicesResponse = api.ListInvoicesResponse
)

type (
	Meter              = api.Meter
	ListMetersResponse = api.ListMetersResponse
	GetMeterResponse   = api.GetMeterResponse
)

type (
	Plan               = api.Plan
	ListPlansResponse  = api.ListPlansResponse
	GetPlanResponse    = api.GetPlanResponse
	DeletePlanResponse = api.DeletePlanResponse
)

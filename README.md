# Vayu's GO API Client Library


## Overview

The Vayu API client library in GO allows you to submit events for processing and storage, manage billing-related entities, and perform various other operations seamlessly.



## Usage

### Importing

In order to use the Vayu Client Library you would need to import it first into your go project.

```go
import (
	VayuSDK "github.com/weft-finance/vayu-go"
)
```

### Initialization

Initialize the Vayu API client.

```go
    vayu := VayuSDK.NewVayu("VAYU_API_KEY")
```


### Authentication

#### Login and authenticate

```go
err := vayu.Login()

if err != nil {
	panic(err)
}
```


### Events

#### Sending Events

To send a batch of events for processing and storage:

```go
events := []VayuSDK.Event{
    {
        Name:          "api_call",                             // The distinctive label assigned to an event
        Timestamp:     time.Now().UTC(),                       // The exact moment of event occurrence
        CustomerAlias: "customer_alias",                       // A unique identifier assigned to each customer
        Ref:           "4f6cf35x-2c4y-483z-a0a9-158621f77a21", // A universally unique identifier for each event
        Data: map[string]interface{}{
            "key1": "processing_duration", // Example additional data
            "key2": "api_url",             // Example additional data
        },
    },
}

result, err := vayu.Events.SendEvents(events)

if err != nil {
    panic(err)
}

for _, event := range result.ValidEvents {
    println(event.Name)
}
```


#### Querying Events

To fetch events occurring within a specified timestamp range:

```go
limit := float32(10.0)
eventsQuery := VayuSDK.QueryEventsRequest{
    StartTime: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
    EndTime:   time.Date(2024, 9, 30, 23, 59, 59, 999, time.UTC),
    Name:      "api_call",
    Limit:    &float32(10.0),
}
result, err := vayu.Events.QueryEvents(eventsQuery)

if err != nil {
    panic(err)
}

for _, event := range result {
    println(event.Name)
}
```


#### Getting Event by Ref ID

To get a specific event using its reference ID:

```go
event, err := vayu.Events.GetEvent("4f6cf35x-2c4y-483z-a0a9-158621f77a21")
if err != nil {
    panic(err)
}
println(event.Name)
```

#### Deleting Event by Ref ID

To delete a specific event using its reference ID:

```go
event, err := vayu.Events.DeleteEvent("4f6cf35x-2c4y-483z-a0a9-158621f77a21")
if err != nil {
    panic(err)
}
println(event.Name)
```


### Customers

#### Creating a Customer

To create a new customer:

```go
alias := "customer_12345"
customerPayload := VayuSDK.CreateCustomerRequest{
    Name: "JJane Doe",
    Alias: &alias,
}
customer, err := vayu.Customers.CreateCustomer(customerPayload)
if err != nil {
    panic(err)
}
println(customer.Id)
```

#### Updating a Customer

To update an existing customer by ID:

```go
newAlias := "customer_67890"
updateCustomerPayload := VayuSDK.UpdateCustomerRequest{
    Alias: &newAlias,
}
updatedCustomer, err := vayu.Customers.UpdateCustomer("customer-id", updateCustomerPayload)
if err != nil {
    panic(err)
}
println(updatedCustomer.Alias)
```

#### Deleting a Customer

To delete a customer by ID:

```go
customer, err := vayu.Customers.DeleteCustomer("customer-id")
if err != nil {
    panic(err)
}
println(customer.Id)
```

### Contracts

#### Assigning a contract to a customer

In order to assign a contract to a customer you would need to provide the customer Id and the relevant plan


    customerId: '1f4cf23x-2c4y-483z-1111-158621f77a21',
    planId: '4f6cf35x-1234-483z-a0a9-158621f77a21',

```go
	contract, err := vayu.Contracts.CreateContract(VayuSDK.CreateContractRequest{
		StartDate:  time.Now().UTC(),                       // The start date of the contract
		EndDate:    nil,                                    // The end date of the contract
		CustomerId: "1f4cf23x-2c4y-483z-1111-158621f77a21", // The id of the customer that the contract is associated with
		PlanId:     "4f6cf35x-1234-483z-a0a9-158621f77a21", // The id of the plan that the contract is associated with
	})
```

### Meters

Meters are entities that track and aggregate usage data based on events. They are crucial for billing and monitoring purposes.

#### Listing Meters

To get a list of meters:

```go
limit := float32(10)
cursor := "cursor"
meters, err := vayu.Meters.ListMeters(&limit, &cursor)
if err != nil {
    panic(err)
}
for _, meter := range meters.Meters {
    println(meter.Id)
}
```

## Features

The Vayu API client library provides access to the following features:

- **Auth**
  - `Login()`
- **Events**
  - `Events.SendEvents`
  - `Events.QueryEvents`
  - `Events.GetEvent`
  - `Events.DeleteEvent`
  - `Events.SendEventsDryRun`
- **Customers**
  - `Customers.CreateCustomer`
  - `Customers.UpdateCustomer`
  - `Customers.DeleteCustomer`
  - `Customers.ListCustomers`
  - `Customers.GetCustomer`
- **Meters**
  - `Meters.GetMeter`
  - `Meters.UpdateMeter`
  - `Meters.DeleteMeter`
  - `Meters.ListMeters`
- **Plans**
  - `Plans.GetPlan`
  - `Plans.DeletePlan`
  - `Plans.ListPlans`
- **Contracts**
  - `Contracts.CreateContract`
  - `Contracts.DeleteContract`
  - `Contracts.ListContracts`
  - `Contracts.GetContract`
- **Invoices**
  - `Invoices.GetInvoice`
  - `Invoices.ListInvoices`

## Support

If you have any questions or need further assistance, please contact Vayu at `team@withvayu.com`.

## License

This project is licensed under the MIT License.

---

This README provides an overview and usage examples for the Vayu API client library. For more detailed information on each method, please refer to the official Vayu API documentation.
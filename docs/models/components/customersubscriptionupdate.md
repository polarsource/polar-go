# CustomerSubscriptionUpdate


## Supported Types

### CustomerSubscriptionUpdateProduct

```go
customerSubscriptionUpdate := components.CreateCustomerSubscriptionUpdateCustomerSubscriptionUpdateProduct(components.CustomerSubscriptionUpdateProduct{/* values here */})
```

### CustomerSubscriptionUpdateSeats

```go
customerSubscriptionUpdate := components.CreateCustomerSubscriptionUpdateCustomerSubscriptionUpdateSeats(components.CustomerSubscriptionUpdateSeats{/* values here */})
```

### CustomerSubscriptionCancel

```go
customerSubscriptionUpdate := components.CreateCustomerSubscriptionUpdateCustomerSubscriptionCancel(components.CustomerSubscriptionCancel{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerSubscriptionUpdate.Type {
	case components.CustomerSubscriptionUpdateTypeCustomerSubscriptionUpdateProduct:
		// customerSubscriptionUpdate.CustomerSubscriptionUpdateProduct is populated
	case components.CustomerSubscriptionUpdateTypeCustomerSubscriptionUpdateSeats:
		// customerSubscriptionUpdate.CustomerSubscriptionUpdateSeats is populated
	case components.CustomerSubscriptionUpdateTypeCustomerSubscriptionCancel:
		// customerSubscriptionUpdate.CustomerSubscriptionCancel is populated
}
```

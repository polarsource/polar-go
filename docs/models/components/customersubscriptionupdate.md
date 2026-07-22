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

### CustomerSubscriptionPause

```go
customerSubscriptionUpdate := components.CreateCustomerSubscriptionUpdateCustomerSubscriptionPause(components.CustomerSubscriptionPause{/* values here */})
```

### CustomerSubscriptionResume

```go
customerSubscriptionUpdate := components.CreateCustomerSubscriptionUpdateCustomerSubscriptionResume(components.CustomerSubscriptionResume{/* values here */})
```

### CustomerSubscriptionUpdateClear

```go
customerSubscriptionUpdate := components.CreateCustomerSubscriptionUpdateCustomerSubscriptionUpdateClear(components.CustomerSubscriptionUpdateClear{/* values here */})
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
	case components.CustomerSubscriptionUpdateTypeCustomerSubscriptionPause:
		// customerSubscriptionUpdate.CustomerSubscriptionPause is populated
	case components.CustomerSubscriptionUpdateTypeCustomerSubscriptionResume:
		// customerSubscriptionUpdate.CustomerSubscriptionResume is populated
	case components.CustomerSubscriptionUpdateTypeCustomerSubscriptionUpdateClear:
		// customerSubscriptionUpdate.CustomerSubscriptionUpdateClear is populated
}
```

# SubscriptionUpdate


## Supported Types

### SubscriptionUpdateProduct

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateProduct(components.SubscriptionUpdateProduct{/* values here */})
```

### SubscriptionUpdateDiscount

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateDiscount(components.SubscriptionUpdateDiscount{/* values here */})
```

### SubscriptionUpdateTrial

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateTrial(components.SubscriptionUpdateTrial{/* values here */})
```

### SubscriptionUpdateSeats

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateSeats(components.SubscriptionUpdateSeats{/* values here */})
```

### SubscriptionUpdateBillingPeriod

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateBillingPeriod(components.SubscriptionUpdateBillingPeriod{/* values here */})
```

### SubscriptionCancel

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionCancel(components.SubscriptionCancel{/* values here */})
```

### SubscriptionRevoke

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionRevoke(components.SubscriptionRevoke{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionUpdate.Type {
	case components.SubscriptionUpdateTypeSubscriptionUpdateProduct:
		// subscriptionUpdate.SubscriptionUpdateProduct is populated
	case components.SubscriptionUpdateTypeSubscriptionUpdateDiscount:
		// subscriptionUpdate.SubscriptionUpdateDiscount is populated
	case components.SubscriptionUpdateTypeSubscriptionUpdateTrial:
		// subscriptionUpdate.SubscriptionUpdateTrial is populated
	case components.SubscriptionUpdateTypeSubscriptionUpdateSeats:
		// subscriptionUpdate.SubscriptionUpdateSeats is populated
	case components.SubscriptionUpdateTypeSubscriptionUpdateBillingPeriod:
		// subscriptionUpdate.SubscriptionUpdateBillingPeriod is populated
	case components.SubscriptionUpdateTypeSubscriptionCancel:
		// subscriptionUpdate.SubscriptionCancel is populated
	case components.SubscriptionUpdateTypeSubscriptionRevoke:
		// subscriptionUpdate.SubscriptionRevoke is populated
}
```

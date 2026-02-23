# SubscriptionsCreateSubscriptionCreate


## Supported Types

### SubscriptionCreateCustomer

```go
subscriptionsCreateSubscriptionCreate := operations.CreateSubscriptionsCreateSubscriptionCreateSubscriptionCreateCustomer(components.SubscriptionCreateCustomer{/* values here */})
```

### SubscriptionCreateExternalCustomer

```go
subscriptionsCreateSubscriptionCreate := operations.CreateSubscriptionsCreateSubscriptionCreateSubscriptionCreateExternalCustomer(components.SubscriptionCreateExternalCustomer{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionsCreateSubscriptionCreate.Type {
	case operations.SubscriptionsCreateSubscriptionCreateTypeSubscriptionCreateCustomer:
		// subscriptionsCreateSubscriptionCreate.SubscriptionCreateCustomer is populated
	case operations.SubscriptionsCreateSubscriptionCreateTypeSubscriptionCreateExternalCustomer:
		// subscriptionsCreateSubscriptionCreate.SubscriptionCreateExternalCustomer is populated
}
```

# SubscriptionPrices


## Supported Types

### LegacyRecurringProductPrice

```go
subscriptionPrices := components.CreateSubscriptionPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
subscriptionPrices := components.CreateSubscriptionPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionPrices.Type {
	case components.SubscriptionPricesTypeLegacyRecurringProductPrice:
		// subscriptionPrices.LegacyRecurringProductPrice is populated
	case components.SubscriptionPricesTypeProductPrice:
		// subscriptionPrices.ProductPrice is populated
}
```

# CustomerSubscriptionPrices


## Supported Types

### LegacyRecurringProductPrice

```go
customerSubscriptionPrices := components.CreateCustomerSubscriptionPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
customerSubscriptionPrices := components.CreateCustomerSubscriptionPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerSubscriptionPrices.Type {
	case components.CustomerSubscriptionPricesTypeLegacyRecurringProductPrice:
		// customerSubscriptionPrices.LegacyRecurringProductPrice is populated
	case components.CustomerSubscriptionPricesTypeProductPrice:
		// customerSubscriptionPrices.ProductPrice is populated
}
```

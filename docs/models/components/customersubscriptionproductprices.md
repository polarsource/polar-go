# CustomerSubscriptionProductPrices


## Supported Types

### LegacyRecurringProductPrice

```go
customerSubscriptionProductPrices := components.CreateCustomerSubscriptionProductPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
customerSubscriptionProductPrices := components.CreateCustomerSubscriptionProductPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerSubscriptionProductPrices.Type {
	case components.CustomerSubscriptionProductPricesTypeLegacyRecurringProductPrice:
		// customerSubscriptionProductPrices.LegacyRecurringProductPrice is populated
	case components.CustomerSubscriptionProductPricesTypeProductPrice:
		// customerSubscriptionProductPrices.ProductPrice is populated
}
```

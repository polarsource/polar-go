# CheckoutPrices


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutPrices := components.CreateCheckoutPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutPrices := components.CreateCheckoutPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPrices.Type {
	case components.CheckoutPricesTypeLegacyRecurringProductPrice:
		// checkoutPrices.LegacyRecurringProductPrice is populated
	case components.CheckoutPricesTypeProductPrice:
		// checkoutPrices.ProductPrice is populated
}
```

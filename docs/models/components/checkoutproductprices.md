# CheckoutProductPrices


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutProductPrices := components.CreateCheckoutProductPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutProductPrices := components.CreateCheckoutProductPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutProductPrices.Type {
	case components.CheckoutProductPricesTypeLegacyRecurringProductPrice:
		// checkoutProductPrices.LegacyRecurringProductPrice is populated
	case components.CheckoutProductPricesTypeProductPrice:
		// checkoutProductPrices.ProductPrice is populated
}
```

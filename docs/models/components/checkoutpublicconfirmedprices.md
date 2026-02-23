# CheckoutPublicConfirmedPrices


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutPublicConfirmedPrices := components.CreateCheckoutPublicConfirmedPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutPublicConfirmedPrices := components.CreateCheckoutPublicConfirmedPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicConfirmedPrices.Type {
	case components.CheckoutPublicConfirmedPricesTypeLegacyRecurringProductPrice:
		// checkoutPublicConfirmedPrices.LegacyRecurringProductPrice is populated
	case components.CheckoutPublicConfirmedPricesTypeProductPrice:
		// checkoutPublicConfirmedPrices.ProductPrice is populated
}
```

# CheckoutPublicPrices


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutPublicPrices := components.CreateCheckoutPublicPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutPublicPrices := components.CreateCheckoutPublicPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicPrices.Type {
	case components.CheckoutPublicPricesTypeLegacyRecurringProductPrice:
		// checkoutPublicPrices.LegacyRecurringProductPrice is populated
	case components.CheckoutPublicPricesTypeProductPrice:
		// checkoutPublicPrices.ProductPrice is populated
}
```

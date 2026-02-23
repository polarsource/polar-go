# CheckoutLinkProductPrices


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutLinkProductPrices := components.CreateCheckoutLinkProductPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutLinkProductPrices := components.CreateCheckoutLinkProductPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinkProductPrices.Type {
	case components.CheckoutLinkProductPricesTypeLegacyRecurringProductPrice:
		// checkoutLinkProductPrices.LegacyRecurringProductPrice is populated
	case components.CheckoutLinkProductPricesTypeProductPrice:
		// checkoutLinkProductPrices.ProductPrice is populated
}
```

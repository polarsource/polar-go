# CheckoutProductPrice


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutProductPrice := components.CreateCheckoutProductPriceLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutProductPrice := components.CreateCheckoutProductPriceProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutProductPrice.Type {
	case components.CheckoutProductPriceTypeLegacyRecurringProductPrice:
		// checkoutProductPrice.LegacyRecurringProductPrice is populated
	case components.CheckoutProductPriceTypeProductPrice:
		// checkoutProductPrice.ProductPrice is populated
}
```

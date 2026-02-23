# CheckoutPublicProductPrice


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutPublicProductPrice := components.CreateCheckoutPublicProductPriceLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutPublicProductPrice := components.CreateCheckoutPublicProductPriceProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicProductPrice.Type {
	case components.CheckoutPublicProductPriceTypeLegacyRecurringProductPrice:
		// checkoutPublicProductPrice.LegacyRecurringProductPrice is populated
	case components.CheckoutPublicProductPriceTypeProductPrice:
		// checkoutPublicProductPrice.ProductPrice is populated
}
```

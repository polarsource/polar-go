# CheckoutPublicConfirmedProductPrice


## Supported Types

### LegacyRecurringProductPrice

```go
checkoutPublicConfirmedProductPrice := components.CreateCheckoutPublicConfirmedProductPriceLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
checkoutPublicConfirmedProductPrice := components.CreateCheckoutPublicConfirmedProductPriceProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicConfirmedProductPrice.Type {
	case components.CheckoutPublicConfirmedProductPriceTypeLegacyRecurringProductPrice:
		// checkoutPublicConfirmedProductPrice.LegacyRecurringProductPrice is populated
	case components.CheckoutPublicConfirmedProductPriceTypeProductPrice:
		// checkoutPublicConfirmedProductPrice.ProductPrice is populated
}
```

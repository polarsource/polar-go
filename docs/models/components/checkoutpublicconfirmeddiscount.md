# CheckoutPublicConfirmedDiscount


## Supported Types

### CheckoutDiscountFixedOnceForeverDuration

```go
checkoutPublicConfirmedDiscount := components.CreateCheckoutPublicConfirmedDiscountCheckoutDiscountFixedOnceForeverDuration(components.CheckoutDiscountFixedOnceForeverDuration{/* values here */})
```

### CheckoutDiscountFixedRepeatDuration

```go
checkoutPublicConfirmedDiscount := components.CreateCheckoutPublicConfirmedDiscountCheckoutDiscountFixedRepeatDuration(components.CheckoutDiscountFixedRepeatDuration{/* values here */})
```

### CheckoutDiscountPercentageOnceForeverDuration

```go
checkoutPublicConfirmedDiscount := components.CreateCheckoutPublicConfirmedDiscountCheckoutDiscountPercentageOnceForeverDuration(components.CheckoutDiscountPercentageOnceForeverDuration{/* values here */})
```

### CheckoutDiscountPercentageRepeatDuration

```go
checkoutPublicConfirmedDiscount := components.CreateCheckoutPublicConfirmedDiscountCheckoutDiscountPercentageRepeatDuration(components.CheckoutDiscountPercentageRepeatDuration{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicConfirmedDiscount.Type {
	case components.CheckoutPublicConfirmedDiscountTypeCheckoutDiscountFixedOnceForeverDuration:
		// checkoutPublicConfirmedDiscount.CheckoutDiscountFixedOnceForeverDuration is populated
	case components.CheckoutPublicConfirmedDiscountTypeCheckoutDiscountFixedRepeatDuration:
		// checkoutPublicConfirmedDiscount.CheckoutDiscountFixedRepeatDuration is populated
	case components.CheckoutPublicConfirmedDiscountTypeCheckoutDiscountPercentageOnceForeverDuration:
		// checkoutPublicConfirmedDiscount.CheckoutDiscountPercentageOnceForeverDuration is populated
	case components.CheckoutPublicConfirmedDiscountTypeCheckoutDiscountPercentageRepeatDuration:
		// checkoutPublicConfirmedDiscount.CheckoutDiscountPercentageRepeatDuration is populated
}
```

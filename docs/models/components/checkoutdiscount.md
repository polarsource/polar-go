# CheckoutDiscount


## Supported Types

### CheckoutDiscountFixedOnceForeverDuration

```go
checkoutDiscount := components.CreateCheckoutDiscountCheckoutDiscountFixedOnceForeverDuration(components.CheckoutDiscountFixedOnceForeverDuration{/* values here */})
```

### CheckoutDiscountFixedRepeatDuration

```go
checkoutDiscount := components.CreateCheckoutDiscountCheckoutDiscountFixedRepeatDuration(components.CheckoutDiscountFixedRepeatDuration{/* values here */})
```

### CheckoutDiscountPercentageOnceForeverDuration

```go
checkoutDiscount := components.CreateCheckoutDiscountCheckoutDiscountPercentageOnceForeverDuration(components.CheckoutDiscountPercentageOnceForeverDuration{/* values here */})
```

### CheckoutDiscountPercentageRepeatDuration

```go
checkoutDiscount := components.CreateCheckoutDiscountCheckoutDiscountPercentageRepeatDuration(components.CheckoutDiscountPercentageRepeatDuration{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutDiscount.Type {
	case components.CheckoutDiscountTypeCheckoutDiscountFixedOnceForeverDuration:
		// checkoutDiscount.CheckoutDiscountFixedOnceForeverDuration is populated
	case components.CheckoutDiscountTypeCheckoutDiscountFixedRepeatDuration:
		// checkoutDiscount.CheckoutDiscountFixedRepeatDuration is populated
	case components.CheckoutDiscountTypeCheckoutDiscountPercentageOnceForeverDuration:
		// checkoutDiscount.CheckoutDiscountPercentageOnceForeverDuration is populated
	case components.CheckoutDiscountTypeCheckoutDiscountPercentageRepeatDuration:
		// checkoutDiscount.CheckoutDiscountPercentageRepeatDuration is populated
}
```

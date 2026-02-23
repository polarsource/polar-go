# CheckoutPublicDiscount


## Supported Types

### CheckoutDiscountFixedOnceForeverDuration

```go
checkoutPublicDiscount := components.CreateCheckoutPublicDiscountCheckoutDiscountFixedOnceForeverDuration(components.CheckoutDiscountFixedOnceForeverDuration{/* values here */})
```

### CheckoutDiscountFixedRepeatDuration

```go
checkoutPublicDiscount := components.CreateCheckoutPublicDiscountCheckoutDiscountFixedRepeatDuration(components.CheckoutDiscountFixedRepeatDuration{/* values here */})
```

### CheckoutDiscountPercentageOnceForeverDuration

```go
checkoutPublicDiscount := components.CreateCheckoutPublicDiscountCheckoutDiscountPercentageOnceForeverDuration(components.CheckoutDiscountPercentageOnceForeverDuration{/* values here */})
```

### CheckoutDiscountPercentageRepeatDuration

```go
checkoutPublicDiscount := components.CreateCheckoutPublicDiscountCheckoutDiscountPercentageRepeatDuration(components.CheckoutDiscountPercentageRepeatDuration{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicDiscount.Type {
	case components.CheckoutPublicDiscountTypeCheckoutDiscountFixedOnceForeverDuration:
		// checkoutPublicDiscount.CheckoutDiscountFixedOnceForeverDuration is populated
	case components.CheckoutPublicDiscountTypeCheckoutDiscountFixedRepeatDuration:
		// checkoutPublicDiscount.CheckoutDiscountFixedRepeatDuration is populated
	case components.CheckoutPublicDiscountTypeCheckoutDiscountPercentageOnceForeverDuration:
		// checkoutPublicDiscount.CheckoutDiscountPercentageOnceForeverDuration is populated
	case components.CheckoutPublicDiscountTypeCheckoutDiscountPercentageRepeatDuration:
		// checkoutPublicDiscount.CheckoutDiscountPercentageRepeatDuration is populated
}
```

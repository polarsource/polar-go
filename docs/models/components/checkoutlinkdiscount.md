# CheckoutLinkDiscount


## Supported Types

### DiscountFixedOnceForeverDurationBase

```go
checkoutLinkDiscount := components.CreateCheckoutLinkDiscountDiscountFixedOnceForeverDurationBase(components.DiscountFixedOnceForeverDurationBase{/* values here */})
```

### DiscountFixedRepeatDurationBase

```go
checkoutLinkDiscount := components.CreateCheckoutLinkDiscountDiscountFixedRepeatDurationBase(components.DiscountFixedRepeatDurationBase{/* values here */})
```

### DiscountPercentageOnceForeverDurationBase

```go
checkoutLinkDiscount := components.CreateCheckoutLinkDiscountDiscountPercentageOnceForeverDurationBase(components.DiscountPercentageOnceForeverDurationBase{/* values here */})
```

### DiscountPercentageRepeatDurationBase

```go
checkoutLinkDiscount := components.CreateCheckoutLinkDiscountDiscountPercentageRepeatDurationBase(components.DiscountPercentageRepeatDurationBase{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinkDiscount.Type {
	case components.CheckoutLinkDiscountTypeDiscountFixedOnceForeverDurationBase:
		// checkoutLinkDiscount.DiscountFixedOnceForeverDurationBase is populated
	case components.CheckoutLinkDiscountTypeDiscountFixedRepeatDurationBase:
		// checkoutLinkDiscount.DiscountFixedRepeatDurationBase is populated
	case components.CheckoutLinkDiscountTypeDiscountPercentageOnceForeverDurationBase:
		// checkoutLinkDiscount.DiscountPercentageOnceForeverDurationBase is populated
	case components.CheckoutLinkDiscountTypeDiscountPercentageRepeatDurationBase:
		// checkoutLinkDiscount.DiscountPercentageRepeatDurationBase is populated
}
```

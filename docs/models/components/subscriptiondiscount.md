# SubscriptionDiscount


## Supported Types

### DiscountFixedOnceForeverDurationBase

```go
subscriptionDiscount := components.CreateSubscriptionDiscountDiscountFixedOnceForeverDurationBase(components.DiscountFixedOnceForeverDurationBase{/* values here */})
```

### DiscountFixedRepeatDurationBase

```go
subscriptionDiscount := components.CreateSubscriptionDiscountDiscountFixedRepeatDurationBase(components.DiscountFixedRepeatDurationBase{/* values here */})
```

### DiscountPercentageOnceForeverDurationBase

```go
subscriptionDiscount := components.CreateSubscriptionDiscountDiscountPercentageOnceForeverDurationBase(components.DiscountPercentageOnceForeverDurationBase{/* values here */})
```

### DiscountPercentageRepeatDurationBase

```go
subscriptionDiscount := components.CreateSubscriptionDiscountDiscountPercentageRepeatDurationBase(components.DiscountPercentageRepeatDurationBase{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionDiscount.Type {
	case components.SubscriptionDiscountTypeDiscountFixedOnceForeverDurationBase:
		// subscriptionDiscount.DiscountFixedOnceForeverDurationBase is populated
	case components.SubscriptionDiscountTypeDiscountFixedRepeatDurationBase:
		// subscriptionDiscount.DiscountFixedRepeatDurationBase is populated
	case components.SubscriptionDiscountTypeDiscountPercentageOnceForeverDurationBase:
		// subscriptionDiscount.DiscountPercentageOnceForeverDurationBase is populated
	case components.SubscriptionDiscountTypeDiscountPercentageRepeatDurationBase:
		// subscriptionDiscount.DiscountPercentageRepeatDurationBase is populated
}
```

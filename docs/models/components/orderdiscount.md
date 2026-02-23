# OrderDiscount


## Supported Types

### DiscountFixedOnceForeverDurationBase

```go
orderDiscount := components.CreateOrderDiscountDiscountFixedOnceForeverDurationBase(components.DiscountFixedOnceForeverDurationBase{/* values here */})
```

### DiscountFixedRepeatDurationBase

```go
orderDiscount := components.CreateOrderDiscountDiscountFixedRepeatDurationBase(components.DiscountFixedRepeatDurationBase{/* values here */})
```

### DiscountPercentageOnceForeverDurationBase

```go
orderDiscount := components.CreateOrderDiscountDiscountPercentageOnceForeverDurationBase(components.DiscountPercentageOnceForeverDurationBase{/* values here */})
```

### DiscountPercentageRepeatDurationBase

```go
orderDiscount := components.CreateOrderDiscountDiscountPercentageRepeatDurationBase(components.DiscountPercentageRepeatDurationBase{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch orderDiscount.Type {
	case components.OrderDiscountTypeDiscountFixedOnceForeverDurationBase:
		// orderDiscount.DiscountFixedOnceForeverDurationBase is populated
	case components.OrderDiscountTypeDiscountFixedRepeatDurationBase:
		// orderDiscount.DiscountFixedRepeatDurationBase is populated
	case components.OrderDiscountTypeDiscountPercentageOnceForeverDurationBase:
		// orderDiscount.DiscountPercentageOnceForeverDurationBase is populated
	case components.OrderDiscountTypeDiscountPercentageRepeatDurationBase:
		// orderDiscount.DiscountPercentageRepeatDurationBase is populated
}
```

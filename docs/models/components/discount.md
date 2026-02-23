# Discount


## Supported Types

### DiscountFixedOnceForeverDuration

```go
discount := components.CreateDiscountDiscountFixedOnceForeverDuration(components.DiscountFixedOnceForeverDuration{/* values here */})
```

### DiscountFixedRepeatDuration

```go
discount := components.CreateDiscountDiscountFixedRepeatDuration(components.DiscountFixedRepeatDuration{/* values here */})
```

### DiscountPercentageOnceForeverDuration

```go
discount := components.CreateDiscountDiscountPercentageOnceForeverDuration(components.DiscountPercentageOnceForeverDuration{/* values here */})
```

### DiscountPercentageRepeatDuration

```go
discount := components.CreateDiscountDiscountPercentageRepeatDuration(components.DiscountPercentageRepeatDuration{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discount.Type {
	case components.DiscountUnionTypeDiscountFixedOnceForeverDuration:
		// discount.DiscountFixedOnceForeverDuration is populated
	case components.DiscountUnionTypeDiscountFixedRepeatDuration:
		// discount.DiscountFixedRepeatDuration is populated
	case components.DiscountUnionTypeDiscountPercentageOnceForeverDuration:
		// discount.DiscountPercentageOnceForeverDuration is populated
	case components.DiscountUnionTypeDiscountPercentageRepeatDuration:
		// discount.DiscountPercentageRepeatDuration is populated
}
```

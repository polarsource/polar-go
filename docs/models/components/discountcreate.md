# DiscountCreate


## Supported Types

### DiscountFixedOnceForeverDurationCreate

```go
discountCreate := components.CreateDiscountCreateDiscountFixedOnceForeverDurationCreate(components.DiscountFixedOnceForeverDurationCreate{/* values here */})
```

### DiscountFixedRepeatDurationCreate

```go
discountCreate := components.CreateDiscountCreateDiscountFixedRepeatDurationCreate(components.DiscountFixedRepeatDurationCreate{/* values here */})
```

### DiscountPercentageOnceForeverDurationCreate

```go
discountCreate := components.CreateDiscountCreateDiscountPercentageOnceForeverDurationCreate(components.DiscountPercentageOnceForeverDurationCreate{/* values here */})
```

### DiscountPercentageRepeatDurationCreate

```go
discountCreate := components.CreateDiscountCreateDiscountPercentageRepeatDurationCreate(components.DiscountPercentageRepeatDurationCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountCreate.Type {
	case components.DiscountCreateTypeDiscountFixedOnceForeverDurationCreate:
		// discountCreate.DiscountFixedOnceForeverDurationCreate is populated
	case components.DiscountCreateTypeDiscountFixedRepeatDurationCreate:
		// discountCreate.DiscountFixedRepeatDurationCreate is populated
	case components.DiscountCreateTypeDiscountPercentageOnceForeverDurationCreate:
		// discountCreate.DiscountPercentageOnceForeverDurationCreate is populated
	case components.DiscountCreateTypeDiscountPercentageRepeatDurationCreate:
		// discountCreate.DiscountPercentageRepeatDurationCreate is populated
}
```

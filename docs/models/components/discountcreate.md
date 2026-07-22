# DiscountCreate


## Supported Types

### DiscountFixedCreate

```go
discountCreate := components.CreateDiscountCreateFixed(components.DiscountFixedCreate{/* values here */})
```

### DiscountPercentageCreate

```go
discountCreate := components.CreateDiscountCreatePercentage(components.DiscountPercentageCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountCreate.Type {
	case components.DiscountCreateTypeFixed:
		// discountCreate.DiscountFixedCreate is populated
	case components.DiscountCreateTypePercentage:
		// discountCreate.DiscountPercentageCreate is populated
}
```

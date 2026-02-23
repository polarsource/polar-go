# Two


## Supported Types

### ProductPriceCustomCreate

```go
two := components.CreateTwoCustom(components.ProductPriceCustomCreate{/* values here */})
```

### ProductPriceFixedCreate

```go
two := components.CreateTwoFixed(components.ProductPriceFixedCreate{/* values here */})
```

### ProductPriceFreeCreate

```go
two := components.CreateTwoFree(components.ProductPriceFreeCreate{/* values here */})
```

### ProductPriceMeteredUnitCreate

```go
two := components.CreateTwoMeteredUnit(components.ProductPriceMeteredUnitCreate{/* values here */})
```

### ProductPriceSeatBasedCreate

```go
two := components.CreateTwoSeatBased(components.ProductPriceSeatBasedCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch two.Type {
	case components.TwoTypeCustom:
		// two.ProductPriceCustomCreate is populated
	case components.TwoTypeFixed:
		// two.ProductPriceFixedCreate is populated
	case components.TwoTypeFree:
		// two.ProductPriceFreeCreate is populated
	case components.TwoTypeMeteredUnit:
		// two.ProductPriceMeteredUnitCreate is populated
	case components.TwoTypeSeatBased:
		// two.ProductPriceSeatBasedCreate is populated
}
```

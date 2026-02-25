# ProductPrice


## Supported Types

### ProductPriceCustom

```go
productPrice := components.CreateProductPriceCustom(components.ProductPriceCustom{/* values here */})
```

### ProductPriceFixed

```go
productPrice := components.CreateProductPriceFixed(components.ProductPriceFixed{/* values here */})
```

### ProductPriceFree

```go
productPrice := components.CreateProductPriceFree(components.ProductPriceFree{/* values here */})
```

### ProductPriceMeteredUnit

```go
productPrice := components.CreateProductPriceMeteredUnit(components.ProductPriceMeteredUnit{/* values here */})
```

### ProductPriceSeatBased

```go
productPrice := components.CreateProductPriceSeatBased(components.ProductPriceSeatBased{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productPrice.Type {
	case components.ProductPriceTypeCustom:
		// productPrice.ProductPriceCustom is populated
	case components.ProductPriceTypeFixed:
		// productPrice.ProductPriceFixed is populated
	case components.ProductPriceTypeFree:
		// productPrice.ProductPriceFree is populated
	case components.ProductPriceTypeMeteredUnit:
		// productPrice.ProductPriceMeteredUnit is populated
	case components.ProductPriceTypeSeatBased:
		// productPrice.ProductPriceSeatBased is populated
}
```

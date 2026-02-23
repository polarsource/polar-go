# ProductCreateOneTimePrices


## Supported Types

### ProductPriceCustomCreate

```go
productCreateOneTimePrices := components.CreateProductCreateOneTimePricesCustom(components.ProductPriceCustomCreate{/* values here */})
```

### ProductPriceFixedCreate

```go
productCreateOneTimePrices := components.CreateProductCreateOneTimePricesFixed(components.ProductPriceFixedCreate{/* values here */})
```

### ProductPriceFreeCreate

```go
productCreateOneTimePrices := components.CreateProductCreateOneTimePricesFree(components.ProductPriceFreeCreate{/* values here */})
```

### ProductPriceMeteredUnitCreate

```go
productCreateOneTimePrices := components.CreateProductCreateOneTimePricesMeteredUnit(components.ProductPriceMeteredUnitCreate{/* values here */})
```

### ProductPriceSeatBasedCreate

```go
productCreateOneTimePrices := components.CreateProductCreateOneTimePricesSeatBased(components.ProductPriceSeatBasedCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productCreateOneTimePrices.Type {
	case components.ProductCreateOneTimePricesTypeCustom:
		// productCreateOneTimePrices.ProductPriceCustomCreate is populated
	case components.ProductCreateOneTimePricesTypeFixed:
		// productCreateOneTimePrices.ProductPriceFixedCreate is populated
	case components.ProductCreateOneTimePricesTypeFree:
		// productCreateOneTimePrices.ProductPriceFreeCreate is populated
	case components.ProductCreateOneTimePricesTypeMeteredUnit:
		// productCreateOneTimePrices.ProductPriceMeteredUnitCreate is populated
	case components.ProductCreateOneTimePricesTypeSeatBased:
		// productCreateOneTimePrices.ProductPriceSeatBasedCreate is populated
}
```

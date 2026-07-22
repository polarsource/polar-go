# ProductCreateRecurringPrices


## Supported Types

### ProductPriceCustomCreate

```go
productCreateRecurringPrices := components.CreateProductCreateRecurringPricesCustom(components.ProductPriceCustomCreate{/* values here */})
```

### ProductPriceFixedCreate

```go
productCreateRecurringPrices := components.CreateProductCreateRecurringPricesFixed(components.ProductPriceFixedCreate{/* values here */})
```

### ProductPriceMeteredUnitCreate

```go
productCreateRecurringPrices := components.CreateProductCreateRecurringPricesMeteredUnit(components.ProductPriceMeteredUnitCreate{/* values here */})
```

### ProductPriceSeatBasedCreate

```go
productCreateRecurringPrices := components.CreateProductCreateRecurringPricesSeatBased(components.ProductPriceSeatBasedCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productCreateRecurringPrices.Type {
	case components.ProductCreateRecurringPricesTypeCustom:
		// productCreateRecurringPrices.ProductPriceCustomCreate is populated
	case components.ProductCreateRecurringPricesTypeFixed:
		// productCreateRecurringPrices.ProductPriceFixedCreate is populated
	case components.ProductCreateRecurringPricesTypeMeteredUnit:
		// productCreateRecurringPrices.ProductPriceMeteredUnitCreate is populated
	case components.ProductCreateRecurringPricesTypeSeatBased:
		// productCreateRecurringPrices.ProductPriceSeatBasedCreate is populated
}
```

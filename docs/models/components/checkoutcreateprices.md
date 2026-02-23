# CheckoutCreatePrices


## Supported Types

### ProductPriceCustomCreate

```go
checkoutCreatePrices := components.CreateCheckoutCreatePricesCustom(components.ProductPriceCustomCreate{/* values here */})
```

### ProductPriceFixedCreate

```go
checkoutCreatePrices := components.CreateCheckoutCreatePricesFixed(components.ProductPriceFixedCreate{/* values here */})
```

### ProductPriceFreeCreate

```go
checkoutCreatePrices := components.CreateCheckoutCreatePricesFree(components.ProductPriceFreeCreate{/* values here */})
```

### ProductPriceMeteredUnitCreate

```go
checkoutCreatePrices := components.CreateCheckoutCreatePricesMeteredUnit(components.ProductPriceMeteredUnitCreate{/* values here */})
```

### ProductPriceSeatBasedCreate

```go
checkoutCreatePrices := components.CreateCheckoutCreatePricesSeatBased(components.ProductPriceSeatBasedCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutCreatePrices.Type {
	case components.CheckoutCreatePricesTypeCustom:
		// checkoutCreatePrices.ProductPriceCustomCreate is populated
	case components.CheckoutCreatePricesTypeFixed:
		// checkoutCreatePrices.ProductPriceFixedCreate is populated
	case components.CheckoutCreatePricesTypeFree:
		// checkoutCreatePrices.ProductPriceFreeCreate is populated
	case components.CheckoutCreatePricesTypeMeteredUnit:
		// checkoutCreatePrices.ProductPriceMeteredUnitCreate is populated
	case components.CheckoutCreatePricesTypeSeatBased:
		// checkoutCreatePrices.ProductPriceSeatBasedCreate is populated
}
```

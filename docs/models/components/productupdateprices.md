# ProductUpdatePrices


## Supported Types

### ExistingProductPrice

```go
productUpdatePrices := components.CreateProductUpdatePricesExistingProductPrice(components.ExistingProductPrice{/* values here */})
```

### Two

```go
productUpdatePrices := components.CreateProductUpdatePricesTwo(components.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productUpdatePrices.Type {
	case components.ProductUpdatePricesTypeExistingProductPrice:
		// productUpdatePrices.ExistingProductPrice is populated
	case components.ProductUpdatePricesTypeTwo:
		// productUpdatePrices.Two is populated
}
```

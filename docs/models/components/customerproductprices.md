# CustomerProductPrices


## Supported Types

### LegacyRecurringProductPrice

```go
customerProductPrices := components.CreateCustomerProductPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
customerProductPrices := components.CreateCustomerProductPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerProductPrices.Type {
	case components.CustomerProductPricesTypeLegacyRecurringProductPrice:
		// customerProductPrices.LegacyRecurringProductPrice is populated
	case components.CustomerProductPricesTypeProductPrice:
		// customerProductPrices.ProductPrice is populated
}
```

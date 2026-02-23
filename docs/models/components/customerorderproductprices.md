# CustomerOrderProductPrices


## Supported Types

### LegacyRecurringProductPrice

```go
customerOrderProductPrices := components.CreateCustomerOrderProductPricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
customerOrderProductPrices := components.CreateCustomerOrderProductPricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerOrderProductPrices.Type {
	case components.CustomerOrderProductPricesTypeLegacyRecurringProductPrice:
		// customerOrderProductPrices.LegacyRecurringProductPrice is populated
	case components.CustomerOrderProductPricesTypeProductPrice:
		// customerOrderProductPrices.ProductPrice is populated
}
```

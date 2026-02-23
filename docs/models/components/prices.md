# Prices


## Supported Types

### LegacyRecurringProductPrice

```go
prices := components.CreatePricesLegacyRecurringProductPrice(components.LegacyRecurringProductPrice{/* values here */})
```

### ProductPrice

```go
prices := components.CreatePricesProductPrice(components.ProductPrice{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch prices.Type {
	case components.PricesTypeLegacyRecurringProductPrice:
		// prices.LegacyRecurringProductPrice is populated
	case components.PricesTypeProductPrice:
		// prices.ProductPrice is populated
}
```

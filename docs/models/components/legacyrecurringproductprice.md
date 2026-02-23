# LegacyRecurringProductPrice


## Supported Types

### LegacyRecurringProductPriceCustom

```go
legacyRecurringProductPrice := components.CreateLegacyRecurringProductPriceCustom(components.LegacyRecurringProductPriceCustom{/* values here */})
```

### LegacyRecurringProductPriceFixed

```go
legacyRecurringProductPrice := components.CreateLegacyRecurringProductPriceFixed(components.LegacyRecurringProductPriceFixed{/* values here */})
```

### LegacyRecurringProductPriceFree

```go
legacyRecurringProductPrice := components.CreateLegacyRecurringProductPriceFree(components.LegacyRecurringProductPriceFree{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch legacyRecurringProductPrice.Type {
	case components.LegacyRecurringProductPriceTypeCustom:
		// legacyRecurringProductPrice.LegacyRecurringProductPriceCustom is populated
	case components.LegacyRecurringProductPriceTypeFixed:
		// legacyRecurringProductPrice.LegacyRecurringProductPriceFixed is populated
	case components.LegacyRecurringProductPriceTypeFree:
		// legacyRecurringProductPrice.LegacyRecurringProductPriceFree is populated
}
```

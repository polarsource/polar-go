# CustomerStateTaxID


## Supported Types

### 

```go
customerStateTaxID := components.CreateCustomerStateTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerStateTaxID := components.CreateCustomerStateTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerStateTaxID.Type {
	case components.CustomerStateTaxIDTypeStr:
		// customerStateTaxID.Str is populated
	case components.CustomerStateTaxIDTypeTaxIDFormat:
		// customerStateTaxID.TaxIDFormat is populated
}
```

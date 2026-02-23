# CustomerUpdateTaxID


## Supported Types

### 

```go
customerUpdateTaxID := components.CreateCustomerUpdateTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerUpdateTaxID := components.CreateCustomerUpdateTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerUpdateTaxID.Type {
	case components.CustomerUpdateTaxIDTypeStr:
		// customerUpdateTaxID.Str is populated
	case components.CustomerUpdateTaxIDTypeTaxIDFormat:
		// customerUpdateTaxID.TaxIDFormat is populated
}
```

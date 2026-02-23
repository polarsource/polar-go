# CustomerUpdateExternalIDTaxID


## Supported Types

### 

```go
customerUpdateExternalIDTaxID := components.CreateCustomerUpdateExternalIDTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerUpdateExternalIDTaxID := components.CreateCustomerUpdateExternalIDTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerUpdateExternalIDTaxID.Type {
	case components.CustomerUpdateExternalIDTaxIDTypeStr:
		// customerUpdateExternalIDTaxID.Str is populated
	case components.CustomerUpdateExternalIDTaxIDTypeTaxIDFormat:
		// customerUpdateExternalIDTaxID.TaxIDFormat is populated
}
```

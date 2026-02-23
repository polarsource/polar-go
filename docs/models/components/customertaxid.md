# CustomerTaxID


## Supported Types

### 

```go
customerTaxID := components.CreateCustomerTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerTaxID := components.CreateCustomerTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerTaxID.Type {
	case components.CustomerTaxIDTypeStr:
		// customerTaxID.Str is populated
	case components.CustomerTaxIDTypeTaxIDFormat:
		// customerTaxID.TaxIDFormat is populated
}
```

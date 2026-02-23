# CustomerCreateTaxID


## Supported Types

### 

```go
customerCreateTaxID := components.CreateCustomerCreateTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerCreateTaxID := components.CreateCustomerCreateTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerCreateTaxID.Type {
	case components.CustomerCreateTaxIDTypeStr:
		// customerCreateTaxID.Str is populated
	case components.CustomerCreateTaxIDTypeTaxIDFormat:
		// customerCreateTaxID.TaxIDFormat is populated
}
```

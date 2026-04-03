# CustomerIndividualTaxID


## Supported Types

### 

```go
customerIndividualTaxID := components.CreateCustomerIndividualTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerIndividualTaxID := components.CreateCustomerIndividualTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerIndividualTaxID.Type {
	case components.CustomerIndividualTaxIDTypeStr:
		// customerIndividualTaxID.Str is populated
	case components.CustomerIndividualTaxIDTypeTaxIDFormat:
		// customerIndividualTaxID.TaxIDFormat is populated
}
```

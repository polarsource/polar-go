# CustomerStateIndividualTaxID


## Supported Types

### 

```go
customerStateIndividualTaxID := components.CreateCustomerStateIndividualTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerStateIndividualTaxID := components.CreateCustomerStateIndividualTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerStateIndividualTaxID.Type {
	case components.CustomerStateIndividualTaxIDTypeStr:
		// customerStateIndividualTaxID.Str is populated
	case components.CustomerStateIndividualTaxIDTypeTaxIDFormat:
		// customerStateIndividualTaxID.TaxIDFormat is populated
}
```

# CustomerTeamTaxID


## Supported Types

### 

```go
customerTeamTaxID := components.CreateCustomerTeamTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerTeamTaxID := components.CreateCustomerTeamTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerTeamTaxID.Type {
	case components.CustomerTeamTaxIDTypeStr:
		// customerTeamTaxID.Str is populated
	case components.CustomerTeamTaxIDTypeTaxIDFormat:
		// customerTeamTaxID.TaxIDFormat is populated
}
```

# CustomerStateTeamTaxID


## Supported Types

### 

```go
customerStateTeamTaxID := components.CreateCustomerStateTeamTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerStateTeamTaxID := components.CreateCustomerStateTeamTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerStateTeamTaxID.Type {
	case components.CustomerStateTeamTaxIDTypeStr:
		// customerStateTeamTaxID.Str is populated
	case components.CustomerStateTeamTaxIDTypeTaxIDFormat:
		// customerStateTeamTaxID.TaxIDFormat is populated
}
```

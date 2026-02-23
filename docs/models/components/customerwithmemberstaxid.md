# CustomerWithMembersTaxID


## Supported Types

### 

```go
customerWithMembersTaxID := components.CreateCustomerWithMembersTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerWithMembersTaxID := components.CreateCustomerWithMembersTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerWithMembersTaxID.Type {
	case components.CustomerWithMembersTaxIDTypeStr:
		// customerWithMembersTaxID.Str is populated
	case components.CustomerWithMembersTaxIDTypeTaxIDFormat:
		// customerWithMembersTaxID.TaxIDFormat is populated
}
```

# CustomerPortalCustomerTaxID


## Supported Types

### 

```go
customerPortalCustomerTaxID := components.CreateCustomerPortalCustomerTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
customerPortalCustomerTaxID := components.CreateCustomerPortalCustomerTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalCustomerTaxID.Type {
	case components.CustomerPortalCustomerTaxIDTypeStr:
		// customerPortalCustomerTaxID.Str is populated
	case components.CustomerPortalCustomerTaxIDTypeTaxIDFormat:
		// customerPortalCustomerTaxID.TaxIDFormat is populated
}
```

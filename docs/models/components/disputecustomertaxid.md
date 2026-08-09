# DisputeCustomerTaxID


## Supported Types

### 

```go
disputeCustomerTaxID := components.CreateDisputeCustomerTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
disputeCustomerTaxID := components.CreateDisputeCustomerTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch disputeCustomerTaxID.Type {
	case components.DisputeCustomerTaxIDTypeStr:
		// disputeCustomerTaxID.Str is populated
	case components.DisputeCustomerTaxIDTypeTaxIDFormat:
		// disputeCustomerTaxID.TaxIDFormat is populated
}
```

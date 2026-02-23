# OrderCustomerTaxID


## Supported Types

### 

```go
orderCustomerTaxID := components.CreateOrderCustomerTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
orderCustomerTaxID := components.CreateOrderCustomerTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch orderCustomerTaxID.Type {
	case components.OrderCustomerTaxIDTypeStr:
		// orderCustomerTaxID.Str is populated
	case components.OrderCustomerTaxIDTypeTaxIDFormat:
		// orderCustomerTaxID.TaxIDFormat is populated
}
```

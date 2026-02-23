# ProductBillingTypeFilter

Filter by product billing type. `recurring` will filter data corresponding to subscriptions creations or renewals. `one_time` will filter data corresponding to one-time purchases.


## Supported Types

### ProductBillingType

```go
productBillingTypeFilter := operations.CreateProductBillingTypeFilterProductBillingType(components.ProductBillingType{/* values here */})
```

### 

```go
productBillingTypeFilter := operations.CreateProductBillingTypeFilterArrayOfProductBillingType([]components.ProductBillingType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productBillingTypeFilter.Type {
	case operations.ProductBillingTypeFilterTypeProductBillingType:
		// productBillingTypeFilter.ProductBillingType is populated
	case operations.ProductBillingTypeFilterTypeArrayOfProductBillingType:
		// productBillingTypeFilter.ArrayOfProductBillingType is populated
}
```

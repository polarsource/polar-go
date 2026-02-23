# QueryParamProductBillingTypeFilter

Filter by billing type. `recurring` will filter data corresponding to subscriptions creations or renewals. `one_time` will filter data corresponding to one-time purchases.


## Supported Types

### ProductBillingType

```go
queryParamProductBillingTypeFilter := operations.CreateQueryParamProductBillingTypeFilterProductBillingType(components.ProductBillingType{/* values here */})
```

### 

```go
queryParamProductBillingTypeFilter := operations.CreateQueryParamProductBillingTypeFilterArrayOfProductBillingType([]components.ProductBillingType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamProductBillingTypeFilter.Type {
	case operations.QueryParamProductBillingTypeFilterTypeProductBillingType:
		// queryParamProductBillingTypeFilter.ProductBillingType is populated
	case operations.QueryParamProductBillingTypeFilterTypeArrayOfProductBillingType:
		// queryParamProductBillingTypeFilter.ArrayOfProductBillingType is populated
}
```

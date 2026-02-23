# CustomerPortalOrdersListQueryParamProductBillingTypeFilter

Filter by product billing type. `recurring` will filter data corresponding to subscriptions creations or renewals. `one_time` will filter data corresponding to one-time purchases.


## Supported Types

### ProductBillingType

```go
customerPortalOrdersListQueryParamProductBillingTypeFilter := operations.CreateCustomerPortalOrdersListQueryParamProductBillingTypeFilterProductBillingType(components.ProductBillingType{/* values here */})
```

### 

```go
customerPortalOrdersListQueryParamProductBillingTypeFilter := operations.CreateCustomerPortalOrdersListQueryParamProductBillingTypeFilterArrayOfProductBillingType([]components.ProductBillingType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalOrdersListQueryParamProductBillingTypeFilter.Type {
	case operations.CustomerPortalOrdersListQueryParamProductBillingTypeFilterTypeProductBillingType:
		// customerPortalOrdersListQueryParamProductBillingTypeFilter.ProductBillingType is populated
	case operations.CustomerPortalOrdersListQueryParamProductBillingTypeFilterTypeArrayOfProductBillingType:
		// customerPortalOrdersListQueryParamProductBillingTypeFilter.ArrayOfProductBillingType is populated
}
```

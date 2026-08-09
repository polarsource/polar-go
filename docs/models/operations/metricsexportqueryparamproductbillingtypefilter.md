# MetricsExportQueryParamProductBillingTypeFilter

Filter by billing type. `recurring` will filter data corresponding to subscriptions creations or renewals. `one_time` will filter data corresponding to one-time purchases.


## Supported Types

### ProductBillingType

```go
metricsExportQueryParamProductBillingTypeFilter := operations.CreateMetricsExportQueryParamProductBillingTypeFilterProductBillingType(components.ProductBillingType{/* values here */})
```

### 

```go
metricsExportQueryParamProductBillingTypeFilter := operations.CreateMetricsExportQueryParamProductBillingTypeFilterArrayOfProductBillingType([]components.ProductBillingType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsExportQueryParamProductBillingTypeFilter.Type {
	case operations.MetricsExportQueryParamProductBillingTypeFilterTypeProductBillingType:
		// metricsExportQueryParamProductBillingTypeFilter.ProductBillingType is populated
	case operations.MetricsExportQueryParamProductBillingTypeFilterTypeArrayOfProductBillingType:
		// metricsExportQueryParamProductBillingTypeFilter.ArrayOfProductBillingType is populated
}
```

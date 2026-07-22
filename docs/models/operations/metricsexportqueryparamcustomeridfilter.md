# MetricsExportQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
metricsExportQueryParamCustomerIDFilter := operations.CreateMetricsExportQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
metricsExportQueryParamCustomerIDFilter := operations.CreateMetricsExportQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsExportQueryParamCustomerIDFilter.Type {
	case operations.MetricsExportQueryParamCustomerIDFilterTypeStr:
		// metricsExportQueryParamCustomerIDFilter.Str is populated
	case operations.MetricsExportQueryParamCustomerIDFilterTypeArrayOfStr:
		// metricsExportQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

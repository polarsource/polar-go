# MetricsExportQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
metricsExportQueryParamOrganizationIDFilter := operations.CreateMetricsExportQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
metricsExportQueryParamOrganizationIDFilter := operations.CreateMetricsExportQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsExportQueryParamOrganizationIDFilter.Type {
	case operations.MetricsExportQueryParamOrganizationIDFilterTypeStr:
		// metricsExportQueryParamOrganizationIDFilter.Str is populated
	case operations.MetricsExportQueryParamOrganizationIDFilterTypeArrayOfStr:
		// metricsExportQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

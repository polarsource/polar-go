# MetricsListDashboardsQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
metricsListDashboardsQueryParamOrganizationIDFilter := operations.CreateMetricsListDashboardsQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
metricsListDashboardsQueryParamOrganizationIDFilter := operations.CreateMetricsListDashboardsQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsListDashboardsQueryParamOrganizationIDFilter.Type {
	case operations.MetricsListDashboardsQueryParamOrganizationIDFilterTypeStr:
		// metricsListDashboardsQueryParamOrganizationIDFilter.Str is populated
	case operations.MetricsListDashboardsQueryParamOrganizationIDFilterTypeArrayOfStr:
		// metricsListDashboardsQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

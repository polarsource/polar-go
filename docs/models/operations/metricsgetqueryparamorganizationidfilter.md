# MetricsGetQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
metricsGetQueryParamOrganizationIDFilter := operations.CreateMetricsGetQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
metricsGetQueryParamOrganizationIDFilter := operations.CreateMetricsGetQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsGetQueryParamOrganizationIDFilter.Type {
	case operations.MetricsGetQueryParamOrganizationIDFilterTypeStr:
		// metricsGetQueryParamOrganizationIDFilter.Str is populated
	case operations.MetricsGetQueryParamOrganizationIDFilterTypeArrayOfStr:
		// metricsGetQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

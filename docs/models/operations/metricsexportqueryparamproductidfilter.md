# MetricsExportQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
metricsExportQueryParamProductIDFilter := operations.CreateMetricsExportQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
metricsExportQueryParamProductIDFilter := operations.CreateMetricsExportQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsExportQueryParamProductIDFilter.Type {
	case operations.MetricsExportQueryParamProductIDFilterTypeStr:
		// metricsExportQueryParamProductIDFilter.Str is populated
	case operations.MetricsExportQueryParamProductIDFilterTypeArrayOfStr:
		// metricsExportQueryParamProductIDFilter.ArrayOfStr is populated
}
```

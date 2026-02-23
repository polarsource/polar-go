# MetricsGetQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
metricsGetQueryParamProductIDFilter := operations.CreateMetricsGetQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
metricsGetQueryParamProductIDFilter := operations.CreateMetricsGetQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsGetQueryParamProductIDFilter.Type {
	case operations.MetricsGetQueryParamProductIDFilterTypeStr:
		// metricsGetQueryParamProductIDFilter.Str is populated
	case operations.MetricsGetQueryParamProductIDFilterTypeArrayOfStr:
		// metricsGetQueryParamProductIDFilter.ArrayOfStr is populated
}
```

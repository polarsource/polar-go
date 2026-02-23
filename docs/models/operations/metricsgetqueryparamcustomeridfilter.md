# MetricsGetQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
metricsGetQueryParamCustomerIDFilter := operations.CreateMetricsGetQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
metricsGetQueryParamCustomerIDFilter := operations.CreateMetricsGetQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsGetQueryParamCustomerIDFilter.Type {
	case operations.MetricsGetQueryParamCustomerIDFilterTypeStr:
		// metricsGetQueryParamCustomerIDFilter.Str is populated
	case operations.MetricsGetQueryParamCustomerIDFilterTypeArrayOfStr:
		// metricsGetQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

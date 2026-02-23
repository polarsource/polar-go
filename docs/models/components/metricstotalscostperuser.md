# MetricsTotalsCostPerUser


## Supported Types

### 

```go
metricsTotalsCostPerUser := components.CreateMetricsTotalsCostPerUserInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCostPerUser := components.CreateMetricsTotalsCostPerUserNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCostPerUser.Type {
	case components.MetricsTotalsCostPerUserTypeInteger:
		// metricsTotalsCostPerUser.Integer is populated
	case components.MetricsTotalsCostPerUserTypeNumber:
		// metricsTotalsCostPerUser.Number is populated
}
```

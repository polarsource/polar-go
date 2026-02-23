# MetricsTotalsCashflow


## Supported Types

### 

```go
metricsTotalsCashflow := components.CreateMetricsTotalsCashflowInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCashflow := components.CreateMetricsTotalsCashflowNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCashflow.Type {
	case components.MetricsTotalsCashflowTypeInteger:
		// metricsTotalsCashflow.Integer is populated
	case components.MetricsTotalsCashflowTypeNumber:
		// metricsTotalsCashflow.Number is populated
}
```

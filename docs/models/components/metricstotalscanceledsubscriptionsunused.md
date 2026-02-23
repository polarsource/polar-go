# MetricsTotalsCanceledSubscriptionsUnused


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptionsUnused := components.CreateMetricsTotalsCanceledSubscriptionsUnusedInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptionsUnused := components.CreateMetricsTotalsCanceledSubscriptionsUnusedNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptionsUnused.Type {
	case components.MetricsTotalsCanceledSubscriptionsUnusedTypeInteger:
		// metricsTotalsCanceledSubscriptionsUnused.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsUnusedTypeNumber:
		// metricsTotalsCanceledSubscriptionsUnused.Number is populated
}
```

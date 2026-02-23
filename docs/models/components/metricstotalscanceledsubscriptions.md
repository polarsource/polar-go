# MetricsTotalsCanceledSubscriptions


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptions := components.CreateMetricsTotalsCanceledSubscriptionsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptions := components.CreateMetricsTotalsCanceledSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptions.Type {
	case components.MetricsTotalsCanceledSubscriptionsTypeInteger:
		// metricsTotalsCanceledSubscriptions.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsTypeNumber:
		// metricsTotalsCanceledSubscriptions.Number is populated
}
```

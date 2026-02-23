# MetricsTotalsCanceledSubscriptionsTooExpensive


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptionsTooExpensive := components.CreateMetricsTotalsCanceledSubscriptionsTooExpensiveInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptionsTooExpensive := components.CreateMetricsTotalsCanceledSubscriptionsTooExpensiveNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptionsTooExpensive.Type {
	case components.MetricsTotalsCanceledSubscriptionsTooExpensiveTypeInteger:
		// metricsTotalsCanceledSubscriptionsTooExpensive.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsTooExpensiveTypeNumber:
		// metricsTotalsCanceledSubscriptionsTooExpensive.Number is populated
}
```

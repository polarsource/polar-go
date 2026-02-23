# MetricsTotalsCanceledSubscriptionsOther


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptionsOther := components.CreateMetricsTotalsCanceledSubscriptionsOtherInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptionsOther := components.CreateMetricsTotalsCanceledSubscriptionsOtherNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptionsOther.Type {
	case components.MetricsTotalsCanceledSubscriptionsOtherTypeInteger:
		// metricsTotalsCanceledSubscriptionsOther.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsOtherTypeNumber:
		// metricsTotalsCanceledSubscriptionsOther.Number is populated
}
```

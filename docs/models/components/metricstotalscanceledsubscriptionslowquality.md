# MetricsTotalsCanceledSubscriptionsLowQuality


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptionsLowQuality := components.CreateMetricsTotalsCanceledSubscriptionsLowQualityInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptionsLowQuality := components.CreateMetricsTotalsCanceledSubscriptionsLowQualityNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptionsLowQuality.Type {
	case components.MetricsTotalsCanceledSubscriptionsLowQualityTypeInteger:
		// metricsTotalsCanceledSubscriptionsLowQuality.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsLowQualityTypeNumber:
		// metricsTotalsCanceledSubscriptionsLowQuality.Number is populated
}
```

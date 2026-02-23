# MetricsTotalsCanceledSubscriptionsTooComplex


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptionsTooComplex := components.CreateMetricsTotalsCanceledSubscriptionsTooComplexInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptionsTooComplex := components.CreateMetricsTotalsCanceledSubscriptionsTooComplexNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptionsTooComplex.Type {
	case components.MetricsTotalsCanceledSubscriptionsTooComplexTypeInteger:
		// metricsTotalsCanceledSubscriptionsTooComplex.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsTooComplexTypeNumber:
		// metricsTotalsCanceledSubscriptionsTooComplex.Number is populated
}
```

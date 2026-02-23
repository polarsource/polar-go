# MetricsTotalsActiveSubscriptions


## Supported Types

### 

```go
metricsTotalsActiveSubscriptions := components.CreateMetricsTotalsActiveSubscriptionsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsActiveSubscriptions := components.CreateMetricsTotalsActiveSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsActiveSubscriptions.Type {
	case components.MetricsTotalsActiveSubscriptionsTypeInteger:
		// metricsTotalsActiveSubscriptions.Integer is populated
	case components.MetricsTotalsActiveSubscriptionsTypeNumber:
		// metricsTotalsActiveSubscriptions.Number is populated
}
```

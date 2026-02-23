# MetricsTotalsNewSubscriptions


## Supported Types

### 

```go
metricsTotalsNewSubscriptions := components.CreateMetricsTotalsNewSubscriptionsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsNewSubscriptions := components.CreateMetricsTotalsNewSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsNewSubscriptions.Type {
	case components.MetricsTotalsNewSubscriptionsTypeInteger:
		// metricsTotalsNewSubscriptions.Integer is populated
	case components.MetricsTotalsNewSubscriptionsTypeNumber:
		// metricsTotalsNewSubscriptions.Number is populated
}
```

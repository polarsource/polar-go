# MetricsTotalsChurnedSubscriptions


## Supported Types

### 

```go
metricsTotalsChurnedSubscriptions := components.CreateMetricsTotalsChurnedSubscriptionsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsChurnedSubscriptions := components.CreateMetricsTotalsChurnedSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsChurnedSubscriptions.Type {
	case components.MetricsTotalsChurnedSubscriptionsTypeInteger:
		// metricsTotalsChurnedSubscriptions.Integer is populated
	case components.MetricsTotalsChurnedSubscriptionsTypeNumber:
		// metricsTotalsChurnedSubscriptions.Number is populated
}
```

# MetricsTotalsCommittedSubscriptions


## Supported Types

### 

```go
metricsTotalsCommittedSubscriptions := components.CreateMetricsTotalsCommittedSubscriptionsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCommittedSubscriptions := components.CreateMetricsTotalsCommittedSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCommittedSubscriptions.Type {
	case components.MetricsTotalsCommittedSubscriptionsTypeInteger:
		// metricsTotalsCommittedSubscriptions.Integer is populated
	case components.MetricsTotalsCommittedSubscriptionsTypeNumber:
		// metricsTotalsCommittedSubscriptions.Number is populated
}
```

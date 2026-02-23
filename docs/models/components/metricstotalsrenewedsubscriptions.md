# MetricsTotalsRenewedSubscriptions


## Supported Types

### 

```go
metricsTotalsRenewedSubscriptions := components.CreateMetricsTotalsRenewedSubscriptionsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsRenewedSubscriptions := components.CreateMetricsTotalsRenewedSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsRenewedSubscriptions.Type {
	case components.MetricsTotalsRenewedSubscriptionsTypeInteger:
		// metricsTotalsRenewedSubscriptions.Integer is populated
	case components.MetricsTotalsRenewedSubscriptionsTypeNumber:
		// metricsTotalsRenewedSubscriptions.Number is populated
}
```

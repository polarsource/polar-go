# MetricsTotalsRevenue


## Supported Types

### 

```go
metricsTotalsRevenue := components.CreateMetricsTotalsRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsRevenue := components.CreateMetricsTotalsRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsRevenue.Type {
	case components.MetricsTotalsRevenueTypeInteger:
		// metricsTotalsRevenue.Integer is populated
	case components.MetricsTotalsRevenueTypeNumber:
		// metricsTotalsRevenue.Number is populated
}
```

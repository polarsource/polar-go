# MetricsTotalsNetRevenue


## Supported Types

### 

```go
metricsTotalsNetRevenue := components.CreateMetricsTotalsNetRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsNetRevenue := components.CreateMetricsTotalsNetRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsNetRevenue.Type {
	case components.MetricsTotalsNetRevenueTypeInteger:
		// metricsTotalsNetRevenue.Integer is populated
	case components.MetricsTotalsNetRevenueTypeNumber:
		// metricsTotalsNetRevenue.Number is populated
}
```

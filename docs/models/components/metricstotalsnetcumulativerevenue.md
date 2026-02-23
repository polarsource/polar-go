# MetricsTotalsNetCumulativeRevenue


## Supported Types

### 

```go
metricsTotalsNetCumulativeRevenue := components.CreateMetricsTotalsNetCumulativeRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsNetCumulativeRevenue := components.CreateMetricsTotalsNetCumulativeRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsNetCumulativeRevenue.Type {
	case components.MetricsTotalsNetCumulativeRevenueTypeInteger:
		// metricsTotalsNetCumulativeRevenue.Integer is populated
	case components.MetricsTotalsNetCumulativeRevenueTypeNumber:
		// metricsTotalsNetCumulativeRevenue.Number is populated
}
```

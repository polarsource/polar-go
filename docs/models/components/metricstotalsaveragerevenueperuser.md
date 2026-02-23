# MetricsTotalsAverageRevenuePerUser


## Supported Types

### 

```go
metricsTotalsAverageRevenuePerUser := components.CreateMetricsTotalsAverageRevenuePerUserInteger(int64{/* values here */})
```

### 

```go
metricsTotalsAverageRevenuePerUser := components.CreateMetricsTotalsAverageRevenuePerUserNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsAverageRevenuePerUser.Type {
	case components.MetricsTotalsAverageRevenuePerUserTypeInteger:
		// metricsTotalsAverageRevenuePerUser.Integer is populated
	case components.MetricsTotalsAverageRevenuePerUserTypeNumber:
		// metricsTotalsAverageRevenuePerUser.Number is populated
}
```

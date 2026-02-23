# MetricsTotalsOneTimeProductsRevenue


## Supported Types

### 

```go
metricsTotalsOneTimeProductsRevenue := components.CreateMetricsTotalsOneTimeProductsRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsOneTimeProductsRevenue := components.CreateMetricsTotalsOneTimeProductsRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsOneTimeProductsRevenue.Type {
	case components.MetricsTotalsOneTimeProductsRevenueTypeInteger:
		// metricsTotalsOneTimeProductsRevenue.Integer is populated
	case components.MetricsTotalsOneTimeProductsRevenueTypeNumber:
		// metricsTotalsOneTimeProductsRevenue.Number is populated
}
```

# MetricsTotalsOneTimeProductsNetRevenue


## Supported Types

### 

```go
metricsTotalsOneTimeProductsNetRevenue := components.CreateMetricsTotalsOneTimeProductsNetRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsOneTimeProductsNetRevenue := components.CreateMetricsTotalsOneTimeProductsNetRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsOneTimeProductsNetRevenue.Type {
	case components.MetricsTotalsOneTimeProductsNetRevenueTypeInteger:
		// metricsTotalsOneTimeProductsNetRevenue.Integer is populated
	case components.MetricsTotalsOneTimeProductsNetRevenueTypeNumber:
		// metricsTotalsOneTimeProductsNetRevenue.Number is populated
}
```

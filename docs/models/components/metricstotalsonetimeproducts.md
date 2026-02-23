# MetricsTotalsOneTimeProducts


## Supported Types

### 

```go
metricsTotalsOneTimeProducts := components.CreateMetricsTotalsOneTimeProductsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsOneTimeProducts := components.CreateMetricsTotalsOneTimeProductsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsOneTimeProducts.Type {
	case components.MetricsTotalsOneTimeProductsTypeInteger:
		// metricsTotalsOneTimeProducts.Integer is populated
	case components.MetricsTotalsOneTimeProductsTypeNumber:
		// metricsTotalsOneTimeProducts.Number is populated
}
```

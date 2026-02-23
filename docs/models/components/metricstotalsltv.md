# MetricsTotalsLtv


## Supported Types

### 

```go
metricsTotalsLtv := components.CreateMetricsTotalsLtvInteger(int64{/* values here */})
```

### 

```go
metricsTotalsLtv := components.CreateMetricsTotalsLtvNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsLtv.Type {
	case components.MetricsTotalsLtvTypeInteger:
		// metricsTotalsLtv.Integer is populated
	case components.MetricsTotalsLtvTypeNumber:
		// metricsTotalsLtv.Number is populated
}
```

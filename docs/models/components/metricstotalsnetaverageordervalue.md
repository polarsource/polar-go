# MetricsTotalsNetAverageOrderValue


## Supported Types

### 

```go
metricsTotalsNetAverageOrderValue := components.CreateMetricsTotalsNetAverageOrderValueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsNetAverageOrderValue := components.CreateMetricsTotalsNetAverageOrderValueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsNetAverageOrderValue.Type {
	case components.MetricsTotalsNetAverageOrderValueTypeInteger:
		// metricsTotalsNetAverageOrderValue.Integer is populated
	case components.MetricsTotalsNetAverageOrderValueTypeNumber:
		// metricsTotalsNetAverageOrderValue.Number is populated
}
```

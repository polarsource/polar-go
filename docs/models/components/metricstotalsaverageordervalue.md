# MetricsTotalsAverageOrderValue


## Supported Types

### 

```go
metricsTotalsAverageOrderValue := components.CreateMetricsTotalsAverageOrderValueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsAverageOrderValue := components.CreateMetricsTotalsAverageOrderValueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsAverageOrderValue.Type {
	case components.MetricsTotalsAverageOrderValueTypeInteger:
		// metricsTotalsAverageOrderValue.Integer is populated
	case components.MetricsTotalsAverageOrderValueTypeNumber:
		// metricsTotalsAverageOrderValue.Number is populated
}
```

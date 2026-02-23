# MetricsTotalsCheckouts


## Supported Types

### 

```go
metricsTotalsCheckouts := components.CreateMetricsTotalsCheckoutsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCheckouts := components.CreateMetricsTotalsCheckoutsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCheckouts.Type {
	case components.MetricsTotalsCheckoutsTypeInteger:
		// metricsTotalsCheckouts.Integer is populated
	case components.MetricsTotalsCheckoutsTypeNumber:
		// metricsTotalsCheckouts.Number is populated
}
```

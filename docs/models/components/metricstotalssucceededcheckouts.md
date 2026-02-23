# MetricsTotalsSucceededCheckouts


## Supported Types

### 

```go
metricsTotalsSucceededCheckouts := components.CreateMetricsTotalsSucceededCheckoutsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsSucceededCheckouts := components.CreateMetricsTotalsSucceededCheckoutsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsSucceededCheckouts.Type {
	case components.MetricsTotalsSucceededCheckoutsTypeInteger:
		// metricsTotalsSucceededCheckouts.Integer is populated
	case components.MetricsTotalsSucceededCheckoutsTypeNumber:
		// metricsTotalsSucceededCheckouts.Number is populated
}
```

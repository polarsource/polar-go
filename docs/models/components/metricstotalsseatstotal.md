# MetricsTotalsSeatsTotal


## Supported Types

### 

```go
metricsTotalsSeatsTotal := components.CreateMetricsTotalsSeatsTotalInteger(int64{/* values here */})
```

### 

```go
metricsTotalsSeatsTotal := components.CreateMetricsTotalsSeatsTotalNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsSeatsTotal.Type {
	case components.MetricsTotalsSeatsTotalTypeInteger:
		// metricsTotalsSeatsTotal.Integer is populated
	case components.MetricsTotalsSeatsTotalTypeNumber:
		// metricsTotalsSeatsTotal.Number is populated
}
```

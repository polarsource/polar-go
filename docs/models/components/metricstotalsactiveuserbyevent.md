# MetricsTotalsActiveUserByEvent


## Supported Types

### 

```go
metricsTotalsActiveUserByEvent := components.CreateMetricsTotalsActiveUserByEventInteger(int64{/* values here */})
```

### 

```go
metricsTotalsActiveUserByEvent := components.CreateMetricsTotalsActiveUserByEventNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsActiveUserByEvent.Type {
	case components.MetricsTotalsActiveUserByEventTypeInteger:
		// metricsTotalsActiveUserByEvent.Integer is populated
	case components.MetricsTotalsActiveUserByEventTypeNumber:
		// metricsTotalsActiveUserByEvent.Number is populated
}
```

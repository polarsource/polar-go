# MetricsTotalsSeatsPending


## Supported Types

### 

```go
metricsTotalsSeatsPending := components.CreateMetricsTotalsSeatsPendingInteger(int64{/* values here */})
```

### 

```go
metricsTotalsSeatsPending := components.CreateMetricsTotalsSeatsPendingNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsSeatsPending.Type {
	case components.MetricsTotalsSeatsPendingTypeInteger:
		// metricsTotalsSeatsPending.Integer is populated
	case components.MetricsTotalsSeatsPendingTypeNumber:
		// metricsTotalsSeatsPending.Number is populated
}
```

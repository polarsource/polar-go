# MetricsTotalsSeatsClaimed


## Supported Types

### 

```go
metricsTotalsSeatsClaimed := components.CreateMetricsTotalsSeatsClaimedInteger(int64{/* values here */})
```

### 

```go
metricsTotalsSeatsClaimed := components.CreateMetricsTotalsSeatsClaimedNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsSeatsClaimed.Type {
	case components.MetricsTotalsSeatsClaimedTypeInteger:
		// metricsTotalsSeatsClaimed.Integer is populated
	case components.MetricsTotalsSeatsClaimedTypeNumber:
		// metricsTotalsSeatsClaimed.Number is populated
}
```

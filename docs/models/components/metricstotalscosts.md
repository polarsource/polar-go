# MetricsTotalsCosts


## Supported Types

### 

```go
metricsTotalsCosts := components.CreateMetricsTotalsCostsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCosts := components.CreateMetricsTotalsCostsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCosts.Type {
	case components.MetricsTotalsCostsTypeInteger:
		// metricsTotalsCosts.Integer is populated
	case components.MetricsTotalsCostsTypeNumber:
		// metricsTotalsCosts.Number is populated
}
```

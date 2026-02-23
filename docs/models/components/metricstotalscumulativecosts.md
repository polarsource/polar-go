# MetricsTotalsCumulativeCosts


## Supported Types

### 

```go
metricsTotalsCumulativeCosts := components.CreateMetricsTotalsCumulativeCostsInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCumulativeCosts := components.CreateMetricsTotalsCumulativeCostsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCumulativeCosts.Type {
	case components.MetricsTotalsCumulativeCostsTypeInteger:
		// metricsTotalsCumulativeCosts.Integer is populated
	case components.MetricsTotalsCumulativeCostsTypeNumber:
		// metricsTotalsCumulativeCosts.Number is populated
}
```

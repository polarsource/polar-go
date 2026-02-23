# MetricsTotalsCumulativeRevenue


## Supported Types

### 

```go
metricsTotalsCumulativeRevenue := components.CreateMetricsTotalsCumulativeRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCumulativeRevenue := components.CreateMetricsTotalsCumulativeRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCumulativeRevenue.Type {
	case components.MetricsTotalsCumulativeRevenueTypeInteger:
		// metricsTotalsCumulativeRevenue.Integer is populated
	case components.MetricsTotalsCumulativeRevenueTypeNumber:
		// metricsTotalsCumulativeRevenue.Number is populated
}
```

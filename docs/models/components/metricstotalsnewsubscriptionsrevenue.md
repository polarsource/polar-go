# MetricsTotalsNewSubscriptionsRevenue


## Supported Types

### 

```go
metricsTotalsNewSubscriptionsRevenue := components.CreateMetricsTotalsNewSubscriptionsRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsNewSubscriptionsRevenue := components.CreateMetricsTotalsNewSubscriptionsRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsNewSubscriptionsRevenue.Type {
	case components.MetricsTotalsNewSubscriptionsRevenueTypeInteger:
		// metricsTotalsNewSubscriptionsRevenue.Integer is populated
	case components.MetricsTotalsNewSubscriptionsRevenueTypeNumber:
		// metricsTotalsNewSubscriptionsRevenue.Number is populated
}
```

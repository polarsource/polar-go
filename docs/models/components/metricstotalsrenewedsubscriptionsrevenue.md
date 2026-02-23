# MetricsTotalsRenewedSubscriptionsRevenue


## Supported Types

### 

```go
metricsTotalsRenewedSubscriptionsRevenue := components.CreateMetricsTotalsRenewedSubscriptionsRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsRenewedSubscriptionsRevenue := components.CreateMetricsTotalsRenewedSubscriptionsRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsRenewedSubscriptionsRevenue.Type {
	case components.MetricsTotalsRenewedSubscriptionsRevenueTypeInteger:
		// metricsTotalsRenewedSubscriptionsRevenue.Integer is populated
	case components.MetricsTotalsRenewedSubscriptionsRevenueTypeNumber:
		// metricsTotalsRenewedSubscriptionsRevenue.Number is populated
}
```

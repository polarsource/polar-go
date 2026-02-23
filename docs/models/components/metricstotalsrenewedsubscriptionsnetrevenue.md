# MetricsTotalsRenewedSubscriptionsNetRevenue


## Supported Types

### 

```go
metricsTotalsRenewedSubscriptionsNetRevenue := components.CreateMetricsTotalsRenewedSubscriptionsNetRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsRenewedSubscriptionsNetRevenue := components.CreateMetricsTotalsRenewedSubscriptionsNetRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsRenewedSubscriptionsNetRevenue.Type {
	case components.MetricsTotalsRenewedSubscriptionsNetRevenueTypeInteger:
		// metricsTotalsRenewedSubscriptionsNetRevenue.Integer is populated
	case components.MetricsTotalsRenewedSubscriptionsNetRevenueTypeNumber:
		// metricsTotalsRenewedSubscriptionsNetRevenue.Number is populated
}
```

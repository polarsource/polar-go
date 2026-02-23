# MetricsTotalsMonthlyRecurringRevenue


## Supported Types

### 

```go
metricsTotalsMonthlyRecurringRevenue := components.CreateMetricsTotalsMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsMonthlyRecurringRevenue := components.CreateMetricsTotalsMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsMonthlyRecurringRevenue.Type {
	case components.MetricsTotalsMonthlyRecurringRevenueTypeInteger:
		// metricsTotalsMonthlyRecurringRevenue.Integer is populated
	case components.MetricsTotalsMonthlyRecurringRevenueTypeNumber:
		// metricsTotalsMonthlyRecurringRevenue.Number is populated
}
```

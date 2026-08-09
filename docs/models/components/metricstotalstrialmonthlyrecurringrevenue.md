# MetricsTotalsTrialMonthlyRecurringRevenue


## Supported Types

### 

```go
metricsTotalsTrialMonthlyRecurringRevenue := components.CreateMetricsTotalsTrialMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsTrialMonthlyRecurringRevenue := components.CreateMetricsTotalsTrialMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsTrialMonthlyRecurringRevenue.Type {
	case components.MetricsTotalsTrialMonthlyRecurringRevenueTypeInteger:
		// metricsTotalsTrialMonthlyRecurringRevenue.Integer is populated
	case components.MetricsTotalsTrialMonthlyRecurringRevenueTypeNumber:
		// metricsTotalsTrialMonthlyRecurringRevenue.Number is populated
}
```

# MetricsTotalsTrialCommittedMonthlyRecurringRevenue


## Supported Types

### 

```go
metricsTotalsTrialCommittedMonthlyRecurringRevenue := components.CreateMetricsTotalsTrialCommittedMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsTrialCommittedMonthlyRecurringRevenue := components.CreateMetricsTotalsTrialCommittedMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsTrialCommittedMonthlyRecurringRevenue.Type {
	case components.MetricsTotalsTrialCommittedMonthlyRecurringRevenueTypeInteger:
		// metricsTotalsTrialCommittedMonthlyRecurringRevenue.Integer is populated
	case components.MetricsTotalsTrialCommittedMonthlyRecurringRevenueTypeNumber:
		// metricsTotalsTrialCommittedMonthlyRecurringRevenue.Number is populated
}
```

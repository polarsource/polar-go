# MetricsTotalsCommittedMonthlyRecurringRevenue


## Supported Types

### 

```go
metricsTotalsCommittedMonthlyRecurringRevenue := components.CreateMetricsTotalsCommittedMonthlyRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCommittedMonthlyRecurringRevenue := components.CreateMetricsTotalsCommittedMonthlyRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCommittedMonthlyRecurringRevenue.Type {
	case components.MetricsTotalsCommittedMonthlyRecurringRevenueTypeInteger:
		// metricsTotalsCommittedMonthlyRecurringRevenue.Integer is populated
	case components.MetricsTotalsCommittedMonthlyRecurringRevenueTypeNumber:
		// metricsTotalsCommittedMonthlyRecurringRevenue.Number is populated
}
```

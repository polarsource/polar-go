# MetricsTotalsCommittedAnnualRecurringRevenue


## Supported Types

### 

```go
metricsTotalsCommittedAnnualRecurringRevenue := components.CreateMetricsTotalsCommittedAnnualRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCommittedAnnualRecurringRevenue := components.CreateMetricsTotalsCommittedAnnualRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCommittedAnnualRecurringRevenue.Type {
	case components.MetricsTotalsCommittedAnnualRecurringRevenueTypeInteger:
		// metricsTotalsCommittedAnnualRecurringRevenue.Integer is populated
	case components.MetricsTotalsCommittedAnnualRecurringRevenueTypeNumber:
		// metricsTotalsCommittedAnnualRecurringRevenue.Number is populated
}
```

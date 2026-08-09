# MetricsTotalsAnnualRecurringRevenue


## Supported Types

### 

```go
metricsTotalsAnnualRecurringRevenue := components.CreateMetricsTotalsAnnualRecurringRevenueInteger(int64{/* values here */})
```

### 

```go
metricsTotalsAnnualRecurringRevenue := components.CreateMetricsTotalsAnnualRecurringRevenueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsAnnualRecurringRevenue.Type {
	case components.MetricsTotalsAnnualRecurringRevenueTypeInteger:
		// metricsTotalsAnnualRecurringRevenue.Integer is populated
	case components.MetricsTotalsAnnualRecurringRevenueTypeNumber:
		// metricsTotalsAnnualRecurringRevenue.Number is populated
}
```

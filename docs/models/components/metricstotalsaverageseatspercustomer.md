# MetricsTotalsAverageSeatsPerCustomer


## Supported Types

### 

```go
metricsTotalsAverageSeatsPerCustomer := components.CreateMetricsTotalsAverageSeatsPerCustomerInteger(int64{/* values here */})
```

### 

```go
metricsTotalsAverageSeatsPerCustomer := components.CreateMetricsTotalsAverageSeatsPerCustomerNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsAverageSeatsPerCustomer.Type {
	case components.MetricsTotalsAverageSeatsPerCustomerTypeInteger:
		// metricsTotalsAverageSeatsPerCustomer.Integer is populated
	case components.MetricsTotalsAverageSeatsPerCustomerTypeNumber:
		// metricsTotalsAverageSeatsPerCustomer.Number is populated
}
```

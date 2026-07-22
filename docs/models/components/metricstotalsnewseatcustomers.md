# MetricsTotalsNewSeatCustomers


## Supported Types

### 

```go
metricsTotalsNewSeatCustomers := components.CreateMetricsTotalsNewSeatCustomersInteger(int64{/* values here */})
```

### 

```go
metricsTotalsNewSeatCustomers := components.CreateMetricsTotalsNewSeatCustomersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsNewSeatCustomers.Type {
	case components.MetricsTotalsNewSeatCustomersTypeInteger:
		// metricsTotalsNewSeatCustomers.Integer is populated
	case components.MetricsTotalsNewSeatCustomersTypeNumber:
		// metricsTotalsNewSeatCustomers.Number is populated
}
```

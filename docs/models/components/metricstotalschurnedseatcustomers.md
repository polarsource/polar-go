# MetricsTotalsChurnedSeatCustomers


## Supported Types

### 

```go
metricsTotalsChurnedSeatCustomers := components.CreateMetricsTotalsChurnedSeatCustomersInteger(int64{/* values here */})
```

### 

```go
metricsTotalsChurnedSeatCustomers := components.CreateMetricsTotalsChurnedSeatCustomersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsChurnedSeatCustomers.Type {
	case components.MetricsTotalsChurnedSeatCustomersTypeInteger:
		// metricsTotalsChurnedSeatCustomers.Integer is populated
	case components.MetricsTotalsChurnedSeatCustomersTypeNumber:
		// metricsTotalsChurnedSeatCustomers.Number is populated
}
```

# MetricsTotalsSeatCustomers


## Supported Types

### 

```go
metricsTotalsSeatCustomers := components.CreateMetricsTotalsSeatCustomersInteger(int64{/* values here */})
```

### 

```go
metricsTotalsSeatCustomers := components.CreateMetricsTotalsSeatCustomersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsSeatCustomers.Type {
	case components.MetricsTotalsSeatCustomersTypeInteger:
		// metricsTotalsSeatCustomers.Integer is populated
	case components.MetricsTotalsSeatCustomersTypeNumber:
		// metricsTotalsSeatCustomers.Number is populated
}
```

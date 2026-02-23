# MetricsTotalsCanceledSubscriptionsCustomerService


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptionsCustomerService := components.CreateMetricsTotalsCanceledSubscriptionsCustomerServiceInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptionsCustomerService := components.CreateMetricsTotalsCanceledSubscriptionsCustomerServiceNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptionsCustomerService.Type {
	case components.MetricsTotalsCanceledSubscriptionsCustomerServiceTypeInteger:
		// metricsTotalsCanceledSubscriptionsCustomerService.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsCustomerServiceTypeNumber:
		// metricsTotalsCanceledSubscriptionsCustomerService.Number is populated
}
```

# MetricsTotalsCanceledSubscriptionsSwitchedService


## Supported Types

### 

```go
metricsTotalsCanceledSubscriptionsSwitchedService := components.CreateMetricsTotalsCanceledSubscriptionsSwitchedServiceInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCanceledSubscriptionsSwitchedService := components.CreateMetricsTotalsCanceledSubscriptionsSwitchedServiceNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCanceledSubscriptionsSwitchedService.Type {
	case components.MetricsTotalsCanceledSubscriptionsSwitchedServiceTypeInteger:
		// metricsTotalsCanceledSubscriptionsSwitchedService.Integer is populated
	case components.MetricsTotalsCanceledSubscriptionsSwitchedServiceTypeNumber:
		// metricsTotalsCanceledSubscriptionsSwitchedService.Number is populated
}
```

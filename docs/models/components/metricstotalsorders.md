# MetricsTotalsOrders


## Supported Types

### 

```go
metricsTotalsOrders := components.CreateMetricsTotalsOrdersInteger(int64{/* values here */})
```

### 

```go
metricsTotalsOrders := components.CreateMetricsTotalsOrdersNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsOrders.Type {
	case components.MetricsTotalsOrdersTypeInteger:
		// metricsTotalsOrders.Integer is populated
	case components.MetricsTotalsOrdersTypeNumber:
		// metricsTotalsOrders.Number is populated
}
```

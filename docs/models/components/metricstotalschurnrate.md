# MetricsTotalsChurnRate


## Supported Types

### 

```go
metricsTotalsChurnRate := components.CreateMetricsTotalsChurnRateInteger(int64{/* values here */})
```

### 

```go
metricsTotalsChurnRate := components.CreateMetricsTotalsChurnRateNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsChurnRate.Type {
	case components.MetricsTotalsChurnRateTypeInteger:
		// metricsTotalsChurnRate.Integer is populated
	case components.MetricsTotalsChurnRateTypeNumber:
		// metricsTotalsChurnRate.Number is populated
}
```

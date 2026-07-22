# MetricsTotalsSeatUtilizationRate


## Supported Types

### 

```go
metricsTotalsSeatUtilizationRate := components.CreateMetricsTotalsSeatUtilizationRateInteger(int64{/* values here */})
```

### 

```go
metricsTotalsSeatUtilizationRate := components.CreateMetricsTotalsSeatUtilizationRateNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsSeatUtilizationRate.Type {
	case components.MetricsTotalsSeatUtilizationRateTypeInteger:
		// metricsTotalsSeatUtilizationRate.Integer is populated
	case components.MetricsTotalsSeatUtilizationRateTypeNumber:
		// metricsTotalsSeatUtilizationRate.Number is populated
}
```

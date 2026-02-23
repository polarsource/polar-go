# MetricsTotalsCheckoutsConversion


## Supported Types

### 

```go
metricsTotalsCheckoutsConversion := components.CreateMetricsTotalsCheckoutsConversionInteger(int64{/* values here */})
```

### 

```go
metricsTotalsCheckoutsConversion := components.CreateMetricsTotalsCheckoutsConversionNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsCheckoutsConversion.Type {
	case components.MetricsTotalsCheckoutsConversionTypeInteger:
		// metricsTotalsCheckoutsConversion.Integer is populated
	case components.MetricsTotalsCheckoutsConversionTypeNumber:
		// metricsTotalsCheckoutsConversion.Number is populated
}
```

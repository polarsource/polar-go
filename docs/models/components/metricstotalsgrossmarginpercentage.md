# MetricsTotalsGrossMarginPercentage


## Supported Types

### 

```go
metricsTotalsGrossMarginPercentage := components.CreateMetricsTotalsGrossMarginPercentageInteger(int64{/* values here */})
```

### 

```go
metricsTotalsGrossMarginPercentage := components.CreateMetricsTotalsGrossMarginPercentageNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsGrossMarginPercentage.Type {
	case components.MetricsTotalsGrossMarginPercentageTypeInteger:
		// metricsTotalsGrossMarginPercentage.Integer is populated
	case components.MetricsTotalsGrossMarginPercentageTypeNumber:
		// metricsTotalsGrossMarginPercentage.Number is populated
}
```

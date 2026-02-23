# MetricsTotalsGrossMargin


## Supported Types

### 

```go
metricsTotalsGrossMargin := components.CreateMetricsTotalsGrossMarginInteger(int64{/* values here */})
```

### 

```go
metricsTotalsGrossMargin := components.CreateMetricsTotalsGrossMarginNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metricsTotalsGrossMargin.Type {
	case components.MetricsTotalsGrossMarginTypeInteger:
		// metricsTotalsGrossMargin.Integer is populated
	case components.MetricsTotalsGrossMarginTypeNumber:
		// metricsTotalsGrossMargin.Number is populated
}
```

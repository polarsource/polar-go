# MeterCreateAggregation

The aggregation to apply on the filtered events to calculate the meter.


## Supported Types

### PropertyAggregation

```go
meterCreateAggregation := components.CreateMeterCreateAggregationAvg(components.PropertyAggregation{/* values here */})
```

### CountAggregation

```go
meterCreateAggregation := components.CreateMeterCreateAggregationCount(components.CountAggregation{/* values here */})
```

### PropertyAggregation

```go
meterCreateAggregation := components.CreateMeterCreateAggregationMax(components.PropertyAggregation{/* values here */})
```

### PropertyAggregation

```go
meterCreateAggregation := components.CreateMeterCreateAggregationMin(components.PropertyAggregation{/* values here */})
```

### PropertyAggregation

```go
meterCreateAggregation := components.CreateMeterCreateAggregationSum(components.PropertyAggregation{/* values here */})
```

### UniqueAggregation

```go
meterCreateAggregation := components.CreateMeterCreateAggregationUnique(components.UniqueAggregation{/* values here */})
```


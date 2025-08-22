# MeterAggregation

The aggregation to apply on the filtered events to calculate the meter.


## Supported Types

### PropertyAggregation

```go
meterAggregation := components.CreateMeterAggregationAvg(components.PropertyAggregation{/* values here */})
```

### CountAggregation

```go
meterAggregation := components.CreateMeterAggregationCount(components.CountAggregation{/* values here */})
```

### PropertyAggregation

```go
meterAggregation := components.CreateMeterAggregationMax(components.PropertyAggregation{/* values here */})
```

### PropertyAggregation

```go
meterAggregation := components.CreateMeterAggregationMin(components.PropertyAggregation{/* values here */})
```

### PropertyAggregation

```go
meterAggregation := components.CreateMeterAggregationSum(components.PropertyAggregation{/* values here */})
```

### UniqueAggregation

```go
meterAggregation := components.CreateMeterAggregationUnique(components.UniqueAggregation{/* values here */})
```


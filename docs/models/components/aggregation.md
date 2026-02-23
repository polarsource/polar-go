# Aggregation


## Supported Types

### PropertyAggregation

```go
aggregation := components.CreateAggregationAvg(components.PropertyAggregation{/* values here */})
```

### CountAggregation

```go
aggregation := components.CreateAggregationCount(components.CountAggregation{/* values here */})
```

### PropertyAggregation

```go
aggregation := components.CreateAggregationMax(components.PropertyAggregation{/* values here */})
```

### PropertyAggregation

```go
aggregation := components.CreateAggregationMin(components.PropertyAggregation{/* values here */})
```

### PropertyAggregation

```go
aggregation := components.CreateAggregationSum(components.PropertyAggregation{/* values here */})
```

### UniqueAggregation

```go
aggregation := components.CreateAggregationUnique(components.UniqueAggregation{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch aggregation.Type {
	case components.AggregationTypeAvg:
		// aggregation.PropertyAggregation is populated
	case components.AggregationTypeCount:
		// aggregation.CountAggregation is populated
	case components.AggregationTypeMax:
		// aggregation.PropertyAggregation is populated
	case components.AggregationTypeMin:
		// aggregation.PropertyAggregation is populated
	case components.AggregationTypeSum:
		// aggregation.PropertyAggregation is populated
	case components.AggregationTypeUnique:
		// aggregation.UniqueAggregation is populated
}
```

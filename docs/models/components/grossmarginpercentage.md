# GrossMarginPercentage


## Supported Types

### 

```go
grossMarginPercentage := components.CreateGrossMarginPercentageInteger(int64{/* values here */})
```

### 

```go
grossMarginPercentage := components.CreateGrossMarginPercentageNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch grossMarginPercentage.Type {
	case components.GrossMarginPercentageTypeInteger:
		// grossMarginPercentage.Integer is populated
	case components.GrossMarginPercentageTypeNumber:
		// grossMarginPercentage.Number is populated
}
```

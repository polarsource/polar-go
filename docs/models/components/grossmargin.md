# GrossMargin


## Supported Types

### 

```go
grossMargin := components.CreateGrossMarginInteger(int64{/* values here */})
```

### 

```go
grossMargin := components.CreateGrossMarginNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch grossMargin.Type {
	case components.GrossMarginTypeInteger:
		// grossMargin.Integer is populated
	case components.GrossMarginTypeNumber:
		// grossMargin.Number is populated
}
```

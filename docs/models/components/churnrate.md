# ChurnRate


## Supported Types

### 

```go
churnRate := components.CreateChurnRateInteger(int64{/* values here */})
```

### 

```go
churnRate := components.CreateChurnRateNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch churnRate.Type {
	case components.ChurnRateTypeInteger:
		// churnRate.Integer is populated
	case components.ChurnRateTypeNumber:
		// churnRate.Number is populated
}
```

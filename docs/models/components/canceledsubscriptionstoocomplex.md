# CanceledSubscriptionsTooComplex


## Supported Types

### 

```go
canceledSubscriptionsTooComplex := components.CreateCanceledSubscriptionsTooComplexInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsTooComplex := components.CreateCanceledSubscriptionsTooComplexNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsTooComplex.Type {
	case components.CanceledSubscriptionsTooComplexTypeInteger:
		// canceledSubscriptionsTooComplex.Integer is populated
	case components.CanceledSubscriptionsTooComplexTypeNumber:
		// canceledSubscriptionsTooComplex.Number is populated
}
```

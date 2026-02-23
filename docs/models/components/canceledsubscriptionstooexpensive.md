# CanceledSubscriptionsTooExpensive


## Supported Types

### 

```go
canceledSubscriptionsTooExpensive := components.CreateCanceledSubscriptionsTooExpensiveInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsTooExpensive := components.CreateCanceledSubscriptionsTooExpensiveNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsTooExpensive.Type {
	case components.CanceledSubscriptionsTooExpensiveTypeInteger:
		// canceledSubscriptionsTooExpensive.Integer is populated
	case components.CanceledSubscriptionsTooExpensiveTypeNumber:
		// canceledSubscriptionsTooExpensive.Number is populated
}
```

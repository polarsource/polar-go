# CanceledSubscriptionsUnused


## Supported Types

### 

```go
canceledSubscriptionsUnused := components.CreateCanceledSubscriptionsUnusedInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsUnused := components.CreateCanceledSubscriptionsUnusedNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsUnused.Type {
	case components.CanceledSubscriptionsUnusedTypeInteger:
		// canceledSubscriptionsUnused.Integer is populated
	case components.CanceledSubscriptionsUnusedTypeNumber:
		// canceledSubscriptionsUnused.Number is populated
}
```

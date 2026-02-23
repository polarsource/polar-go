# CanceledSubscriptions


## Supported Types

### 

```go
canceledSubscriptions := components.CreateCanceledSubscriptionsInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptions := components.CreateCanceledSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptions.Type {
	case components.CanceledSubscriptionsTypeInteger:
		// canceledSubscriptions.Integer is populated
	case components.CanceledSubscriptionsTypeNumber:
		// canceledSubscriptions.Number is populated
}
```

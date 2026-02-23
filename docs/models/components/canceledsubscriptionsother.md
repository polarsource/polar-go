# CanceledSubscriptionsOther


## Supported Types

### 

```go
canceledSubscriptionsOther := components.CreateCanceledSubscriptionsOtherInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsOther := components.CreateCanceledSubscriptionsOtherNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsOther.Type {
	case components.CanceledSubscriptionsOtherTypeInteger:
		// canceledSubscriptionsOther.Integer is populated
	case components.CanceledSubscriptionsOtherTypeNumber:
		// canceledSubscriptionsOther.Number is populated
}
```

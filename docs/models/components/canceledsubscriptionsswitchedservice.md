# CanceledSubscriptionsSwitchedService


## Supported Types

### 

```go
canceledSubscriptionsSwitchedService := components.CreateCanceledSubscriptionsSwitchedServiceInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsSwitchedService := components.CreateCanceledSubscriptionsSwitchedServiceNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsSwitchedService.Type {
	case components.CanceledSubscriptionsSwitchedServiceTypeInteger:
		// canceledSubscriptionsSwitchedService.Integer is populated
	case components.CanceledSubscriptionsSwitchedServiceTypeNumber:
		// canceledSubscriptionsSwitchedService.Number is populated
}
```

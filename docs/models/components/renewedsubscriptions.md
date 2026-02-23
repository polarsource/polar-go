# RenewedSubscriptions


## Supported Types

### 

```go
renewedSubscriptions := components.CreateRenewedSubscriptionsInteger(int64{/* values here */})
```

### 

```go
renewedSubscriptions := components.CreateRenewedSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch renewedSubscriptions.Type {
	case components.RenewedSubscriptionsTypeInteger:
		// renewedSubscriptions.Integer is populated
	case components.RenewedSubscriptionsTypeNumber:
		// renewedSubscriptions.Number is populated
}
```

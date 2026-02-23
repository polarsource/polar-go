# ActiveSubscriptions


## Supported Types

### 

```go
activeSubscriptions := components.CreateActiveSubscriptionsInteger(int64{/* values here */})
```

### 

```go
activeSubscriptions := components.CreateActiveSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch activeSubscriptions.Type {
	case components.ActiveSubscriptionsTypeInteger:
		// activeSubscriptions.Integer is populated
	case components.ActiveSubscriptionsTypeNumber:
		// activeSubscriptions.Number is populated
}
```

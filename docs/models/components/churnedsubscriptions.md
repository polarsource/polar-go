# ChurnedSubscriptions


## Supported Types

### 

```go
churnedSubscriptions := components.CreateChurnedSubscriptionsInteger(int64{/* values here */})
```

### 

```go
churnedSubscriptions := components.CreateChurnedSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch churnedSubscriptions.Type {
	case components.ChurnedSubscriptionsTypeInteger:
		// churnedSubscriptions.Integer is populated
	case components.ChurnedSubscriptionsTypeNumber:
		// churnedSubscriptions.Number is populated
}
```

# NewSubscriptions


## Supported Types

### 

```go
newSubscriptions := components.CreateNewSubscriptionsInteger(int64{/* values here */})
```

### 

```go
newSubscriptions := components.CreateNewSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch newSubscriptions.Type {
	case components.NewSubscriptionsTypeInteger:
		// newSubscriptions.Integer is populated
	case components.NewSubscriptionsTypeNumber:
		// newSubscriptions.Number is populated
}
```

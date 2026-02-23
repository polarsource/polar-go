# CommittedSubscriptions


## Supported Types

### 

```go
committedSubscriptions := components.CreateCommittedSubscriptionsInteger(int64{/* values here */})
```

### 

```go
committedSubscriptions := components.CreateCommittedSubscriptionsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch committedSubscriptions.Type {
	case components.CommittedSubscriptionsTypeInteger:
		// committedSubscriptions.Integer is populated
	case components.CommittedSubscriptionsTypeNumber:
		// committedSubscriptions.Number is populated
}
```

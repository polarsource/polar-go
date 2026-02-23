# SubscriptionIDFilter

Filter by subscription ID.


## Supported Types

### 

```go
subscriptionIDFilter := operations.CreateSubscriptionIDFilterStr(string{/* values here */})
```

### 

```go
subscriptionIDFilter := operations.CreateSubscriptionIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionIDFilter.Type {
	case operations.SubscriptionIDFilterTypeStr:
		// subscriptionIDFilter.Str is populated
	case operations.SubscriptionIDFilterTypeArrayOfStr:
		// subscriptionIDFilter.ArrayOfStr is populated
}
```

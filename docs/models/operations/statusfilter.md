# StatusFilter

Filter by subscription status.


## Supported Types

### SubscriptionStatus

```go
statusFilter := operations.CreateStatusFilterSubscriptionStatus(components.SubscriptionStatus{/* values here */})
```

### 

```go
statusFilter := operations.CreateStatusFilterArrayOfSubscriptionStatus([]components.SubscriptionStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch statusFilter.Type {
	case operations.StatusFilterTypeSubscriptionStatus:
		// statusFilter.SubscriptionStatus is populated
	case operations.StatusFilterTypeArrayOfSubscriptionStatus:
		// statusFilter.ArrayOfSubscriptionStatus is populated
}
```

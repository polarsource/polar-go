# SubscriptionUpdate


## Supported Types

### SubscriptionUpdateBase

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateBase(components.SubscriptionUpdateBase{/* values here */})
```

### SubscriptionUpdateSeats

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateSeats(components.SubscriptionUpdateSeats{/* values here */})
```

### SubscriptionUpdateBillingPeriod

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateBillingPeriod(components.SubscriptionUpdateBillingPeriod{/* values here */})
```

### SubscriptionCancel

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionCancel(components.SubscriptionCancel{/* values here */})
```

### SubscriptionRevoke

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionRevoke(components.SubscriptionRevoke{/* values here */})
```

### SubscriptionPause

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionPause(components.SubscriptionPause{/* values here */})
```

### SubscriptionResume

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionResume(components.SubscriptionResume{/* values here */})
```

### SubscriptionUpdateClear

```go
subscriptionUpdate := components.CreateSubscriptionUpdateSubscriptionUpdateClear(components.SubscriptionUpdateClear{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionUpdate.Type {
	case components.SubscriptionUpdateTypeSubscriptionUpdateBase:
		// subscriptionUpdate.SubscriptionUpdateBase is populated
	case components.SubscriptionUpdateTypeSubscriptionUpdateSeats:
		// subscriptionUpdate.SubscriptionUpdateSeats is populated
	case components.SubscriptionUpdateTypeSubscriptionUpdateBillingPeriod:
		// subscriptionUpdate.SubscriptionUpdateBillingPeriod is populated
	case components.SubscriptionUpdateTypeSubscriptionCancel:
		// subscriptionUpdate.SubscriptionCancel is populated
	case components.SubscriptionUpdateTypeSubscriptionRevoke:
		// subscriptionUpdate.SubscriptionRevoke is populated
	case components.SubscriptionUpdateTypeSubscriptionPause:
		// subscriptionUpdate.SubscriptionPause is populated
	case components.SubscriptionUpdateTypeSubscriptionResume:
		// subscriptionUpdate.SubscriptionResume is populated
	case components.SubscriptionUpdateTypeSubscriptionUpdateClear:
		// subscriptionUpdate.SubscriptionUpdateClear is populated
}
```

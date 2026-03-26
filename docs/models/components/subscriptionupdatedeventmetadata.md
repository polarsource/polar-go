# SubscriptionUpdatedEventMetadata


## Supported Types

### SubscriptionUpdatedProductMetadata

```go
subscriptionUpdatedEventMetadata := components.CreateSubscriptionUpdatedEventMetadataSubscriptionUpdatedProductMetadata(components.SubscriptionUpdatedProductMetadata{/* values here */})
```

### SubscriptionUpdatedDiscountMetadata

```go
subscriptionUpdatedEventMetadata := components.CreateSubscriptionUpdatedEventMetadataSubscriptionUpdatedDiscountMetadata(components.SubscriptionUpdatedDiscountMetadata{/* values here */})
```

### SubscriptionUpdatedTrialMetadata

```go
subscriptionUpdatedEventMetadata := components.CreateSubscriptionUpdatedEventMetadataSubscriptionUpdatedTrialMetadata(components.SubscriptionUpdatedTrialMetadata{/* values here */})
```

### SubscriptionUpdatedSeatsMetadata

```go
subscriptionUpdatedEventMetadata := components.CreateSubscriptionUpdatedEventMetadataSubscriptionUpdatedSeatsMetadata(components.SubscriptionUpdatedSeatsMetadata{/* values here */})
```

### SubscriptionUpdatedBillingPeriodMetadata

```go
subscriptionUpdatedEventMetadata := components.CreateSubscriptionUpdatedEventMetadataSubscriptionUpdatedBillingPeriodMetadata(components.SubscriptionUpdatedBillingPeriodMetadata{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionUpdatedEventMetadata.Type {
	case components.SubscriptionUpdatedEventMetadataTypeSubscriptionUpdatedProductMetadata:
		// subscriptionUpdatedEventMetadata.SubscriptionUpdatedProductMetadata is populated
	case components.SubscriptionUpdatedEventMetadataTypeSubscriptionUpdatedDiscountMetadata:
		// subscriptionUpdatedEventMetadata.SubscriptionUpdatedDiscountMetadata is populated
	case components.SubscriptionUpdatedEventMetadataTypeSubscriptionUpdatedTrialMetadata:
		// subscriptionUpdatedEventMetadata.SubscriptionUpdatedTrialMetadata is populated
	case components.SubscriptionUpdatedEventMetadataTypeSubscriptionUpdatedSeatsMetadata:
		// subscriptionUpdatedEventMetadata.SubscriptionUpdatedSeatsMetadata is populated
	case components.SubscriptionUpdatedEventMetadataTypeSubscriptionUpdatedBillingPeriodMetadata:
		// subscriptionUpdatedEventMetadata.SubscriptionUpdatedBillingPeriodMetadata is populated
}
```

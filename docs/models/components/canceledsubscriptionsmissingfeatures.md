# CanceledSubscriptionsMissingFeatures


## Supported Types

### 

```go
canceledSubscriptionsMissingFeatures := components.CreateCanceledSubscriptionsMissingFeaturesInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsMissingFeatures := components.CreateCanceledSubscriptionsMissingFeaturesNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsMissingFeatures.Type {
	case components.CanceledSubscriptionsMissingFeaturesTypeInteger:
		// canceledSubscriptionsMissingFeatures.Integer is populated
	case components.CanceledSubscriptionsMissingFeaturesTypeNumber:
		// canceledSubscriptionsMissingFeatures.Number is populated
}
```

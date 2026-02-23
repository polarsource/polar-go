# CanceledSubscriptionsLowQuality


## Supported Types

### 

```go
canceledSubscriptionsLowQuality := components.CreateCanceledSubscriptionsLowQualityInteger(int64{/* values here */})
```

### 

```go
canceledSubscriptionsLowQuality := components.CreateCanceledSubscriptionsLowQualityNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch canceledSubscriptionsLowQuality.Type {
	case components.CanceledSubscriptionsLowQualityTypeInteger:
		// canceledSubscriptionsLowQuality.Integer is populated
	case components.CanceledSubscriptionsLowQualityTypeNumber:
		// canceledSubscriptionsLowQuality.Number is populated
}
```

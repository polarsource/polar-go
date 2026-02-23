# SubscriptionCreateExternalCustomerMetadata


## Supported Types

### 

```go
subscriptionCreateExternalCustomerMetadata := components.CreateSubscriptionCreateExternalCustomerMetadataStr(string{/* values here */})
```

### 

```go
subscriptionCreateExternalCustomerMetadata := components.CreateSubscriptionCreateExternalCustomerMetadataInteger(int64{/* values here */})
```

### 

```go
subscriptionCreateExternalCustomerMetadata := components.CreateSubscriptionCreateExternalCustomerMetadataNumber(float64{/* values here */})
```

### 

```go
subscriptionCreateExternalCustomerMetadata := components.CreateSubscriptionCreateExternalCustomerMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch subscriptionCreateExternalCustomerMetadata.Type {
	case components.SubscriptionCreateExternalCustomerMetadataTypeStr:
		// subscriptionCreateExternalCustomerMetadata.Str is populated
	case components.SubscriptionCreateExternalCustomerMetadataTypeInteger:
		// subscriptionCreateExternalCustomerMetadata.Integer is populated
	case components.SubscriptionCreateExternalCustomerMetadataTypeNumber:
		// subscriptionCreateExternalCustomerMetadata.Number is populated
	case components.SubscriptionCreateExternalCustomerMetadataTypeBoolean:
		// subscriptionCreateExternalCustomerMetadata.Boolean is populated
}
```

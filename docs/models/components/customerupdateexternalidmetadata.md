# CustomerUpdateExternalIDMetadata


## Supported Types

### 

```go
customerUpdateExternalIDMetadata := components.CreateCustomerUpdateExternalIDMetadataStr(string{/* values here */})
```

### 

```go
customerUpdateExternalIDMetadata := components.CreateCustomerUpdateExternalIDMetadataInteger(int64{/* values here */})
```

### 

```go
customerUpdateExternalIDMetadata := components.CreateCustomerUpdateExternalIDMetadataNumber(float64{/* values here */})
```

### 

```go
customerUpdateExternalIDMetadata := components.CreateCustomerUpdateExternalIDMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerUpdateExternalIDMetadata.Type {
	case components.CustomerUpdateExternalIDMetadataTypeStr:
		// customerUpdateExternalIDMetadata.Str is populated
	case components.CustomerUpdateExternalIDMetadataTypeInteger:
		// customerUpdateExternalIDMetadata.Integer is populated
	case components.CustomerUpdateExternalIDMetadataTypeNumber:
		// customerUpdateExternalIDMetadata.Number is populated
	case components.CustomerUpdateExternalIDMetadataTypeBoolean:
		// customerUpdateExternalIDMetadata.Boolean is populated
}
```

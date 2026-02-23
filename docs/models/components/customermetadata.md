# CustomerMetadata


## Supported Types

### 

```go
customerMetadata := components.CreateCustomerMetadataStr(string{/* values here */})
```

### 

```go
customerMetadata := components.CreateCustomerMetadataInteger(int64{/* values here */})
```

### 

```go
customerMetadata := components.CreateCustomerMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerMetadata.Type {
	case components.CustomerMetadataTypeStr:
		// customerMetadata.Str is populated
	case components.CustomerMetadataTypeInteger:
		// customerMetadata.Integer is populated
	case components.CustomerMetadataTypeBoolean:
		// customerMetadata.Boolean is populated
}
```

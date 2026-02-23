# CustomerUpdateMetadata


## Supported Types

### 

```go
customerUpdateMetadata := components.CreateCustomerUpdateMetadataStr(string{/* values here */})
```

### 

```go
customerUpdateMetadata := components.CreateCustomerUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
customerUpdateMetadata := components.CreateCustomerUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
customerUpdateMetadata := components.CreateCustomerUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerUpdateMetadata.Type {
	case components.CustomerUpdateMetadataTypeStr:
		// customerUpdateMetadata.Str is populated
	case components.CustomerUpdateMetadataTypeInteger:
		// customerUpdateMetadata.Integer is populated
	case components.CustomerUpdateMetadataTypeNumber:
		// customerUpdateMetadata.Number is populated
	case components.CustomerUpdateMetadataTypeBoolean:
		// customerUpdateMetadata.Boolean is populated
}
```

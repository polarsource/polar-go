# CustomerUpdatedFieldsMetadata


## Supported Types

### 

```go
customerUpdatedFieldsMetadata := components.CreateCustomerUpdatedFieldsMetadataStr(string{/* values here */})
```

### 

```go
customerUpdatedFieldsMetadata := components.CreateCustomerUpdatedFieldsMetadataInteger(int64{/* values here */})
```

### 

```go
customerUpdatedFieldsMetadata := components.CreateCustomerUpdatedFieldsMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerUpdatedFieldsMetadata.Type {
	case components.CustomerUpdatedFieldsMetadataTypeStr:
		// customerUpdatedFieldsMetadata.Str is populated
	case components.CustomerUpdatedFieldsMetadataTypeInteger:
		// customerUpdatedFieldsMetadata.Integer is populated
	case components.CustomerUpdatedFieldsMetadataTypeBoolean:
		// customerUpdatedFieldsMetadata.Boolean is populated
}
```

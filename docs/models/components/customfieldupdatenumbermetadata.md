# CustomFieldUpdateNumberMetadata


## Supported Types

### 

```go
customFieldUpdateNumberMetadata := components.CreateCustomFieldUpdateNumberMetadataStr(string{/* values here */})
```

### 

```go
customFieldUpdateNumberMetadata := components.CreateCustomFieldUpdateNumberMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldUpdateNumberMetadata := components.CreateCustomFieldUpdateNumberMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldUpdateNumberMetadata := components.CreateCustomFieldUpdateNumberMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldUpdateNumberMetadata.Type {
	case components.CustomFieldUpdateNumberMetadataTypeStr:
		// customFieldUpdateNumberMetadata.Str is populated
	case components.CustomFieldUpdateNumberMetadataTypeInteger:
		// customFieldUpdateNumberMetadata.Integer is populated
	case components.CustomFieldUpdateNumberMetadataTypeNumber:
		// customFieldUpdateNumberMetadata.Number is populated
	case components.CustomFieldUpdateNumberMetadataTypeBoolean:
		// customFieldUpdateNumberMetadata.Boolean is populated
}
```

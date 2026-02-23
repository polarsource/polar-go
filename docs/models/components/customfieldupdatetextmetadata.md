# CustomFieldUpdateTextMetadata


## Supported Types

### 

```go
customFieldUpdateTextMetadata := components.CreateCustomFieldUpdateTextMetadataStr(string{/* values here */})
```

### 

```go
customFieldUpdateTextMetadata := components.CreateCustomFieldUpdateTextMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldUpdateTextMetadata := components.CreateCustomFieldUpdateTextMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldUpdateTextMetadata := components.CreateCustomFieldUpdateTextMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldUpdateTextMetadata.Type {
	case components.CustomFieldUpdateTextMetadataTypeStr:
		// customFieldUpdateTextMetadata.Str is populated
	case components.CustomFieldUpdateTextMetadataTypeInteger:
		// customFieldUpdateTextMetadata.Integer is populated
	case components.CustomFieldUpdateTextMetadataTypeNumber:
		// customFieldUpdateTextMetadata.Number is populated
	case components.CustomFieldUpdateTextMetadataTypeBoolean:
		// customFieldUpdateTextMetadata.Boolean is populated
}
```

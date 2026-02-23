# CustomFieldUpdateDateMetadata


## Supported Types

### 

```go
customFieldUpdateDateMetadata := components.CreateCustomFieldUpdateDateMetadataStr(string{/* values here */})
```

### 

```go
customFieldUpdateDateMetadata := components.CreateCustomFieldUpdateDateMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldUpdateDateMetadata := components.CreateCustomFieldUpdateDateMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldUpdateDateMetadata := components.CreateCustomFieldUpdateDateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldUpdateDateMetadata.Type {
	case components.CustomFieldUpdateDateMetadataTypeStr:
		// customFieldUpdateDateMetadata.Str is populated
	case components.CustomFieldUpdateDateMetadataTypeInteger:
		// customFieldUpdateDateMetadata.Integer is populated
	case components.CustomFieldUpdateDateMetadataTypeNumber:
		// customFieldUpdateDateMetadata.Number is populated
	case components.CustomFieldUpdateDateMetadataTypeBoolean:
		// customFieldUpdateDateMetadata.Boolean is populated
}
```

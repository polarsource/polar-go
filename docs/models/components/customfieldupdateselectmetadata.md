# CustomFieldUpdateSelectMetadata


## Supported Types

### 

```go
customFieldUpdateSelectMetadata := components.CreateCustomFieldUpdateSelectMetadataStr(string{/* values here */})
```

### 

```go
customFieldUpdateSelectMetadata := components.CreateCustomFieldUpdateSelectMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldUpdateSelectMetadata := components.CreateCustomFieldUpdateSelectMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldUpdateSelectMetadata := components.CreateCustomFieldUpdateSelectMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldUpdateSelectMetadata.Type {
	case components.CustomFieldUpdateSelectMetadataTypeStr:
		// customFieldUpdateSelectMetadata.Str is populated
	case components.CustomFieldUpdateSelectMetadataTypeInteger:
		// customFieldUpdateSelectMetadata.Integer is populated
	case components.CustomFieldUpdateSelectMetadataTypeNumber:
		// customFieldUpdateSelectMetadata.Number is populated
	case components.CustomFieldUpdateSelectMetadataTypeBoolean:
		// customFieldUpdateSelectMetadata.Boolean is populated
}
```

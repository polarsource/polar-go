# CustomFieldCreateTextMetadata


## Supported Types

### 

```go
customFieldCreateTextMetadata := components.CreateCustomFieldCreateTextMetadataStr(string{/* values here */})
```

### 

```go
customFieldCreateTextMetadata := components.CreateCustomFieldCreateTextMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldCreateTextMetadata := components.CreateCustomFieldCreateTextMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldCreateTextMetadata := components.CreateCustomFieldCreateTextMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldCreateTextMetadata.Type {
	case components.CustomFieldCreateTextMetadataTypeStr:
		// customFieldCreateTextMetadata.Str is populated
	case components.CustomFieldCreateTextMetadataTypeInteger:
		// customFieldCreateTextMetadata.Integer is populated
	case components.CustomFieldCreateTextMetadataTypeNumber:
		// customFieldCreateTextMetadata.Number is populated
	case components.CustomFieldCreateTextMetadataTypeBoolean:
		// customFieldCreateTextMetadata.Boolean is populated
}
```

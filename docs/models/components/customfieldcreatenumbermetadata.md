# CustomFieldCreateNumberMetadata


## Supported Types

### 

```go
customFieldCreateNumberMetadata := components.CreateCustomFieldCreateNumberMetadataStr(string{/* values here */})
```

### 

```go
customFieldCreateNumberMetadata := components.CreateCustomFieldCreateNumberMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldCreateNumberMetadata := components.CreateCustomFieldCreateNumberMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldCreateNumberMetadata := components.CreateCustomFieldCreateNumberMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldCreateNumberMetadata.Type {
	case components.CustomFieldCreateNumberMetadataTypeStr:
		// customFieldCreateNumberMetadata.Str is populated
	case components.CustomFieldCreateNumberMetadataTypeInteger:
		// customFieldCreateNumberMetadata.Integer is populated
	case components.CustomFieldCreateNumberMetadataTypeNumber:
		// customFieldCreateNumberMetadata.Number is populated
	case components.CustomFieldCreateNumberMetadataTypeBoolean:
		// customFieldCreateNumberMetadata.Boolean is populated
}
```

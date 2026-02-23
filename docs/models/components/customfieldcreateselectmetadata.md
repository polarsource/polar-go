# CustomFieldCreateSelectMetadata


## Supported Types

### 

```go
customFieldCreateSelectMetadata := components.CreateCustomFieldCreateSelectMetadataStr(string{/* values here */})
```

### 

```go
customFieldCreateSelectMetadata := components.CreateCustomFieldCreateSelectMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldCreateSelectMetadata := components.CreateCustomFieldCreateSelectMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldCreateSelectMetadata := components.CreateCustomFieldCreateSelectMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldCreateSelectMetadata.Type {
	case components.CustomFieldCreateSelectMetadataTypeStr:
		// customFieldCreateSelectMetadata.Str is populated
	case components.CustomFieldCreateSelectMetadataTypeInteger:
		// customFieldCreateSelectMetadata.Integer is populated
	case components.CustomFieldCreateSelectMetadataTypeNumber:
		// customFieldCreateSelectMetadata.Number is populated
	case components.CustomFieldCreateSelectMetadataTypeBoolean:
		// customFieldCreateSelectMetadata.Boolean is populated
}
```

# CustomFieldCreateDateMetadata


## Supported Types

### 

```go
customFieldCreateDateMetadata := components.CreateCustomFieldCreateDateMetadataStr(string{/* values here */})
```

### 

```go
customFieldCreateDateMetadata := components.CreateCustomFieldCreateDateMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldCreateDateMetadata := components.CreateCustomFieldCreateDateMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldCreateDateMetadata := components.CreateCustomFieldCreateDateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldCreateDateMetadata.Type {
	case components.CustomFieldCreateDateMetadataTypeStr:
		// customFieldCreateDateMetadata.Str is populated
	case components.CustomFieldCreateDateMetadataTypeInteger:
		// customFieldCreateDateMetadata.Integer is populated
	case components.CustomFieldCreateDateMetadataTypeNumber:
		// customFieldCreateDateMetadata.Number is populated
	case components.CustomFieldCreateDateMetadataTypeBoolean:
		// customFieldCreateDateMetadata.Boolean is populated
}
```

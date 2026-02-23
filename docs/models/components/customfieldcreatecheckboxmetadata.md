# CustomFieldCreateCheckboxMetadata


## Supported Types

### 

```go
customFieldCreateCheckboxMetadata := components.CreateCustomFieldCreateCheckboxMetadataStr(string{/* values here */})
```

### 

```go
customFieldCreateCheckboxMetadata := components.CreateCustomFieldCreateCheckboxMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldCreateCheckboxMetadata := components.CreateCustomFieldCreateCheckboxMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldCreateCheckboxMetadata := components.CreateCustomFieldCreateCheckboxMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldCreateCheckboxMetadata.Type {
	case components.CustomFieldCreateCheckboxMetadataTypeStr:
		// customFieldCreateCheckboxMetadata.Str is populated
	case components.CustomFieldCreateCheckboxMetadataTypeInteger:
		// customFieldCreateCheckboxMetadata.Integer is populated
	case components.CustomFieldCreateCheckboxMetadataTypeNumber:
		// customFieldCreateCheckboxMetadata.Number is populated
	case components.CustomFieldCreateCheckboxMetadataTypeBoolean:
		// customFieldCreateCheckboxMetadata.Boolean is populated
}
```

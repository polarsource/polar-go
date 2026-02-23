# CustomFieldUpdateCheckboxMetadata


## Supported Types

### 

```go
customFieldUpdateCheckboxMetadata := components.CreateCustomFieldUpdateCheckboxMetadataStr(string{/* values here */})
```

### 

```go
customFieldUpdateCheckboxMetadata := components.CreateCustomFieldUpdateCheckboxMetadataInteger(int64{/* values here */})
```

### 

```go
customFieldUpdateCheckboxMetadata := components.CreateCustomFieldUpdateCheckboxMetadataNumber(float64{/* values here */})
```

### 

```go
customFieldUpdateCheckboxMetadata := components.CreateCustomFieldUpdateCheckboxMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldUpdateCheckboxMetadata.Type {
	case components.CustomFieldUpdateCheckboxMetadataTypeStr:
		// customFieldUpdateCheckboxMetadata.Str is populated
	case components.CustomFieldUpdateCheckboxMetadataTypeInteger:
		// customFieldUpdateCheckboxMetadata.Integer is populated
	case components.CustomFieldUpdateCheckboxMetadataTypeNumber:
		// customFieldUpdateCheckboxMetadata.Number is populated
	case components.CustomFieldUpdateCheckboxMetadataTypeBoolean:
		// customFieldUpdateCheckboxMetadata.Boolean is populated
}
```

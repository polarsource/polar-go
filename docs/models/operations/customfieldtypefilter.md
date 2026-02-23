# CustomFieldTypeFilter

Filter by custom field type.


## Supported Types

### CustomFieldType

```go
customFieldTypeFilter := operations.CreateCustomFieldTypeFilterCustomFieldType(components.CustomFieldType{/* values here */})
```

### 

```go
customFieldTypeFilter := operations.CreateCustomFieldTypeFilterArrayOfCustomFieldType([]components.CustomFieldType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldTypeFilter.Type {
	case operations.CustomFieldTypeFilterTypeCustomFieldType:
		// customFieldTypeFilter.CustomFieldType is populated
	case operations.CustomFieldTypeFilterTypeArrayOfCustomFieldType:
		// customFieldTypeFilter.ArrayOfCustomFieldType is populated
}
```

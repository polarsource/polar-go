# CustomField


## Supported Types

### CustomFieldCheckbox

```go
customField := components.CreateCustomFieldCheckbox(components.CustomFieldCheckbox{/* values here */})
```

### CustomFieldDate

```go
customField := components.CreateCustomFieldDate(components.CustomFieldDate{/* values here */})
```

### CustomFieldNumber

```go
customField := components.CreateCustomFieldNumber(components.CustomFieldNumber{/* values here */})
```

### CustomFieldSelect

```go
customField := components.CreateCustomFieldSelect(components.CustomFieldSelect{/* values here */})
```

### CustomFieldText

```go
customField := components.CreateCustomFieldText(components.CustomFieldText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customField.Type {
	case components.CustomFieldUnionTypeCheckbox:
		// customField.CustomFieldCheckbox is populated
	case components.CustomFieldUnionTypeDate:
		// customField.CustomFieldDate is populated
	case components.CustomFieldUnionTypeNumber:
		// customField.CustomFieldNumber is populated
	case components.CustomFieldUnionTypeSelect:
		// customField.CustomFieldSelect is populated
	case components.CustomFieldUnionTypeText:
		// customField.CustomFieldText is populated
}
```

# CustomFieldUpdate


## Supported Types

### CustomFieldUpdateCheckbox

```go
customFieldUpdate := components.CreateCustomFieldUpdateCheckbox(components.CustomFieldUpdateCheckbox{/* values here */})
```

### CustomFieldUpdateDate

```go
customFieldUpdate := components.CreateCustomFieldUpdateDate(components.CustomFieldUpdateDate{/* values here */})
```

### CustomFieldUpdateNumber

```go
customFieldUpdate := components.CreateCustomFieldUpdateNumber(components.CustomFieldUpdateNumber{/* values here */})
```

### CustomFieldUpdateSelect

```go
customFieldUpdate := components.CreateCustomFieldUpdateSelect(components.CustomFieldUpdateSelect{/* values here */})
```

### CustomFieldUpdateText

```go
customFieldUpdate := components.CreateCustomFieldUpdateText(components.CustomFieldUpdateText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldUpdate.Type {
	case components.CustomFieldUpdateTypeCheckbox:
		// customFieldUpdate.CustomFieldUpdateCheckbox is populated
	case components.CustomFieldUpdateTypeDate:
		// customFieldUpdate.CustomFieldUpdateDate is populated
	case components.CustomFieldUpdateTypeNumber:
		// customFieldUpdate.CustomFieldUpdateNumber is populated
	case components.CustomFieldUpdateTypeSelect:
		// customFieldUpdate.CustomFieldUpdateSelect is populated
	case components.CustomFieldUpdateTypeText:
		// customFieldUpdate.CustomFieldUpdateText is populated
}
```

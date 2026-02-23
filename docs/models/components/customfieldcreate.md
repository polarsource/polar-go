# CustomFieldCreate


## Supported Types

### CustomFieldCreateCheckbox

```go
customFieldCreate := components.CreateCustomFieldCreateCheckbox(components.CustomFieldCreateCheckbox{/* values here */})
```

### CustomFieldCreateDate

```go
customFieldCreate := components.CreateCustomFieldCreateDate(components.CustomFieldCreateDate{/* values here */})
```

### CustomFieldCreateNumber

```go
customFieldCreate := components.CreateCustomFieldCreateNumber(components.CustomFieldCreateNumber{/* values here */})
```

### CustomFieldCreateSelect

```go
customFieldCreate := components.CreateCustomFieldCreateSelect(components.CustomFieldCreateSelect{/* values here */})
```

### CustomFieldCreateText

```go
customFieldCreate := components.CreateCustomFieldCreateText(components.CustomFieldCreateText{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldCreate.Type {
	case components.CustomFieldCreateTypeCheckbox:
		// customFieldCreate.CustomFieldCreateCheckbox is populated
	case components.CustomFieldCreateTypeDate:
		// customFieldCreate.CustomFieldCreateDate is populated
	case components.CustomFieldCreateTypeNumber:
		// customFieldCreate.CustomFieldCreateNumber is populated
	case components.CustomFieldCreateTypeSelect:
		// customFieldCreate.CustomFieldCreateSelect is populated
	case components.CustomFieldCreateTypeText:
		// customFieldCreate.CustomFieldCreateText is populated
}
```

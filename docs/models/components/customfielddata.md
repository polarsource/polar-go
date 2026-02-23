# CustomFieldData


## Supported Types

### 

```go
customFieldData := components.CreateCustomFieldDataStr(string{/* values here */})
```

### 

```go
customFieldData := components.CreateCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
customFieldData := components.CreateCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
customFieldData := components.CreateCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldData.Type {
	case components.CustomFieldDataTypeStr:
		// customFieldData.Str is populated
	case components.CustomFieldDataTypeInteger:
		// customFieldData.Integer is populated
	case components.CustomFieldDataTypeBoolean:
		// customFieldData.Boolean is populated
	case components.CustomFieldDataTypeDateTime:
		// customFieldData.DateTime is populated
}
```

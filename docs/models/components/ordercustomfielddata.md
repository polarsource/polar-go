# OrderCustomFieldData


## Supported Types

### 

```go
orderCustomFieldData := components.CreateOrderCustomFieldDataStr(string{/* values here */})
```

### 

```go
orderCustomFieldData := components.CreateOrderCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
orderCustomFieldData := components.CreateOrderCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
orderCustomFieldData := components.CreateOrderCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch orderCustomFieldData.Type {
	case components.OrderCustomFieldDataTypeStr:
		// orderCustomFieldData.Str is populated
	case components.OrderCustomFieldDataTypeInteger:
		// orderCustomFieldData.Integer is populated
	case components.OrderCustomFieldDataTypeBoolean:
		// orderCustomFieldData.Boolean is populated
	case components.OrderCustomFieldDataTypeDateTime:
		// orderCustomFieldData.DateTime is populated
}
```

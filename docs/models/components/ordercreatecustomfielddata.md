# OrderCreateCustomFieldData


## Supported Types

### 

```go
orderCreateCustomFieldData := components.CreateOrderCreateCustomFieldDataStr(string{/* values here */})
```

### 

```go
orderCreateCustomFieldData := components.CreateOrderCreateCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
orderCreateCustomFieldData := components.CreateOrderCreateCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
orderCreateCustomFieldData := components.CreateOrderCreateCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch orderCreateCustomFieldData.Type {
	case components.OrderCreateCustomFieldDataTypeStr:
		// orderCreateCustomFieldData.Str is populated
	case components.OrderCreateCustomFieldDataTypeInteger:
		// orderCreateCustomFieldData.Integer is populated
	case components.OrderCreateCustomFieldDataTypeBoolean:
		// orderCreateCustomFieldData.Boolean is populated
	case components.OrderCreateCustomFieldDataTypeDateTime:
		// orderCreateCustomFieldData.DateTime is populated
}
```

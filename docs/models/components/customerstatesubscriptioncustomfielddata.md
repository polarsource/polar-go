# CustomerStateSubscriptionCustomFieldData


## Supported Types

### 

```go
customerStateSubscriptionCustomFieldData := components.CreateCustomerStateSubscriptionCustomFieldDataStr(string{/* values here */})
```

### 

```go
customerStateSubscriptionCustomFieldData := components.CreateCustomerStateSubscriptionCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
customerStateSubscriptionCustomFieldData := components.CreateCustomerStateSubscriptionCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
customerStateSubscriptionCustomFieldData := components.CreateCustomerStateSubscriptionCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerStateSubscriptionCustomFieldData.Type {
	case components.CustomerStateSubscriptionCustomFieldDataTypeStr:
		// customerStateSubscriptionCustomFieldData.Str is populated
	case components.CustomerStateSubscriptionCustomFieldDataTypeInteger:
		// customerStateSubscriptionCustomFieldData.Integer is populated
	case components.CustomerStateSubscriptionCustomFieldDataTypeBoolean:
		// customerStateSubscriptionCustomFieldData.Boolean is populated
	case components.CustomerStateSubscriptionCustomFieldDataTypeDateTime:
		// customerStateSubscriptionCustomFieldData.DateTime is populated
}
```

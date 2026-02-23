# CheckoutCreateCustomFieldData


## Supported Types

### 

```go
checkoutCreateCustomFieldData := components.CreateCheckoutCreateCustomFieldDataStr(string{/* values here */})
```

### 

```go
checkoutCreateCustomFieldData := components.CreateCheckoutCreateCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
checkoutCreateCustomFieldData := components.CreateCheckoutCreateCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
checkoutCreateCustomFieldData := components.CreateCheckoutCreateCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutCreateCustomFieldData.Type {
	case components.CheckoutCreateCustomFieldDataTypeStr:
		// checkoutCreateCustomFieldData.Str is populated
	case components.CheckoutCreateCustomFieldDataTypeInteger:
		// checkoutCreateCustomFieldData.Integer is populated
	case components.CheckoutCreateCustomFieldDataTypeBoolean:
		// checkoutCreateCustomFieldData.Boolean is populated
	case components.CheckoutCreateCustomFieldDataTypeDateTime:
		// checkoutCreateCustomFieldData.DateTime is populated
}
```

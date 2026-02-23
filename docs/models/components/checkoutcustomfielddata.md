# CheckoutCustomFieldData


## Supported Types

### 

```go
checkoutCustomFieldData := components.CreateCheckoutCustomFieldDataStr(string{/* values here */})
```

### 

```go
checkoutCustomFieldData := components.CreateCheckoutCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
checkoutCustomFieldData := components.CreateCheckoutCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
checkoutCustomFieldData := components.CreateCheckoutCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutCustomFieldData.Type {
	case components.CheckoutCustomFieldDataTypeStr:
		// checkoutCustomFieldData.Str is populated
	case components.CheckoutCustomFieldDataTypeInteger:
		// checkoutCustomFieldData.Integer is populated
	case components.CheckoutCustomFieldDataTypeBoolean:
		// checkoutCustomFieldData.Boolean is populated
	case components.CheckoutCustomFieldDataTypeDateTime:
		// checkoutCustomFieldData.DateTime is populated
}
```

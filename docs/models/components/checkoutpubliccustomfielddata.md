# CheckoutPublicCustomFieldData


## Supported Types

### 

```go
checkoutPublicCustomFieldData := components.CreateCheckoutPublicCustomFieldDataStr(string{/* values here */})
```

### 

```go
checkoutPublicCustomFieldData := components.CreateCheckoutPublicCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
checkoutPublicCustomFieldData := components.CreateCheckoutPublicCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
checkoutPublicCustomFieldData := components.CreateCheckoutPublicCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicCustomFieldData.Type {
	case components.CheckoutPublicCustomFieldDataTypeStr:
		// checkoutPublicCustomFieldData.Str is populated
	case components.CheckoutPublicCustomFieldDataTypeInteger:
		// checkoutPublicCustomFieldData.Integer is populated
	case components.CheckoutPublicCustomFieldDataTypeBoolean:
		// checkoutPublicCustomFieldData.Boolean is populated
	case components.CheckoutPublicCustomFieldDataTypeDateTime:
		// checkoutPublicCustomFieldData.DateTime is populated
}
```

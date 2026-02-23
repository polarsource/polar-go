# CheckoutPublicConfirmedCustomFieldData


## Supported Types

### 

```go
checkoutPublicConfirmedCustomFieldData := components.CreateCheckoutPublicConfirmedCustomFieldDataStr(string{/* values here */})
```

### 

```go
checkoutPublicConfirmedCustomFieldData := components.CreateCheckoutPublicConfirmedCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
checkoutPublicConfirmedCustomFieldData := components.CreateCheckoutPublicConfirmedCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
checkoutPublicConfirmedCustomFieldData := components.CreateCheckoutPublicConfirmedCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutPublicConfirmedCustomFieldData.Type {
	case components.CheckoutPublicConfirmedCustomFieldDataTypeStr:
		// checkoutPublicConfirmedCustomFieldData.Str is populated
	case components.CheckoutPublicConfirmedCustomFieldDataTypeInteger:
		// checkoutPublicConfirmedCustomFieldData.Integer is populated
	case components.CheckoutPublicConfirmedCustomFieldDataTypeBoolean:
		// checkoutPublicConfirmedCustomFieldData.Boolean is populated
	case components.CheckoutPublicConfirmedCustomFieldDataTypeDateTime:
		// checkoutPublicConfirmedCustomFieldData.DateTime is populated
}
```

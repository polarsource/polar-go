# CheckoutUpdatePublicCustomFieldData


## Supported Types

### 

```go
checkoutUpdatePublicCustomFieldData := components.CreateCheckoutUpdatePublicCustomFieldDataStr(string{/* values here */})
```

### 

```go
checkoutUpdatePublicCustomFieldData := components.CreateCheckoutUpdatePublicCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
checkoutUpdatePublicCustomFieldData := components.CreateCheckoutUpdatePublicCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
checkoutUpdatePublicCustomFieldData := components.CreateCheckoutUpdatePublicCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutUpdatePublicCustomFieldData.Type {
	case components.CheckoutUpdatePublicCustomFieldDataTypeStr:
		// checkoutUpdatePublicCustomFieldData.Str is populated
	case components.CheckoutUpdatePublicCustomFieldDataTypeInteger:
		// checkoutUpdatePublicCustomFieldData.Integer is populated
	case components.CheckoutUpdatePublicCustomFieldDataTypeBoolean:
		// checkoutUpdatePublicCustomFieldData.Boolean is populated
	case components.CheckoutUpdatePublicCustomFieldDataTypeDateTime:
		// checkoutUpdatePublicCustomFieldData.DateTime is populated
}
```

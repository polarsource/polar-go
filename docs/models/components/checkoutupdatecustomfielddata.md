# CheckoutUpdateCustomFieldData


## Supported Types

### 

```go
checkoutUpdateCustomFieldData := components.CreateCheckoutUpdateCustomFieldDataStr(string{/* values here */})
```

### 

```go
checkoutUpdateCustomFieldData := components.CreateCheckoutUpdateCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
checkoutUpdateCustomFieldData := components.CreateCheckoutUpdateCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
checkoutUpdateCustomFieldData := components.CreateCheckoutUpdateCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutUpdateCustomFieldData.Type {
	case components.CheckoutUpdateCustomFieldDataTypeStr:
		// checkoutUpdateCustomFieldData.Str is populated
	case components.CheckoutUpdateCustomFieldDataTypeInteger:
		// checkoutUpdateCustomFieldData.Integer is populated
	case components.CheckoutUpdateCustomFieldDataTypeBoolean:
		// checkoutUpdateCustomFieldData.Boolean is populated
	case components.CheckoutUpdateCustomFieldDataTypeDateTime:
		// checkoutUpdateCustomFieldData.DateTime is populated
}
```

# CheckoutConfirmStripeCustomFieldData


## Supported Types

### 

```go
checkoutConfirmStripeCustomFieldData := components.CreateCheckoutConfirmStripeCustomFieldDataStr(string{/* values here */})
```

### 

```go
checkoutConfirmStripeCustomFieldData := components.CreateCheckoutConfirmStripeCustomFieldDataInteger(int64{/* values here */})
```

### 

```go
checkoutConfirmStripeCustomFieldData := components.CreateCheckoutConfirmStripeCustomFieldDataBoolean(bool{/* values here */})
```

### 

```go
checkoutConfirmStripeCustomFieldData := components.CreateCheckoutConfirmStripeCustomFieldDataDateTime(time.Time{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutConfirmStripeCustomFieldData.Type {
	case components.CheckoutConfirmStripeCustomFieldDataTypeStr:
		// checkoutConfirmStripeCustomFieldData.Str is populated
	case components.CheckoutConfirmStripeCustomFieldDataTypeInteger:
		// checkoutConfirmStripeCustomFieldData.Integer is populated
	case components.CheckoutConfirmStripeCustomFieldDataTypeBoolean:
		// checkoutConfirmStripeCustomFieldData.Boolean is populated
	case components.CheckoutConfirmStripeCustomFieldDataTypeDateTime:
		// checkoutConfirmStripeCustomFieldData.DateTime is populated
}
```

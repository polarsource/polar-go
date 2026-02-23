# CheckoutIDFilter

Filter by checkout ID.


## Supported Types

### 

```go
checkoutIDFilter := operations.CreateCheckoutIDFilterStr(string{/* values here */})
```

### 

```go
checkoutIDFilter := operations.CreateCheckoutIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutIDFilter.Type {
	case operations.CheckoutIDFilterTypeStr:
		// checkoutIDFilter.Str is populated
	case operations.CheckoutIDFilterTypeArrayOfStr:
		// checkoutIDFilter.ArrayOfStr is populated
}
```

# PaymentsListQueryParamCheckoutIDFilter

Filter by checkout ID.


## Supported Types

### 

```go
paymentsListQueryParamCheckoutIDFilter := operations.CreatePaymentsListQueryParamCheckoutIDFilterStr(string{/* values here */})
```

### 

```go
paymentsListQueryParamCheckoutIDFilter := operations.CreatePaymentsListQueryParamCheckoutIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentsListQueryParamCheckoutIDFilter.Type {
	case operations.PaymentsListQueryParamCheckoutIDFilterTypeStr:
		// paymentsListQueryParamCheckoutIDFilter.Str is populated
	case operations.PaymentsListQueryParamCheckoutIDFilterTypeArrayOfStr:
		// paymentsListQueryParamCheckoutIDFilter.ArrayOfStr is populated
}
```

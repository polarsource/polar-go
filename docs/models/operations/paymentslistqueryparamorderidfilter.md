# PaymentsListQueryParamOrderIDFilter

Filter by order ID.


## Supported Types

### 

```go
paymentsListQueryParamOrderIDFilter := operations.CreatePaymentsListQueryParamOrderIDFilterStr(string{/* values here */})
```

### 

```go
paymentsListQueryParamOrderIDFilter := operations.CreatePaymentsListQueryParamOrderIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentsListQueryParamOrderIDFilter.Type {
	case operations.PaymentsListQueryParamOrderIDFilterTypeStr:
		// paymentsListQueryParamOrderIDFilter.Str is populated
	case operations.PaymentsListQueryParamOrderIDFilterTypeArrayOfStr:
		// paymentsListQueryParamOrderIDFilter.ArrayOfStr is populated
}
```

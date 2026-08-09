# PaymentsListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
paymentsListQueryParamCustomerIDFilter := operations.CreatePaymentsListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
paymentsListQueryParamCustomerIDFilter := operations.CreatePaymentsListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentsListQueryParamCustomerIDFilter.Type {
	case operations.PaymentsListQueryParamCustomerIDFilterTypeStr:
		// paymentsListQueryParamCustomerIDFilter.Str is populated
	case operations.PaymentsListQueryParamCustomerIDFilterTypeArrayOfStr:
		// paymentsListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

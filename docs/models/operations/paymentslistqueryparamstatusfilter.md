# PaymentsListQueryParamStatusFilter

Filter by payment status.


## Supported Types

### PaymentStatus

```go
paymentsListQueryParamStatusFilter := operations.CreatePaymentsListQueryParamStatusFilterPaymentStatus(components.PaymentStatus{/* values here */})
```

### 

```go
paymentsListQueryParamStatusFilter := operations.CreatePaymentsListQueryParamStatusFilterArrayOfPaymentStatus([]components.PaymentStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentsListQueryParamStatusFilter.Type {
	case operations.PaymentsListQueryParamStatusFilterTypePaymentStatus:
		// paymentsListQueryParamStatusFilter.PaymentStatus is populated
	case operations.PaymentsListQueryParamStatusFilterTypeArrayOfPaymentStatus:
		// paymentsListQueryParamStatusFilter.ArrayOfPaymentStatus is populated
}
```

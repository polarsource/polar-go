# Payment


## Supported Types

### CardPayment

```go
payment := components.CreatePaymentCardPayment(components.CardPayment{/* values here */})
```

### GenericPayment

```go
payment := components.CreatePaymentGenericPayment(components.GenericPayment{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch payment.Type {
	case components.PaymentTypeCardPayment:
		// payment.CardPayment is populated
	case components.PaymentTypeGenericPayment:
		// payment.GenericPayment is populated
}
```

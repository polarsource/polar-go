# PaymentMethod


## Supported Types

### CustomerPaymentMethodCard

```go
paymentMethod := components.CreatePaymentMethodCustomerPaymentMethodCard(components.CustomerPaymentMethodCard{/* values here */})
```

### CustomerPaymentMethodGeneric

```go
paymentMethod := components.CreatePaymentMethodCustomerPaymentMethodGeneric(components.CustomerPaymentMethodGeneric{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch paymentMethod.Type {
	case components.PaymentMethodTypeCustomerPaymentMethodCard:
		// paymentMethod.CustomerPaymentMethodCard is populated
	case components.PaymentMethodTypeCustomerPaymentMethodGeneric:
		// paymentMethod.CustomerPaymentMethodGeneric is populated
}
```

# CustomerPaymentMethod


## Supported Types

### PaymentMethodCard

```go
customerPaymentMethod := components.CreateCustomerPaymentMethodPaymentMethodCard(components.PaymentMethodCard{/* values here */})
```

### PaymentMethodGeneric

```go
customerPaymentMethod := components.CreateCustomerPaymentMethodPaymentMethodGeneric(components.PaymentMethodGeneric{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPaymentMethod.Type {
	case components.CustomerPaymentMethodTypePaymentMethodCard:
		// customerPaymentMethod.PaymentMethodCard is populated
	case components.CustomerPaymentMethodTypePaymentMethodGeneric:
		// customerPaymentMethod.PaymentMethodGeneric is populated
}
```

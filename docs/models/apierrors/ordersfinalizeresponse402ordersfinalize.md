# OrdersFinalizeResponse402OrdersFinalize

The charge failed, or requires customer authentication (e.g. a 3DS challenge) that can't be completed off-session.


## Supported Types

### PaymentFailed

```go
ordersFinalizeResponse402OrdersFinalize := apierrors.CreateOrdersFinalizeResponse402OrdersFinalizePaymentFailed(components.PaymentFailed{/* values here */})
```

### PaymentActionRequired

```go
ordersFinalizeResponse402OrdersFinalize := apierrors.CreateOrdersFinalizeResponse402OrdersFinalizePaymentActionRequired(components.PaymentActionRequired{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersFinalizeResponse402OrdersFinalize.Type {
	case apierrors.OrdersFinalizeResponse402OrdersFinalizeTypePaymentFailed:
		// ordersFinalizeResponse402OrdersFinalize.PaymentFailed is populated
	case apierrors.OrdersFinalizeResponse402OrdersFinalizeTypePaymentActionRequired:
		// ordersFinalizeResponse402OrdersFinalize.PaymentActionRequired is populated
}
```

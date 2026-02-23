# CheckoutForbiddenError


## Supported Types

### AlreadyActiveSubscriptionError

```go
checkoutForbiddenError := apierrors.CreateCheckoutForbiddenErrorAlreadyActiveSubscriptionError(components.AlreadyActiveSubscriptionError{/* values here */})
```

### NotOpenCheckout

```go
checkoutForbiddenError := apierrors.CreateCheckoutForbiddenErrorNotOpenCheckout(components.NotOpenCheckout{/* values here */})
```

### PaymentNotReady

```go
checkoutForbiddenError := apierrors.CreateCheckoutForbiddenErrorPaymentNotReady(components.PaymentNotReady{/* values here */})
```

### TrialAlreadyRedeemed

```go
checkoutForbiddenError := apierrors.CreateCheckoutForbiddenErrorTrialAlreadyRedeemed(components.TrialAlreadyRedeemed{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutForbiddenError.Type {
	case apierrors.CheckoutForbiddenErrorTypeAlreadyActiveSubscriptionError:
		// checkoutForbiddenError.AlreadyActiveSubscriptionError is populated
	case apierrors.CheckoutForbiddenErrorTypeNotOpenCheckout:
		// checkoutForbiddenError.NotOpenCheckout is populated
	case apierrors.CheckoutForbiddenErrorTypePaymentNotReady:
		// checkoutForbiddenError.PaymentNotReady is populated
	case apierrors.CheckoutForbiddenErrorTypeTrialAlreadyRedeemed:
		// checkoutForbiddenError.TrialAlreadyRedeemed is populated
}
```

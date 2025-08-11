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


# CustomerPaymentMethodCreateResponse


## Supported Types

### CustomerPaymentMethodCreateRequiresActionResponse

```go
customerPaymentMethodCreateResponse := components.CreateCustomerPaymentMethodCreateResponseRequiresAction(components.CustomerPaymentMethodCreateRequiresActionResponse{/* values here */})
```

### CustomerPaymentMethodCreateSucceededResponse

```go
customerPaymentMethodCreateResponse := components.CreateCustomerPaymentMethodCreateResponseSucceeded(components.CustomerPaymentMethodCreateSucceededResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPaymentMethodCreateResponse.Type {
	case components.CustomerPaymentMethodCreateResponseTypeRequiresAction:
		// customerPaymentMethodCreateResponse.CustomerPaymentMethodCreateRequiresActionResponse is populated
	case components.CustomerPaymentMethodCreateResponseTypeSucceeded:
		// customerPaymentMethodCreateResponse.CustomerPaymentMethodCreateSucceededResponse is populated
}
```

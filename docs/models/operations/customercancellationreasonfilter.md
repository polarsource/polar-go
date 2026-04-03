# CustomerCancellationReasonFilter

Filter by customer cancellation reason.


## Supported Types

### CustomerCancellationReason

```go
customerCancellationReasonFilter := operations.CreateCustomerCancellationReasonFilterCustomerCancellationReason(components.CustomerCancellationReason{/* values here */})
```

### 

```go
customerCancellationReasonFilter := operations.CreateCustomerCancellationReasonFilterArrayOfCustomerCancellationReason([]components.CustomerCancellationReason{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerCancellationReasonFilter.Type {
	case operations.CustomerCancellationReasonFilterTypeCustomerCancellationReason:
		// customerCancellationReasonFilter.CustomerCancellationReason is populated
	case operations.CustomerCancellationReasonFilterTypeArrayOfCustomerCancellationReason:
		// customerCancellationReasonFilter.ArrayOfCustomerCancellationReason is populated
}
```

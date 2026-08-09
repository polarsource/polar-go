# CheckoutsListQueryParamStatusFilter

Filter by checkout session status.


## Supported Types

### CheckoutStatus

```go
checkoutsListQueryParamStatusFilter := operations.CreateCheckoutsListQueryParamStatusFilterCheckoutStatus(components.CheckoutStatus{/* values here */})
```

### 

```go
checkoutsListQueryParamStatusFilter := operations.CreateCheckoutsListQueryParamStatusFilterArrayOfCheckoutStatus([]components.CheckoutStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutsListQueryParamStatusFilter.Type {
	case operations.CheckoutsListQueryParamStatusFilterTypeCheckoutStatus:
		// checkoutsListQueryParamStatusFilter.CheckoutStatus is populated
	case operations.CheckoutsListQueryParamStatusFilterTypeArrayOfCheckoutStatus:
		// checkoutsListQueryParamStatusFilter.ArrayOfCheckoutStatus is populated
}
```

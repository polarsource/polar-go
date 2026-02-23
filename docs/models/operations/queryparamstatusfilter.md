# QueryParamStatusFilter

Filter by checkout session status.


## Supported Types

### CheckoutStatus

```go
queryParamStatusFilter := operations.CreateQueryParamStatusFilterCheckoutStatus(components.CheckoutStatus{/* values here */})
```

### 

```go
queryParamStatusFilter := operations.CreateQueryParamStatusFilterArrayOfCheckoutStatus([]components.CheckoutStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamStatusFilter.Type {
	case operations.QueryParamStatusFilterTypeCheckoutStatus:
		// queryParamStatusFilter.CheckoutStatus is populated
	case operations.QueryParamStatusFilterTypeArrayOfCheckoutStatus:
		// queryParamStatusFilter.ArrayOfCheckoutStatus is populated
}
```

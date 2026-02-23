# OrdersListQueryParamExternalCustomerIDFilter

Filter by customer external ID.


## Supported Types

### 

```go
ordersListQueryParamExternalCustomerIDFilter := operations.CreateOrdersListQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
ordersListQueryParamExternalCustomerIDFilter := operations.CreateOrdersListQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersListQueryParamExternalCustomerIDFilter.Type {
	case operations.OrdersListQueryParamExternalCustomerIDFilterTypeStr:
		// ordersListQueryParamExternalCustomerIDFilter.Str is populated
	case operations.OrdersListQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// ordersListQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```

# OrdersListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
ordersListQueryParamCustomerIDFilter := operations.CreateOrdersListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
ordersListQueryParamCustomerIDFilter := operations.CreateOrdersListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersListQueryParamCustomerIDFilter.Type {
	case operations.OrdersListQueryParamCustomerIDFilterTypeStr:
		// ordersListQueryParamCustomerIDFilter.Str is populated
	case operations.OrdersListQueryParamCustomerIDFilterTypeArrayOfStr:
		// ordersListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

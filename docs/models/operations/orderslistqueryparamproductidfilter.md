# OrdersListQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
ordersListQueryParamProductIDFilter := operations.CreateOrdersListQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
ordersListQueryParamProductIDFilter := operations.CreateOrdersListQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersListQueryParamProductIDFilter.Type {
	case operations.OrdersListQueryParamProductIDFilterTypeStr:
		// ordersListQueryParamProductIDFilter.Str is populated
	case operations.OrdersListQueryParamProductIDFilterTypeArrayOfStr:
		// ordersListQueryParamProductIDFilter.ArrayOfStr is populated
}
```

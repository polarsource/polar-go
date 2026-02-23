# OrdersListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
ordersListQueryParamOrganizationIDFilter := operations.CreateOrdersListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
ordersListQueryParamOrganizationIDFilter := operations.CreateOrdersListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersListQueryParamOrganizationIDFilter.Type {
	case operations.OrdersListQueryParamOrganizationIDFilterTypeStr:
		// ordersListQueryParamOrganizationIDFilter.Str is populated
	case operations.OrdersListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// ordersListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

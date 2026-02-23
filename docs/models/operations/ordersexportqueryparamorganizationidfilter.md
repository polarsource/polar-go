# OrdersExportQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
ordersExportQueryParamOrganizationIDFilter := operations.CreateOrdersExportQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
ordersExportQueryParamOrganizationIDFilter := operations.CreateOrdersExportQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersExportQueryParamOrganizationIDFilter.Type {
	case operations.OrdersExportQueryParamOrganizationIDFilterTypeStr:
		// ordersExportQueryParamOrganizationIDFilter.Str is populated
	case operations.OrdersExportQueryParamOrganizationIDFilterTypeArrayOfStr:
		// ordersExportQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

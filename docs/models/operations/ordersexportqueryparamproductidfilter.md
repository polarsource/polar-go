# OrdersExportQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
ordersExportQueryParamProductIDFilter := operations.CreateOrdersExportQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
ordersExportQueryParamProductIDFilter := operations.CreateOrdersExportQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersExportQueryParamProductIDFilter.Type {
	case operations.OrdersExportQueryParamProductIDFilterTypeStr:
		// ordersExportQueryParamProductIDFilter.Str is populated
	case operations.OrdersExportQueryParamProductIDFilterTypeArrayOfStr:
		// ordersExportQueryParamProductIDFilter.ArrayOfStr is populated
}
```

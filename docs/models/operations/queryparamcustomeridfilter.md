# QueryParamCustomerIDFilter

Filter by customer.


## Supported Types

### 

```go
queryParamCustomerIDFilter := operations.CreateQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
queryParamCustomerIDFilter := operations.CreateQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamCustomerIDFilter.Type {
	case operations.QueryParamCustomerIDFilterTypeStr:
		// queryParamCustomerIDFilter.Str is populated
	case operations.QueryParamCustomerIDFilterTypeArrayOfStr:
		// queryParamCustomerIDFilter.ArrayOfStr is populated
}
```

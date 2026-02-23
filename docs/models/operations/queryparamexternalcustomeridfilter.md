# QueryParamExternalCustomerIDFilter

Filter by customer external ID.


## Supported Types

### 

```go
queryParamExternalCustomerIDFilter := operations.CreateQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
queryParamExternalCustomerIDFilter := operations.CreateQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamExternalCustomerIDFilter.Type {
	case operations.QueryParamExternalCustomerIDFilterTypeStr:
		// queryParamExternalCustomerIDFilter.Str is populated
	case operations.QueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// queryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```

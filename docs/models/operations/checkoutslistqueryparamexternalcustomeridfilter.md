# CheckoutsListQueryParamExternalCustomerIDFilter

Filter by customer external ID.


## Supported Types

### 

```go
checkoutsListQueryParamExternalCustomerIDFilter := operations.CreateCheckoutsListQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
checkoutsListQueryParamExternalCustomerIDFilter := operations.CreateCheckoutsListQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutsListQueryParamExternalCustomerIDFilter.Type {
	case operations.CheckoutsListQueryParamExternalCustomerIDFilterTypeStr:
		// checkoutsListQueryParamExternalCustomerIDFilter.Str is populated
	case operations.CheckoutsListQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// checkoutsListQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```

# CheckoutsListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
checkoutsListQueryParamCustomerIDFilter := operations.CreateCheckoutsListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
checkoutsListQueryParamCustomerIDFilter := operations.CreateCheckoutsListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutsListQueryParamCustomerIDFilter.Type {
	case operations.CheckoutsListQueryParamCustomerIDFilterTypeStr:
		// checkoutsListQueryParamCustomerIDFilter.Str is populated
	case operations.CheckoutsListQueryParamCustomerIDFilterTypeArrayOfStr:
		// checkoutsListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

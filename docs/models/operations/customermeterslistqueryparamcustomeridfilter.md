# CustomerMetersListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
customerMetersListQueryParamCustomerIDFilter := operations.CreateCustomerMetersListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
customerMetersListQueryParamCustomerIDFilter := operations.CreateCustomerMetersListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerMetersListQueryParamCustomerIDFilter.Type {
	case operations.CustomerMetersListQueryParamCustomerIDFilterTypeStr:
		// customerMetersListQueryParamCustomerIDFilter.Str is populated
	case operations.CustomerMetersListQueryParamCustomerIDFilterTypeArrayOfStr:
		// customerMetersListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

# CustomerMetersListQueryParamExternalCustomerIDFilter

Filter by external customer ID.


## Supported Types

### 

```go
customerMetersListQueryParamExternalCustomerIDFilter := operations.CreateCustomerMetersListQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
customerMetersListQueryParamExternalCustomerIDFilter := operations.CreateCustomerMetersListQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerMetersListQueryParamExternalCustomerIDFilter.Type {
	case operations.CustomerMetersListQueryParamExternalCustomerIDFilterTypeStr:
		// customerMetersListQueryParamExternalCustomerIDFilter.Str is populated
	case operations.CustomerMetersListQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// customerMetersListQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```

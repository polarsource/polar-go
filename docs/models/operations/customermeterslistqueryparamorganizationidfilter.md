# CustomerMetersListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
customerMetersListQueryParamOrganizationIDFilter := operations.CreateCustomerMetersListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
customerMetersListQueryParamOrganizationIDFilter := operations.CreateCustomerMetersListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerMetersListQueryParamOrganizationIDFilter.Type {
	case operations.CustomerMetersListQueryParamOrganizationIDFilterTypeStr:
		// customerMetersListQueryParamOrganizationIDFilter.Str is populated
	case operations.CustomerMetersListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// customerMetersListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

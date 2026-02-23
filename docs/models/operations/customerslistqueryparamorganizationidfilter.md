# CustomersListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
customersListQueryParamOrganizationIDFilter := operations.CreateCustomersListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
customersListQueryParamOrganizationIDFilter := operations.CreateCustomersListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customersListQueryParamOrganizationIDFilter.Type {
	case operations.CustomersListQueryParamOrganizationIDFilterTypeStr:
		// customersListQueryParamOrganizationIDFilter.Str is populated
	case operations.CustomersListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// customersListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

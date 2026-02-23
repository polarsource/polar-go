# CustomersExportQueryParamOrganizationID

Filter by organization ID.


## Supported Types

### 

```go
customersExportQueryParamOrganizationID := operations.CreateCustomersExportQueryParamOrganizationIDStr(string{/* values here */})
```

### 

```go
customersExportQueryParamOrganizationID := operations.CreateCustomersExportQueryParamOrganizationIDArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customersExportQueryParamOrganizationID.Type {
	case operations.CustomersExportQueryParamOrganizationIDTypeStr:
		// customersExportQueryParamOrganizationID.Str is populated
	case operations.CustomersExportQueryParamOrganizationIDTypeArrayOfStr:
		// customersExportQueryParamOrganizationID.ArrayOfStr is populated
}
```

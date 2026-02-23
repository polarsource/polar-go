# CustomFieldsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
customFieldsListQueryParamOrganizationIDFilter := operations.CreateCustomFieldsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
customFieldsListQueryParamOrganizationIDFilter := operations.CreateCustomFieldsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customFieldsListQueryParamOrganizationIDFilter.Type {
	case operations.CustomFieldsListQueryParamOrganizationIDFilterTypeStr:
		// customFieldsListQueryParamOrganizationIDFilter.Str is populated
	case operations.CustomFieldsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// customFieldsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

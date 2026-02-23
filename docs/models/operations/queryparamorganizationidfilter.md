# QueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
queryParamOrganizationIDFilter := operations.CreateQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
queryParamOrganizationIDFilter := operations.CreateQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamOrganizationIDFilter.Type {
	case operations.QueryParamOrganizationIDFilterTypeStr:
		// queryParamOrganizationIDFilter.Str is populated
	case operations.QueryParamOrganizationIDFilterTypeArrayOfStr:
		// queryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

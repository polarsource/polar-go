# QueryParamOrganizationID

Filter by organization ID.


## Supported Types

### 

```go
queryParamOrganizationID := operations.CreateQueryParamOrganizationIDStr(string{/* values here */})
```

### 

```go
queryParamOrganizationID := operations.CreateQueryParamOrganizationIDArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamOrganizationID.Type {
	case operations.QueryParamOrganizationIDTypeStr:
		// queryParamOrganizationID.Str is populated
	case operations.QueryParamOrganizationIDTypeArrayOfStr:
		// queryParamOrganizationID.ArrayOfStr is populated
}
```

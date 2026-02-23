# CheckoutsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
checkoutsListQueryParamOrganizationIDFilter := operations.CreateCheckoutsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
checkoutsListQueryParamOrganizationIDFilter := operations.CreateCheckoutsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutsListQueryParamOrganizationIDFilter.Type {
	case operations.CheckoutsListQueryParamOrganizationIDFilterTypeStr:
		// checkoutsListQueryParamOrganizationIDFilter.Str is populated
	case operations.CheckoutsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// checkoutsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

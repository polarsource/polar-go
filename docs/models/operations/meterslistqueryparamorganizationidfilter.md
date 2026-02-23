# MetersListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
metersListQueryParamOrganizationIDFilter := operations.CreateMetersListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
metersListQueryParamOrganizationIDFilter := operations.CreateMetersListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metersListQueryParamOrganizationIDFilter.Type {
	case operations.MetersListQueryParamOrganizationIDFilterTypeStr:
		// metersListQueryParamOrganizationIDFilter.Str is populated
	case operations.MetersListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// metersListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

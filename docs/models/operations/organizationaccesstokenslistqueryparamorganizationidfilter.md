# OrganizationAccessTokensListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
organizationAccessTokensListQueryParamOrganizationIDFilter := operations.CreateOrganizationAccessTokensListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
organizationAccessTokensListQueryParamOrganizationIDFilter := operations.CreateOrganizationAccessTokensListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch organizationAccessTokensListQueryParamOrganizationIDFilter.Type {
	case operations.OrganizationAccessTokensListQueryParamOrganizationIDFilterTypeStr:
		// organizationAccessTokensListQueryParamOrganizationIDFilter.Str is populated
	case operations.OrganizationAccessTokensListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// organizationAccessTokensListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

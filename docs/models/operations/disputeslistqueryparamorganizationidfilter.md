# DisputesListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
disputesListQueryParamOrganizationIDFilter := operations.CreateDisputesListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
disputesListQueryParamOrganizationIDFilter := operations.CreateDisputesListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch disputesListQueryParamOrganizationIDFilter.Type {
	case operations.DisputesListQueryParamOrganizationIDFilterTypeStr:
		// disputesListQueryParamOrganizationIDFilter.Str is populated
	case operations.DisputesListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// disputesListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

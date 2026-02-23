# RefundsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
refundsListQueryParamOrganizationIDFilter := operations.CreateRefundsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
refundsListQueryParamOrganizationIDFilter := operations.CreateRefundsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch refundsListQueryParamOrganizationIDFilter.Type {
	case operations.RefundsListQueryParamOrganizationIDFilterTypeStr:
		// refundsListQueryParamOrganizationIDFilter.Str is populated
	case operations.RefundsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// refundsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

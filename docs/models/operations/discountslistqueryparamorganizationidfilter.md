# DiscountsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
discountsListQueryParamOrganizationIDFilter := operations.CreateDiscountsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
discountsListQueryParamOrganizationIDFilter := operations.CreateDiscountsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountsListQueryParamOrganizationIDFilter.Type {
	case operations.DiscountsListQueryParamOrganizationIDFilterTypeStr:
		// discountsListQueryParamOrganizationIDFilter.Str is populated
	case operations.DiscountsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// discountsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

# BenefitGrantsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
benefitGrantsListQueryParamOrganizationIDFilter := operations.CreateBenefitGrantsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
benefitGrantsListQueryParamOrganizationIDFilter := operations.CreateBenefitGrantsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitGrantsListQueryParamOrganizationIDFilter.Type {
	case operations.BenefitGrantsListQueryParamOrganizationIDFilterTypeStr:
		// benefitGrantsListQueryParamOrganizationIDFilter.Str is populated
	case operations.BenefitGrantsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// benefitGrantsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

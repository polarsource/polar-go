# CustomerPortalBenefitGrantsListQueryParamBenefitIDFilter

Filter by benefit ID.


## Supported Types

### 

```go
customerPortalBenefitGrantsListQueryParamBenefitIDFilter := operations.CreateCustomerPortalBenefitGrantsListQueryParamBenefitIDFilterStr(string{/* values here */})
```

### 

```go
customerPortalBenefitGrantsListQueryParamBenefitIDFilter := operations.CreateCustomerPortalBenefitGrantsListQueryParamBenefitIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalBenefitGrantsListQueryParamBenefitIDFilter.Type {
	case operations.CustomerPortalBenefitGrantsListQueryParamBenefitIDFilterTypeStr:
		// customerPortalBenefitGrantsListQueryParamBenefitIDFilter.Str is populated
	case operations.CustomerPortalBenefitGrantsListQueryParamBenefitIDFilterTypeArrayOfStr:
		// customerPortalBenefitGrantsListQueryParamBenefitIDFilter.ArrayOfStr is populated
}
```

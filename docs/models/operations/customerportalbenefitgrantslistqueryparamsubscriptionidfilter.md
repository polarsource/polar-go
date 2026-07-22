# CustomerPortalBenefitGrantsListQueryParamSubscriptionIDFilter

Filter by subscription ID.


## Supported Types

### 

```go
customerPortalBenefitGrantsListQueryParamSubscriptionIDFilter := operations.CreateCustomerPortalBenefitGrantsListQueryParamSubscriptionIDFilterStr(string{/* values here */})
```

### 

```go
customerPortalBenefitGrantsListQueryParamSubscriptionIDFilter := operations.CreateCustomerPortalBenefitGrantsListQueryParamSubscriptionIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalBenefitGrantsListQueryParamSubscriptionIDFilter.Type {
	case operations.CustomerPortalBenefitGrantsListQueryParamSubscriptionIDFilterTypeStr:
		// customerPortalBenefitGrantsListQueryParamSubscriptionIDFilter.Str is populated
	case operations.CustomerPortalBenefitGrantsListQueryParamSubscriptionIDFilterTypeArrayOfStr:
		// customerPortalBenefitGrantsListQueryParamSubscriptionIDFilter.ArrayOfStr is populated
}
```

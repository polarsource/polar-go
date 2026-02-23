# CustomerPortalBenefitGrantsListQueryParamOrderIDFilter

Filter by order ID.


## Supported Types

### 

```go
customerPortalBenefitGrantsListQueryParamOrderIDFilter := operations.CreateCustomerPortalBenefitGrantsListQueryParamOrderIDFilterStr(string{/* values here */})
```

### 

```go
customerPortalBenefitGrantsListQueryParamOrderIDFilter := operations.CreateCustomerPortalBenefitGrantsListQueryParamOrderIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalBenefitGrantsListQueryParamOrderIDFilter.Type {
	case operations.CustomerPortalBenefitGrantsListQueryParamOrderIDFilterTypeStr:
		// customerPortalBenefitGrantsListQueryParamOrderIDFilter.Str is populated
	case operations.CustomerPortalBenefitGrantsListQueryParamOrderIDFilterTypeArrayOfStr:
		// customerPortalBenefitGrantsListQueryParamOrderIDFilter.ArrayOfStr is populated
}
```

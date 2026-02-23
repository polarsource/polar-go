# BenefitGrantsListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
benefitGrantsListQueryParamCustomerIDFilter := operations.CreateBenefitGrantsListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
benefitGrantsListQueryParamCustomerIDFilter := operations.CreateBenefitGrantsListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitGrantsListQueryParamCustomerIDFilter.Type {
	case operations.BenefitGrantsListQueryParamCustomerIDFilterTypeStr:
		// benefitGrantsListQueryParamCustomerIDFilter.Str is populated
	case operations.BenefitGrantsListQueryParamCustomerIDFilterTypeArrayOfStr:
		// benefitGrantsListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

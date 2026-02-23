# CustomerPortalDownloadablesListQueryParamBenefitIDFilter

Filter by benefit ID.


## Supported Types

### 

```go
customerPortalDownloadablesListQueryParamBenefitIDFilter := operations.CreateCustomerPortalDownloadablesListQueryParamBenefitIDFilterStr(string{/* values here */})
```

### 

```go
customerPortalDownloadablesListQueryParamBenefitIDFilter := operations.CreateCustomerPortalDownloadablesListQueryParamBenefitIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalDownloadablesListQueryParamBenefitIDFilter.Type {
	case operations.CustomerPortalDownloadablesListQueryParamBenefitIDFilterTypeStr:
		// customerPortalDownloadablesListQueryParamBenefitIDFilter.Str is populated
	case operations.CustomerPortalDownloadablesListQueryParamBenefitIDFilterTypeArrayOfStr:
		// customerPortalDownloadablesListQueryParamBenefitIDFilter.ArrayOfStr is populated
}
```

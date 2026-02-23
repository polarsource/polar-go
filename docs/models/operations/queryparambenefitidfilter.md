# QueryParamBenefitIDFilter

Filter by benefit ID.


## Supported Types

### 

```go
queryParamBenefitIDFilter := operations.CreateQueryParamBenefitIDFilterStr(string{/* values here */})
```

### 

```go
queryParamBenefitIDFilter := operations.CreateQueryParamBenefitIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamBenefitIDFilter.Type {
	case operations.QueryParamBenefitIDFilterTypeStr:
		// queryParamBenefitIDFilter.Str is populated
	case operations.QueryParamBenefitIDFilterTypeArrayOfStr:
		// queryParamBenefitIDFilter.ArrayOfStr is populated
}
```

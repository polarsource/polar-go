# QueryParamBenefitTypeFilter

Filter by benefit type.


## Supported Types

### BenefitType

```go
queryParamBenefitTypeFilter := operations.CreateQueryParamBenefitTypeFilterBenefitType(components.BenefitType{/* values here */})
```

### 

```go
queryParamBenefitTypeFilter := operations.CreateQueryParamBenefitTypeFilterArrayOfBenefitType([]components.BenefitType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamBenefitTypeFilter.Type {
	case operations.QueryParamBenefitTypeFilterTypeBenefitType:
		// queryParamBenefitTypeFilter.BenefitType is populated
	case operations.QueryParamBenefitTypeFilterTypeArrayOfBenefitType:
		// queryParamBenefitTypeFilter.ArrayOfBenefitType is populated
}
```

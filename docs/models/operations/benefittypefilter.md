# BenefitTypeFilter

Filter by benefit type.


## Supported Types

### BenefitType

```go
benefitTypeFilter := operations.CreateBenefitTypeFilterBenefitType(components.BenefitType{/* values here */})
```

### 

```go
benefitTypeFilter := operations.CreateBenefitTypeFilterArrayOfBenefitType([]components.BenefitType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitTypeFilter.Type {
	case operations.BenefitTypeFilterTypeBenefitType:
		// benefitTypeFilter.BenefitType is populated
	case operations.BenefitTypeFilterTypeArrayOfBenefitType:
		// benefitTypeFilter.ArrayOfBenefitType is populated
}
```

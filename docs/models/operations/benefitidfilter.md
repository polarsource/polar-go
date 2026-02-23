# BenefitIDFilter

Filter products granting specific benefit.


## Supported Types

### 

```go
benefitIDFilter := operations.CreateBenefitIDFilterStr(string{/* values here */})
```

### 

```go
benefitIDFilter := operations.CreateBenefitIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitIDFilter.Type {
	case operations.BenefitIDFilterTypeStr:
		// benefitIDFilter.Str is populated
	case operations.BenefitIDFilterTypeArrayOfStr:
		// benefitIDFilter.ArrayOfStr is populated
}
```

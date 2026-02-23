# BenefitGitHubRepositoryCreateMetadata


## Supported Types

### 

```go
benefitGitHubRepositoryCreateMetadata := components.CreateBenefitGitHubRepositoryCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitGitHubRepositoryCreateMetadata := components.CreateBenefitGitHubRepositoryCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitGitHubRepositoryCreateMetadata := components.CreateBenefitGitHubRepositoryCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitGitHubRepositoryCreateMetadata := components.CreateBenefitGitHubRepositoryCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitGitHubRepositoryCreateMetadata.Type {
	case components.BenefitGitHubRepositoryCreateMetadataTypeStr:
		// benefitGitHubRepositoryCreateMetadata.Str is populated
	case components.BenefitGitHubRepositoryCreateMetadataTypeInteger:
		// benefitGitHubRepositoryCreateMetadata.Integer is populated
	case components.BenefitGitHubRepositoryCreateMetadataTypeNumber:
		// benefitGitHubRepositoryCreateMetadata.Number is populated
	case components.BenefitGitHubRepositoryCreateMetadataTypeBoolean:
		// benefitGitHubRepositoryCreateMetadata.Boolean is populated
}
```

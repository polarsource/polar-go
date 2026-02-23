# BenefitGitHubRepositoryUpdateMetadata


## Supported Types

### 

```go
benefitGitHubRepositoryUpdateMetadata := components.CreateBenefitGitHubRepositoryUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitGitHubRepositoryUpdateMetadata := components.CreateBenefitGitHubRepositoryUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitGitHubRepositoryUpdateMetadata := components.CreateBenefitGitHubRepositoryUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitGitHubRepositoryUpdateMetadata := components.CreateBenefitGitHubRepositoryUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitGitHubRepositoryUpdateMetadata.Type {
	case components.BenefitGitHubRepositoryUpdateMetadataTypeStr:
		// benefitGitHubRepositoryUpdateMetadata.Str is populated
	case components.BenefitGitHubRepositoryUpdateMetadataTypeInteger:
		// benefitGitHubRepositoryUpdateMetadata.Integer is populated
	case components.BenefitGitHubRepositoryUpdateMetadataTypeNumber:
		// benefitGitHubRepositoryUpdateMetadata.Number is populated
	case components.BenefitGitHubRepositoryUpdateMetadataTypeBoolean:
		// benefitGitHubRepositoryUpdateMetadata.Boolean is populated
}
```

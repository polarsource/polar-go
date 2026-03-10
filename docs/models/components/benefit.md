# Benefit


## Supported Types

### BenefitCustom

```go
benefit := components.CreateBenefitCustom(components.BenefitCustom{/* values here */})
```

### BenefitDiscord

```go
benefit := components.CreateBenefitDiscord(components.BenefitDiscord{/* values here */})
```

### BenefitDownloadables

```go
benefit := components.CreateBenefitDownloadables(components.BenefitDownloadables{/* values here */})
```

### BenefitFeatureFlag

```go
benefit := components.CreateBenefitFeatureFlag(components.BenefitFeatureFlag{/* values here */})
```

### BenefitGitHubRepository

```go
benefit := components.CreateBenefitGithubRepository(components.BenefitGitHubRepository{/* values here */})
```

### BenefitLicenseKeys

```go
benefit := components.CreateBenefitLicenseKeys(components.BenefitLicenseKeys{/* values here */})
```

### BenefitMeterCredit

```go
benefit := components.CreateBenefitMeterCredit(components.BenefitMeterCredit{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefit.Type {
	case components.BenefitUnionTypeCustom:
		// benefit.BenefitCustom is populated
	case components.BenefitUnionTypeDiscord:
		// benefit.BenefitDiscord is populated
	case components.BenefitUnionTypeDownloadables:
		// benefit.BenefitDownloadables is populated
	case components.BenefitUnionTypeFeatureFlag:
		// benefit.BenefitFeatureFlag is populated
	case components.BenefitUnionTypeGithubRepository:
		// benefit.BenefitGitHubRepository is populated
	case components.BenefitUnionTypeLicenseKeys:
		// benefit.BenefitLicenseKeys is populated
	case components.BenefitUnionTypeMeterCredit:
		// benefit.BenefitMeterCredit is populated
}
```

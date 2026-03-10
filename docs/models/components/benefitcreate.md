# BenefitCreate


## Supported Types

### BenefitCustomCreate

```go
benefitCreate := components.CreateBenefitCreateCustom(components.BenefitCustomCreate{/* values here */})
```

### BenefitDiscordCreate

```go
benefitCreate := components.CreateBenefitCreateDiscord(components.BenefitDiscordCreate{/* values here */})
```

### BenefitDownloadablesCreate

```go
benefitCreate := components.CreateBenefitCreateDownloadables(components.BenefitDownloadablesCreate{/* values here */})
```

### BenefitFeatureFlagCreate

```go
benefitCreate := components.CreateBenefitCreateFeatureFlag(components.BenefitFeatureFlagCreate{/* values here */})
```

### BenefitGitHubRepositoryCreate

```go
benefitCreate := components.CreateBenefitCreateGithubRepository(components.BenefitGitHubRepositoryCreate{/* values here */})
```

### BenefitLicenseKeysCreate

```go
benefitCreate := components.CreateBenefitCreateLicenseKeys(components.BenefitLicenseKeysCreate{/* values here */})
```

### BenefitMeterCreditCreate

```go
benefitCreate := components.CreateBenefitCreateMeterCredit(components.BenefitMeterCreditCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitCreate.Type {
	case components.BenefitCreateTypeCustom:
		// benefitCreate.BenefitCustomCreate is populated
	case components.BenefitCreateTypeDiscord:
		// benefitCreate.BenefitDiscordCreate is populated
	case components.BenefitCreateTypeDownloadables:
		// benefitCreate.BenefitDownloadablesCreate is populated
	case components.BenefitCreateTypeFeatureFlag:
		// benefitCreate.BenefitFeatureFlagCreate is populated
	case components.BenefitCreateTypeGithubRepository:
		// benefitCreate.BenefitGitHubRepositoryCreate is populated
	case components.BenefitCreateTypeLicenseKeys:
		// benefitCreate.BenefitLicenseKeysCreate is populated
	case components.BenefitCreateTypeMeterCredit:
		// benefitCreate.BenefitMeterCreditCreate is populated
}
```

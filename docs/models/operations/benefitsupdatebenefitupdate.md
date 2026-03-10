# BenefitsUpdateBenefitUpdate


## Supported Types

### BenefitCustomUpdate

```go
benefitsUpdateBenefitUpdate := operations.CreateBenefitsUpdateBenefitUpdateBenefitCustomUpdate(components.BenefitCustomUpdate{/* values here */})
```

### BenefitDiscordUpdate

```go
benefitsUpdateBenefitUpdate := operations.CreateBenefitsUpdateBenefitUpdateBenefitDiscordUpdate(components.BenefitDiscordUpdate{/* values here */})
```

### BenefitGitHubRepositoryUpdate

```go
benefitsUpdateBenefitUpdate := operations.CreateBenefitsUpdateBenefitUpdateBenefitGitHubRepositoryUpdate(components.BenefitGitHubRepositoryUpdate{/* values here */})
```

### BenefitDownloadablesUpdate

```go
benefitsUpdateBenefitUpdate := operations.CreateBenefitsUpdateBenefitUpdateBenefitDownloadablesUpdate(components.BenefitDownloadablesUpdate{/* values here */})
```

### BenefitLicenseKeysUpdate

```go
benefitsUpdateBenefitUpdate := operations.CreateBenefitsUpdateBenefitUpdateBenefitLicenseKeysUpdate(components.BenefitLicenseKeysUpdate{/* values here */})
```

### BenefitMeterCreditUpdate

```go
benefitsUpdateBenefitUpdate := operations.CreateBenefitsUpdateBenefitUpdateBenefitMeterCreditUpdate(components.BenefitMeterCreditUpdate{/* values here */})
```

### BenefitFeatureFlagUpdate

```go
benefitsUpdateBenefitUpdate := operations.CreateBenefitsUpdateBenefitUpdateBenefitFeatureFlagUpdate(components.BenefitFeatureFlagUpdate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitsUpdateBenefitUpdate.Type {
	case operations.BenefitsUpdateBenefitUpdateTypeBenefitCustomUpdate:
		// benefitsUpdateBenefitUpdate.BenefitCustomUpdate is populated
	case operations.BenefitsUpdateBenefitUpdateTypeBenefitDiscordUpdate:
		// benefitsUpdateBenefitUpdate.BenefitDiscordUpdate is populated
	case operations.BenefitsUpdateBenefitUpdateTypeBenefitGitHubRepositoryUpdate:
		// benefitsUpdateBenefitUpdate.BenefitGitHubRepositoryUpdate is populated
	case operations.BenefitsUpdateBenefitUpdateTypeBenefitDownloadablesUpdate:
		// benefitsUpdateBenefitUpdate.BenefitDownloadablesUpdate is populated
	case operations.BenefitsUpdateBenefitUpdateTypeBenefitLicenseKeysUpdate:
		// benefitsUpdateBenefitUpdate.BenefitLicenseKeysUpdate is populated
	case operations.BenefitsUpdateBenefitUpdateTypeBenefitMeterCreditUpdate:
		// benefitsUpdateBenefitUpdate.BenefitMeterCreditUpdate is populated
	case operations.BenefitsUpdateBenefitUpdateTypeBenefitFeatureFlagUpdate:
		// benefitsUpdateBenefitUpdate.BenefitFeatureFlagUpdate is populated
}
```

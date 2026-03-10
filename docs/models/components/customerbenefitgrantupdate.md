# CustomerBenefitGrantUpdate


## Supported Types

### CustomerBenefitGrantCustomUpdate

```go
customerBenefitGrantUpdate := components.CreateCustomerBenefitGrantUpdateCustom(components.CustomerBenefitGrantCustomUpdate{/* values here */})
```

### CustomerBenefitGrantDiscordUpdate

```go
customerBenefitGrantUpdate := components.CreateCustomerBenefitGrantUpdateDiscord(components.CustomerBenefitGrantDiscordUpdate{/* values here */})
```

### CustomerBenefitGrantDownloadablesUpdate

```go
customerBenefitGrantUpdate := components.CreateCustomerBenefitGrantUpdateDownloadables(components.CustomerBenefitGrantDownloadablesUpdate{/* values here */})
```

### CustomerBenefitGrantFeatureFlagUpdate

```go
customerBenefitGrantUpdate := components.CreateCustomerBenefitGrantUpdateFeatureFlag(components.CustomerBenefitGrantFeatureFlagUpdate{/* values here */})
```

### CustomerBenefitGrantGitHubRepositoryUpdate

```go
customerBenefitGrantUpdate := components.CreateCustomerBenefitGrantUpdateGithubRepository(components.CustomerBenefitGrantGitHubRepositoryUpdate{/* values here */})
```

### CustomerBenefitGrantLicenseKeysUpdate

```go
customerBenefitGrantUpdate := components.CreateCustomerBenefitGrantUpdateLicenseKeys(components.CustomerBenefitGrantLicenseKeysUpdate{/* values here */})
```

### CustomerBenefitGrantMeterCreditUpdate

```go
customerBenefitGrantUpdate := components.CreateCustomerBenefitGrantUpdateMeterCredit(components.CustomerBenefitGrantMeterCreditUpdate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerBenefitGrantUpdate.Type {
	case components.CustomerBenefitGrantUpdateTypeCustom:
		// customerBenefitGrantUpdate.CustomerBenefitGrantCustomUpdate is populated
	case components.CustomerBenefitGrantUpdateTypeDiscord:
		// customerBenefitGrantUpdate.CustomerBenefitGrantDiscordUpdate is populated
	case components.CustomerBenefitGrantUpdateTypeDownloadables:
		// customerBenefitGrantUpdate.CustomerBenefitGrantDownloadablesUpdate is populated
	case components.CustomerBenefitGrantUpdateTypeFeatureFlag:
		// customerBenefitGrantUpdate.CustomerBenefitGrantFeatureFlagUpdate is populated
	case components.CustomerBenefitGrantUpdateTypeGithubRepository:
		// customerBenefitGrantUpdate.CustomerBenefitGrantGitHubRepositoryUpdate is populated
	case components.CustomerBenefitGrantUpdateTypeLicenseKeys:
		// customerBenefitGrantUpdate.CustomerBenefitGrantLicenseKeysUpdate is populated
	case components.CustomerBenefitGrantUpdateTypeMeterCredit:
		// customerBenefitGrantUpdate.CustomerBenefitGrantMeterCreditUpdate is populated
}
```

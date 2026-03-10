# BenefitGrantWebhook


## Supported Types

### BenefitGrantDiscordWebhook

```go
benefitGrantWebhook := components.CreateBenefitGrantWebhookBenefitGrantDiscordWebhook(components.BenefitGrantDiscordWebhook{/* values here */})
```

### BenefitGrantCustomWebhook

```go
benefitGrantWebhook := components.CreateBenefitGrantWebhookBenefitGrantCustomWebhook(components.BenefitGrantCustomWebhook{/* values here */})
```

### BenefitGrantGitHubRepositoryWebhook

```go
benefitGrantWebhook := components.CreateBenefitGrantWebhookBenefitGrantGitHubRepositoryWebhook(components.BenefitGrantGitHubRepositoryWebhook{/* values here */})
```

### BenefitGrantDownloadablesWebhook

```go
benefitGrantWebhook := components.CreateBenefitGrantWebhookBenefitGrantDownloadablesWebhook(components.BenefitGrantDownloadablesWebhook{/* values here */})
```

### BenefitGrantLicenseKeysWebhook

```go
benefitGrantWebhook := components.CreateBenefitGrantWebhookBenefitGrantLicenseKeysWebhook(components.BenefitGrantLicenseKeysWebhook{/* values here */})
```

### BenefitGrantMeterCreditWebhook

```go
benefitGrantWebhook := components.CreateBenefitGrantWebhookBenefitGrantMeterCreditWebhook(components.BenefitGrantMeterCreditWebhook{/* values here */})
```

### BenefitGrantFeatureFlagWebhook

```go
benefitGrantWebhook := components.CreateBenefitGrantWebhookBenefitGrantFeatureFlagWebhook(components.BenefitGrantFeatureFlagWebhook{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitGrantWebhook.Type {
	case components.BenefitGrantWebhookTypeBenefitGrantDiscordWebhook:
		// benefitGrantWebhook.BenefitGrantDiscordWebhook is populated
	case components.BenefitGrantWebhookTypeBenefitGrantCustomWebhook:
		// benefitGrantWebhook.BenefitGrantCustomWebhook is populated
	case components.BenefitGrantWebhookTypeBenefitGrantGitHubRepositoryWebhook:
		// benefitGrantWebhook.BenefitGrantGitHubRepositoryWebhook is populated
	case components.BenefitGrantWebhookTypeBenefitGrantDownloadablesWebhook:
		// benefitGrantWebhook.BenefitGrantDownloadablesWebhook is populated
	case components.BenefitGrantWebhookTypeBenefitGrantLicenseKeysWebhook:
		// benefitGrantWebhook.BenefitGrantLicenseKeysWebhook is populated
	case components.BenefitGrantWebhookTypeBenefitGrantMeterCreditWebhook:
		// benefitGrantWebhook.BenefitGrantMeterCreditWebhook is populated
	case components.BenefitGrantWebhookTypeBenefitGrantFeatureFlagWebhook:
		// benefitGrantWebhook.BenefitGrantFeatureFlagWebhook is populated
}
```

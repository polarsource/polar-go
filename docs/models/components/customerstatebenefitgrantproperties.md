# CustomerStateBenefitGrantProperties


## Supported Types

### BenefitGrantDiscordProperties

```go
customerStateBenefitGrantProperties := components.CreateCustomerStateBenefitGrantPropertiesBenefitGrantDiscordProperties(components.BenefitGrantDiscordProperties{/* values here */})
```

### BenefitGrantGitHubRepositoryProperties

```go
customerStateBenefitGrantProperties := components.CreateCustomerStateBenefitGrantPropertiesBenefitGrantGitHubRepositoryProperties(components.BenefitGrantGitHubRepositoryProperties{/* values here */})
```

### BenefitGrantDownloadablesProperties

```go
customerStateBenefitGrantProperties := components.CreateCustomerStateBenefitGrantPropertiesBenefitGrantDownloadablesProperties(components.BenefitGrantDownloadablesProperties{/* values here */})
```

### BenefitGrantLicenseKeysProperties

```go
customerStateBenefitGrantProperties := components.CreateCustomerStateBenefitGrantPropertiesBenefitGrantLicenseKeysProperties(components.BenefitGrantLicenseKeysProperties{/* values here */})
```

### BenefitGrantCustomProperties

```go
customerStateBenefitGrantProperties := components.CreateCustomerStateBenefitGrantPropertiesBenefitGrantCustomProperties(components.BenefitGrantCustomProperties{/* values here */})
```

### BenefitGrantFeatureFlagProperties

```go
customerStateBenefitGrantProperties := components.CreateCustomerStateBenefitGrantPropertiesBenefitGrantFeatureFlagProperties(components.BenefitGrantFeatureFlagProperties{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerStateBenefitGrantProperties.Type {
	case components.CustomerStateBenefitGrantPropertiesTypeBenefitGrantDiscordProperties:
		// customerStateBenefitGrantProperties.BenefitGrantDiscordProperties is populated
	case components.CustomerStateBenefitGrantPropertiesTypeBenefitGrantGitHubRepositoryProperties:
		// customerStateBenefitGrantProperties.BenefitGrantGitHubRepositoryProperties is populated
	case components.CustomerStateBenefitGrantPropertiesTypeBenefitGrantDownloadablesProperties:
		// customerStateBenefitGrantProperties.BenefitGrantDownloadablesProperties is populated
	case components.CustomerStateBenefitGrantPropertiesTypeBenefitGrantLicenseKeysProperties:
		// customerStateBenefitGrantProperties.BenefitGrantLicenseKeysProperties is populated
	case components.CustomerStateBenefitGrantPropertiesTypeBenefitGrantCustomProperties:
		// customerStateBenefitGrantProperties.BenefitGrantCustomProperties is populated
	case components.CustomerStateBenefitGrantPropertiesTypeBenefitGrantFeatureFlagProperties:
		// customerStateBenefitGrantProperties.BenefitGrantFeatureFlagProperties is populated
}
```

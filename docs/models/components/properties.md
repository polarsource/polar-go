# Properties


## Supported Types

### BenefitGrantDiscordProperties

```go
properties := components.CreatePropertiesBenefitGrantDiscordProperties(components.BenefitGrantDiscordProperties{/* values here */})
```

### BenefitGrantGitHubRepositoryProperties

```go
properties := components.CreatePropertiesBenefitGrantGitHubRepositoryProperties(components.BenefitGrantGitHubRepositoryProperties{/* values here */})
```

### BenefitGrantDownloadablesProperties

```go
properties := components.CreatePropertiesBenefitGrantDownloadablesProperties(components.BenefitGrantDownloadablesProperties{/* values here */})
```

### BenefitGrantLicenseKeysProperties

```go
properties := components.CreatePropertiesBenefitGrantLicenseKeysProperties(components.BenefitGrantLicenseKeysProperties{/* values here */})
```

### BenefitGrantCustomProperties

```go
properties := components.CreatePropertiesBenefitGrantCustomProperties(components.BenefitGrantCustomProperties{/* values here */})
```

### BenefitGrantFeatureFlagProperties

```go
properties := components.CreatePropertiesBenefitGrantFeatureFlagProperties(components.BenefitGrantFeatureFlagProperties{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch properties.Type {
	case components.PropertiesTypeBenefitGrantDiscordProperties:
		// properties.BenefitGrantDiscordProperties is populated
	case components.PropertiesTypeBenefitGrantGitHubRepositoryProperties:
		// properties.BenefitGrantGitHubRepositoryProperties is populated
	case components.PropertiesTypeBenefitGrantDownloadablesProperties:
		// properties.BenefitGrantDownloadablesProperties is populated
	case components.PropertiesTypeBenefitGrantLicenseKeysProperties:
		// properties.BenefitGrantLicenseKeysProperties is populated
	case components.PropertiesTypeBenefitGrantCustomProperties:
		// properties.BenefitGrantCustomProperties is populated
	case components.PropertiesTypeBenefitGrantFeatureFlagProperties:
		// properties.BenefitGrantFeatureFlagProperties is populated
}
```

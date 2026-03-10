# CustomerBenefitGrant


## Supported Types

### CustomerBenefitGrantDiscord

```go
customerBenefitGrant := components.CreateCustomerBenefitGrantCustomerBenefitGrantDiscord(components.CustomerBenefitGrantDiscord{/* values here */})
```

### CustomerBenefitGrantGitHubRepository

```go
customerBenefitGrant := components.CreateCustomerBenefitGrantCustomerBenefitGrantGitHubRepository(components.CustomerBenefitGrantGitHubRepository{/* values here */})
```

### CustomerBenefitGrantDownloadables

```go
customerBenefitGrant := components.CreateCustomerBenefitGrantCustomerBenefitGrantDownloadables(components.CustomerBenefitGrantDownloadables{/* values here */})
```

### CustomerBenefitGrantLicenseKeys

```go
customerBenefitGrant := components.CreateCustomerBenefitGrantCustomerBenefitGrantLicenseKeys(components.CustomerBenefitGrantLicenseKeys{/* values here */})
```

### CustomerBenefitGrantCustom

```go
customerBenefitGrant := components.CreateCustomerBenefitGrantCustomerBenefitGrantCustom(components.CustomerBenefitGrantCustom{/* values here */})
```

### CustomerBenefitGrantMeterCredit

```go
customerBenefitGrant := components.CreateCustomerBenefitGrantCustomerBenefitGrantMeterCredit(components.CustomerBenefitGrantMeterCredit{/* values here */})
```

### CustomerBenefitGrantFeatureFlag

```go
customerBenefitGrant := components.CreateCustomerBenefitGrantCustomerBenefitGrantFeatureFlag(components.CustomerBenefitGrantFeatureFlag{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerBenefitGrant.Type {
	case components.CustomerBenefitGrantTypeCustomerBenefitGrantDiscord:
		// customerBenefitGrant.CustomerBenefitGrantDiscord is populated
	case components.CustomerBenefitGrantTypeCustomerBenefitGrantGitHubRepository:
		// customerBenefitGrant.CustomerBenefitGrantGitHubRepository is populated
	case components.CustomerBenefitGrantTypeCustomerBenefitGrantDownloadables:
		// customerBenefitGrant.CustomerBenefitGrantDownloadables is populated
	case components.CustomerBenefitGrantTypeCustomerBenefitGrantLicenseKeys:
		// customerBenefitGrant.CustomerBenefitGrantLicenseKeys is populated
	case components.CustomerBenefitGrantTypeCustomerBenefitGrantCustom:
		// customerBenefitGrant.CustomerBenefitGrantCustom is populated
	case components.CustomerBenefitGrantTypeCustomerBenefitGrantMeterCredit:
		// customerBenefitGrant.CustomerBenefitGrantMeterCredit is populated
	case components.CustomerBenefitGrantTypeCustomerBenefitGrantFeatureFlag:
		// customerBenefitGrant.CustomerBenefitGrantFeatureFlag is populated
}
```

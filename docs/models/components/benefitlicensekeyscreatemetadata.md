# BenefitLicenseKeysCreateMetadata


## Supported Types

### 

```go
benefitLicenseKeysCreateMetadata := components.CreateBenefitLicenseKeysCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitLicenseKeysCreateMetadata := components.CreateBenefitLicenseKeysCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitLicenseKeysCreateMetadata := components.CreateBenefitLicenseKeysCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitLicenseKeysCreateMetadata := components.CreateBenefitLicenseKeysCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitLicenseKeysCreateMetadata.Type {
	case components.BenefitLicenseKeysCreateMetadataTypeStr:
		// benefitLicenseKeysCreateMetadata.Str is populated
	case components.BenefitLicenseKeysCreateMetadataTypeInteger:
		// benefitLicenseKeysCreateMetadata.Integer is populated
	case components.BenefitLicenseKeysCreateMetadataTypeNumber:
		// benefitLicenseKeysCreateMetadata.Number is populated
	case components.BenefitLicenseKeysCreateMetadataTypeBoolean:
		// benefitLicenseKeysCreateMetadata.Boolean is populated
}
```

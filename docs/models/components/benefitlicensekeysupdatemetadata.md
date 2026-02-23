# BenefitLicenseKeysUpdateMetadata


## Supported Types

### 

```go
benefitLicenseKeysUpdateMetadata := components.CreateBenefitLicenseKeysUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitLicenseKeysUpdateMetadata := components.CreateBenefitLicenseKeysUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitLicenseKeysUpdateMetadata := components.CreateBenefitLicenseKeysUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitLicenseKeysUpdateMetadata := components.CreateBenefitLicenseKeysUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitLicenseKeysUpdateMetadata.Type {
	case components.BenefitLicenseKeysUpdateMetadataTypeStr:
		// benefitLicenseKeysUpdateMetadata.Str is populated
	case components.BenefitLicenseKeysUpdateMetadataTypeInteger:
		// benefitLicenseKeysUpdateMetadata.Integer is populated
	case components.BenefitLicenseKeysUpdateMetadataTypeNumber:
		// benefitLicenseKeysUpdateMetadata.Number is populated
	case components.BenefitLicenseKeysUpdateMetadataTypeBoolean:
		// benefitLicenseKeysUpdateMetadata.Boolean is populated
}
```

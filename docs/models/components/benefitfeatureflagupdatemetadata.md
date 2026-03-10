# BenefitFeatureFlagUpdateMetadata


## Supported Types

### 

```go
benefitFeatureFlagUpdateMetadata := components.CreateBenefitFeatureFlagUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitFeatureFlagUpdateMetadata := components.CreateBenefitFeatureFlagUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitFeatureFlagUpdateMetadata := components.CreateBenefitFeatureFlagUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitFeatureFlagUpdateMetadata := components.CreateBenefitFeatureFlagUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitFeatureFlagUpdateMetadata.Type {
	case components.BenefitFeatureFlagUpdateMetadataTypeStr:
		// benefitFeatureFlagUpdateMetadata.Str is populated
	case components.BenefitFeatureFlagUpdateMetadataTypeInteger:
		// benefitFeatureFlagUpdateMetadata.Integer is populated
	case components.BenefitFeatureFlagUpdateMetadataTypeNumber:
		// benefitFeatureFlagUpdateMetadata.Number is populated
	case components.BenefitFeatureFlagUpdateMetadataTypeBoolean:
		// benefitFeatureFlagUpdateMetadata.Boolean is populated
}
```

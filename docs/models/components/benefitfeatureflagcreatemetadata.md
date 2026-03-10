# BenefitFeatureFlagCreateMetadata


## Supported Types

### 

```go
benefitFeatureFlagCreateMetadata := components.CreateBenefitFeatureFlagCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitFeatureFlagCreateMetadata := components.CreateBenefitFeatureFlagCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitFeatureFlagCreateMetadata := components.CreateBenefitFeatureFlagCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitFeatureFlagCreateMetadata := components.CreateBenefitFeatureFlagCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitFeatureFlagCreateMetadata.Type {
	case components.BenefitFeatureFlagCreateMetadataTypeStr:
		// benefitFeatureFlagCreateMetadata.Str is populated
	case components.BenefitFeatureFlagCreateMetadataTypeInteger:
		// benefitFeatureFlagCreateMetadata.Integer is populated
	case components.BenefitFeatureFlagCreateMetadataTypeNumber:
		// benefitFeatureFlagCreateMetadata.Number is populated
	case components.BenefitFeatureFlagCreateMetadataTypeBoolean:
		// benefitFeatureFlagCreateMetadata.Boolean is populated
}
```

# BenefitCustomUpdateMetadata


## Supported Types

### 

```go
benefitCustomUpdateMetadata := components.CreateBenefitCustomUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitCustomUpdateMetadata := components.CreateBenefitCustomUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitCustomUpdateMetadata := components.CreateBenefitCustomUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitCustomUpdateMetadata := components.CreateBenefitCustomUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitCustomUpdateMetadata.Type {
	case components.BenefitCustomUpdateMetadataTypeStr:
		// benefitCustomUpdateMetadata.Str is populated
	case components.BenefitCustomUpdateMetadataTypeInteger:
		// benefitCustomUpdateMetadata.Integer is populated
	case components.BenefitCustomUpdateMetadataTypeNumber:
		// benefitCustomUpdateMetadata.Number is populated
	case components.BenefitCustomUpdateMetadataTypeBoolean:
		// benefitCustomUpdateMetadata.Boolean is populated
}
```

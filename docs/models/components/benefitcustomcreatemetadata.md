# BenefitCustomCreateMetadata


## Supported Types

### 

```go
benefitCustomCreateMetadata := components.CreateBenefitCustomCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitCustomCreateMetadata := components.CreateBenefitCustomCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitCustomCreateMetadata := components.CreateBenefitCustomCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitCustomCreateMetadata := components.CreateBenefitCustomCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitCustomCreateMetadata.Type {
	case components.BenefitCustomCreateMetadataTypeStr:
		// benefitCustomCreateMetadata.Str is populated
	case components.BenefitCustomCreateMetadataTypeInteger:
		// benefitCustomCreateMetadata.Integer is populated
	case components.BenefitCustomCreateMetadataTypeNumber:
		// benefitCustomCreateMetadata.Number is populated
	case components.BenefitCustomCreateMetadataTypeBoolean:
		// benefitCustomCreateMetadata.Boolean is populated
}
```

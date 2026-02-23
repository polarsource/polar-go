# BenefitMeterCreditUpdateMetadata


## Supported Types

### 

```go
benefitMeterCreditUpdateMetadata := components.CreateBenefitMeterCreditUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitMeterCreditUpdateMetadata := components.CreateBenefitMeterCreditUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitMeterCreditUpdateMetadata := components.CreateBenefitMeterCreditUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitMeterCreditUpdateMetadata := components.CreateBenefitMeterCreditUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitMeterCreditUpdateMetadata.Type {
	case components.BenefitMeterCreditUpdateMetadataTypeStr:
		// benefitMeterCreditUpdateMetadata.Str is populated
	case components.BenefitMeterCreditUpdateMetadataTypeInteger:
		// benefitMeterCreditUpdateMetadata.Integer is populated
	case components.BenefitMeterCreditUpdateMetadataTypeNumber:
		// benefitMeterCreditUpdateMetadata.Number is populated
	case components.BenefitMeterCreditUpdateMetadataTypeBoolean:
		// benefitMeterCreditUpdateMetadata.Boolean is populated
}
```

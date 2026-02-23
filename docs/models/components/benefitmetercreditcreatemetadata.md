# BenefitMeterCreditCreateMetadata


## Supported Types

### 

```go
benefitMeterCreditCreateMetadata := components.CreateBenefitMeterCreditCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitMeterCreditCreateMetadata := components.CreateBenefitMeterCreditCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitMeterCreditCreateMetadata := components.CreateBenefitMeterCreditCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitMeterCreditCreateMetadata := components.CreateBenefitMeterCreditCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitMeterCreditCreateMetadata.Type {
	case components.BenefitMeterCreditCreateMetadataTypeStr:
		// benefitMeterCreditCreateMetadata.Str is populated
	case components.BenefitMeterCreditCreateMetadataTypeInteger:
		// benefitMeterCreditCreateMetadata.Integer is populated
	case components.BenefitMeterCreditCreateMetadataTypeNumber:
		// benefitMeterCreditCreateMetadata.Number is populated
	case components.BenefitMeterCreditCreateMetadataTypeBoolean:
		// benefitMeterCreditCreateMetadata.Boolean is populated
}
```

# BenefitDownloadablesCreateMetadata


## Supported Types

### 

```go
benefitDownloadablesCreateMetadata := components.CreateBenefitDownloadablesCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitDownloadablesCreateMetadata := components.CreateBenefitDownloadablesCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitDownloadablesCreateMetadata := components.CreateBenefitDownloadablesCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitDownloadablesCreateMetadata := components.CreateBenefitDownloadablesCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitDownloadablesCreateMetadata.Type {
	case components.BenefitDownloadablesCreateMetadataTypeStr:
		// benefitDownloadablesCreateMetadata.Str is populated
	case components.BenefitDownloadablesCreateMetadataTypeInteger:
		// benefitDownloadablesCreateMetadata.Integer is populated
	case components.BenefitDownloadablesCreateMetadataTypeNumber:
		// benefitDownloadablesCreateMetadata.Number is populated
	case components.BenefitDownloadablesCreateMetadataTypeBoolean:
		// benefitDownloadablesCreateMetadata.Boolean is populated
}
```

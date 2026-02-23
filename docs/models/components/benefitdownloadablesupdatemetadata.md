# BenefitDownloadablesUpdateMetadata


## Supported Types

### 

```go
benefitDownloadablesUpdateMetadata := components.CreateBenefitDownloadablesUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitDownloadablesUpdateMetadata := components.CreateBenefitDownloadablesUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitDownloadablesUpdateMetadata := components.CreateBenefitDownloadablesUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitDownloadablesUpdateMetadata := components.CreateBenefitDownloadablesUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitDownloadablesUpdateMetadata.Type {
	case components.BenefitDownloadablesUpdateMetadataTypeStr:
		// benefitDownloadablesUpdateMetadata.Str is populated
	case components.BenefitDownloadablesUpdateMetadataTypeInteger:
		// benefitDownloadablesUpdateMetadata.Integer is populated
	case components.BenefitDownloadablesUpdateMetadataTypeNumber:
		// benefitDownloadablesUpdateMetadata.Number is populated
	case components.BenefitDownloadablesUpdateMetadataTypeBoolean:
		// benefitDownloadablesUpdateMetadata.Boolean is populated
}
```

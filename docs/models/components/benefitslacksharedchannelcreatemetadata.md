# BenefitSlackSharedChannelCreateMetadata


## Supported Types

### 

```go
benefitSlackSharedChannelCreateMetadata := components.CreateBenefitSlackSharedChannelCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitSlackSharedChannelCreateMetadata := components.CreateBenefitSlackSharedChannelCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitSlackSharedChannelCreateMetadata := components.CreateBenefitSlackSharedChannelCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitSlackSharedChannelCreateMetadata := components.CreateBenefitSlackSharedChannelCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitSlackSharedChannelCreateMetadata.Type {
	case components.BenefitSlackSharedChannelCreateMetadataTypeStr:
		// benefitSlackSharedChannelCreateMetadata.Str is populated
	case components.BenefitSlackSharedChannelCreateMetadataTypeInteger:
		// benefitSlackSharedChannelCreateMetadata.Integer is populated
	case components.BenefitSlackSharedChannelCreateMetadataTypeNumber:
		// benefitSlackSharedChannelCreateMetadata.Number is populated
	case components.BenefitSlackSharedChannelCreateMetadataTypeBoolean:
		// benefitSlackSharedChannelCreateMetadata.Boolean is populated
}
```

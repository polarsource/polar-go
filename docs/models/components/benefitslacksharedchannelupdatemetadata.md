# BenefitSlackSharedChannelUpdateMetadata


## Supported Types

### 

```go
benefitSlackSharedChannelUpdateMetadata := components.CreateBenefitSlackSharedChannelUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitSlackSharedChannelUpdateMetadata := components.CreateBenefitSlackSharedChannelUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitSlackSharedChannelUpdateMetadata := components.CreateBenefitSlackSharedChannelUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitSlackSharedChannelUpdateMetadata := components.CreateBenefitSlackSharedChannelUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitSlackSharedChannelUpdateMetadata.Type {
	case components.BenefitSlackSharedChannelUpdateMetadataTypeStr:
		// benefitSlackSharedChannelUpdateMetadata.Str is populated
	case components.BenefitSlackSharedChannelUpdateMetadataTypeInteger:
		// benefitSlackSharedChannelUpdateMetadata.Integer is populated
	case components.BenefitSlackSharedChannelUpdateMetadataTypeNumber:
		// benefitSlackSharedChannelUpdateMetadata.Number is populated
	case components.BenefitSlackSharedChannelUpdateMetadataTypeBoolean:
		// benefitSlackSharedChannelUpdateMetadata.Boolean is populated
}
```

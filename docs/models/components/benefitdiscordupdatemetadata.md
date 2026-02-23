# BenefitDiscordUpdateMetadata


## Supported Types

### 

```go
benefitDiscordUpdateMetadata := components.CreateBenefitDiscordUpdateMetadataStr(string{/* values here */})
```

### 

```go
benefitDiscordUpdateMetadata := components.CreateBenefitDiscordUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitDiscordUpdateMetadata := components.CreateBenefitDiscordUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitDiscordUpdateMetadata := components.CreateBenefitDiscordUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitDiscordUpdateMetadata.Type {
	case components.BenefitDiscordUpdateMetadataTypeStr:
		// benefitDiscordUpdateMetadata.Str is populated
	case components.BenefitDiscordUpdateMetadataTypeInteger:
		// benefitDiscordUpdateMetadata.Integer is populated
	case components.BenefitDiscordUpdateMetadataTypeNumber:
		// benefitDiscordUpdateMetadata.Number is populated
	case components.BenefitDiscordUpdateMetadataTypeBoolean:
		// benefitDiscordUpdateMetadata.Boolean is populated
}
```

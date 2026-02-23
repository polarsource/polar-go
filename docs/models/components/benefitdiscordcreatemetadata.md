# BenefitDiscordCreateMetadata


## Supported Types

### 

```go
benefitDiscordCreateMetadata := components.CreateBenefitDiscordCreateMetadataStr(string{/* values here */})
```

### 

```go
benefitDiscordCreateMetadata := components.CreateBenefitDiscordCreateMetadataInteger(int64{/* values here */})
```

### 

```go
benefitDiscordCreateMetadata := components.CreateBenefitDiscordCreateMetadataNumber(float64{/* values here */})
```

### 

```go
benefitDiscordCreateMetadata := components.CreateBenefitDiscordCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch benefitDiscordCreateMetadata.Type {
	case components.BenefitDiscordCreateMetadataTypeStr:
		// benefitDiscordCreateMetadata.Str is populated
	case components.BenefitDiscordCreateMetadataTypeInteger:
		// benefitDiscordCreateMetadata.Integer is populated
	case components.BenefitDiscordCreateMetadataTypeNumber:
		// benefitDiscordCreateMetadata.Number is populated
	case components.BenefitDiscordCreateMetadataTypeBoolean:
		// benefitDiscordCreateMetadata.Boolean is populated
}
```

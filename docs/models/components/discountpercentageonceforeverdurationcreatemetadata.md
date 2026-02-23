# DiscountPercentageOnceForeverDurationCreateMetadata


## Supported Types

### 

```go
discountPercentageOnceForeverDurationCreateMetadata := components.CreateDiscountPercentageOnceForeverDurationCreateMetadataStr(string{/* values here */})
```

### 

```go
discountPercentageOnceForeverDurationCreateMetadata := components.CreateDiscountPercentageOnceForeverDurationCreateMetadataInteger(int64{/* values here */})
```

### 

```go
discountPercentageOnceForeverDurationCreateMetadata := components.CreateDiscountPercentageOnceForeverDurationCreateMetadataNumber(float64{/* values here */})
```

### 

```go
discountPercentageOnceForeverDurationCreateMetadata := components.CreateDiscountPercentageOnceForeverDurationCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountPercentageOnceForeverDurationCreateMetadata.Type {
	case components.DiscountPercentageOnceForeverDurationCreateMetadataTypeStr:
		// discountPercentageOnceForeverDurationCreateMetadata.Str is populated
	case components.DiscountPercentageOnceForeverDurationCreateMetadataTypeInteger:
		// discountPercentageOnceForeverDurationCreateMetadata.Integer is populated
	case components.DiscountPercentageOnceForeverDurationCreateMetadataTypeNumber:
		// discountPercentageOnceForeverDurationCreateMetadata.Number is populated
	case components.DiscountPercentageOnceForeverDurationCreateMetadataTypeBoolean:
		// discountPercentageOnceForeverDurationCreateMetadata.Boolean is populated
}
```

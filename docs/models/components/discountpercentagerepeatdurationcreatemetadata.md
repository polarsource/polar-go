# DiscountPercentageRepeatDurationCreateMetadata


## Supported Types

### 

```go
discountPercentageRepeatDurationCreateMetadata := components.CreateDiscountPercentageRepeatDurationCreateMetadataStr(string{/* values here */})
```

### 

```go
discountPercentageRepeatDurationCreateMetadata := components.CreateDiscountPercentageRepeatDurationCreateMetadataInteger(int64{/* values here */})
```

### 

```go
discountPercentageRepeatDurationCreateMetadata := components.CreateDiscountPercentageRepeatDurationCreateMetadataNumber(float64{/* values here */})
```

### 

```go
discountPercentageRepeatDurationCreateMetadata := components.CreateDiscountPercentageRepeatDurationCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountPercentageRepeatDurationCreateMetadata.Type {
	case components.DiscountPercentageRepeatDurationCreateMetadataTypeStr:
		// discountPercentageRepeatDurationCreateMetadata.Str is populated
	case components.DiscountPercentageRepeatDurationCreateMetadataTypeInteger:
		// discountPercentageRepeatDurationCreateMetadata.Integer is populated
	case components.DiscountPercentageRepeatDurationCreateMetadataTypeNumber:
		// discountPercentageRepeatDurationCreateMetadata.Number is populated
	case components.DiscountPercentageRepeatDurationCreateMetadataTypeBoolean:
		// discountPercentageRepeatDurationCreateMetadata.Boolean is populated
}
```

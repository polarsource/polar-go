# DiscountFixedRepeatDurationCreateMetadata


## Supported Types

### 

```go
discountFixedRepeatDurationCreateMetadata := components.CreateDiscountFixedRepeatDurationCreateMetadataStr(string{/* values here */})
```

### 

```go
discountFixedRepeatDurationCreateMetadata := components.CreateDiscountFixedRepeatDurationCreateMetadataInteger(int64{/* values here */})
```

### 

```go
discountFixedRepeatDurationCreateMetadata := components.CreateDiscountFixedRepeatDurationCreateMetadataNumber(float64{/* values here */})
```

### 

```go
discountFixedRepeatDurationCreateMetadata := components.CreateDiscountFixedRepeatDurationCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountFixedRepeatDurationCreateMetadata.Type {
	case components.DiscountFixedRepeatDurationCreateMetadataTypeStr:
		// discountFixedRepeatDurationCreateMetadata.Str is populated
	case components.DiscountFixedRepeatDurationCreateMetadataTypeInteger:
		// discountFixedRepeatDurationCreateMetadata.Integer is populated
	case components.DiscountFixedRepeatDurationCreateMetadataTypeNumber:
		// discountFixedRepeatDurationCreateMetadata.Number is populated
	case components.DiscountFixedRepeatDurationCreateMetadataTypeBoolean:
		// discountFixedRepeatDurationCreateMetadata.Boolean is populated
}
```

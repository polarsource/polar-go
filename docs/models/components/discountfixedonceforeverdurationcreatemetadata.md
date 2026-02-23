# DiscountFixedOnceForeverDurationCreateMetadata


## Supported Types

### 

```go
discountFixedOnceForeverDurationCreateMetadata := components.CreateDiscountFixedOnceForeverDurationCreateMetadataStr(string{/* values here */})
```

### 

```go
discountFixedOnceForeverDurationCreateMetadata := components.CreateDiscountFixedOnceForeverDurationCreateMetadataInteger(int64{/* values here */})
```

### 

```go
discountFixedOnceForeverDurationCreateMetadata := components.CreateDiscountFixedOnceForeverDurationCreateMetadataNumber(float64{/* values here */})
```

### 

```go
discountFixedOnceForeverDurationCreateMetadata := components.CreateDiscountFixedOnceForeverDurationCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountFixedOnceForeverDurationCreateMetadata.Type {
	case components.DiscountFixedOnceForeverDurationCreateMetadataTypeStr:
		// discountFixedOnceForeverDurationCreateMetadata.Str is populated
	case components.DiscountFixedOnceForeverDurationCreateMetadataTypeInteger:
		// discountFixedOnceForeverDurationCreateMetadata.Integer is populated
	case components.DiscountFixedOnceForeverDurationCreateMetadataTypeNumber:
		// discountFixedOnceForeverDurationCreateMetadata.Number is populated
	case components.DiscountFixedOnceForeverDurationCreateMetadataTypeBoolean:
		// discountFixedOnceForeverDurationCreateMetadata.Boolean is populated
}
```

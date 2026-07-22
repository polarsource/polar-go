# DiscountPercentageCreateMetadata


## Supported Types

### 

```go
discountPercentageCreateMetadata := components.CreateDiscountPercentageCreateMetadataStr(string{/* values here */})
```

### 

```go
discountPercentageCreateMetadata := components.CreateDiscountPercentageCreateMetadataInteger(int64{/* values here */})
```

### 

```go
discountPercentageCreateMetadata := components.CreateDiscountPercentageCreateMetadataNumber(float64{/* values here */})
```

### 

```go
discountPercentageCreateMetadata := components.CreateDiscountPercentageCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountPercentageCreateMetadata.Type {
	case components.DiscountPercentageCreateMetadataTypeStr:
		// discountPercentageCreateMetadata.Str is populated
	case components.DiscountPercentageCreateMetadataTypeInteger:
		// discountPercentageCreateMetadata.Integer is populated
	case components.DiscountPercentageCreateMetadataTypeNumber:
		// discountPercentageCreateMetadata.Number is populated
	case components.DiscountPercentageCreateMetadataTypeBoolean:
		// discountPercentageCreateMetadata.Boolean is populated
}
```

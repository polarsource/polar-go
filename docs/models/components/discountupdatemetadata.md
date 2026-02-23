# DiscountUpdateMetadata


## Supported Types

### 

```go
discountUpdateMetadata := components.CreateDiscountUpdateMetadataStr(string{/* values here */})
```

### 

```go
discountUpdateMetadata := components.CreateDiscountUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
discountUpdateMetadata := components.CreateDiscountUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
discountUpdateMetadata := components.CreateDiscountUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountUpdateMetadata.Type {
	case components.DiscountUpdateMetadataTypeStr:
		// discountUpdateMetadata.Str is populated
	case components.DiscountUpdateMetadataTypeInteger:
		// discountUpdateMetadata.Integer is populated
	case components.DiscountUpdateMetadataTypeNumber:
		// discountUpdateMetadata.Number is populated
	case components.DiscountUpdateMetadataTypeBoolean:
		// discountUpdateMetadata.Boolean is populated
}
```

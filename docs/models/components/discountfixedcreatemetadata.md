# DiscountFixedCreateMetadata


## Supported Types

### 

```go
discountFixedCreateMetadata := components.CreateDiscountFixedCreateMetadataStr(string{/* values here */})
```

### 

```go
discountFixedCreateMetadata := components.CreateDiscountFixedCreateMetadataInteger(int64{/* values here */})
```

### 

```go
discountFixedCreateMetadata := components.CreateDiscountFixedCreateMetadataNumber(float64{/* values here */})
```

### 

```go
discountFixedCreateMetadata := components.CreateDiscountFixedCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch discountFixedCreateMetadata.Type {
	case components.DiscountFixedCreateMetadataTypeStr:
		// discountFixedCreateMetadata.Str is populated
	case components.DiscountFixedCreateMetadataTypeInteger:
		// discountFixedCreateMetadata.Integer is populated
	case components.DiscountFixedCreateMetadataTypeNumber:
		// discountFixedCreateMetadata.Number is populated
	case components.DiscountFixedCreateMetadataTypeBoolean:
		// discountFixedCreateMetadata.Boolean is populated
}
```

# ProductUpdateMetadata


## Supported Types

### 

```go
productUpdateMetadata := components.CreateProductUpdateMetadataStr(string{/* values here */})
```

### 

```go
productUpdateMetadata := components.CreateProductUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
productUpdateMetadata := components.CreateProductUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
productUpdateMetadata := components.CreateProductUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productUpdateMetadata.Type {
	case components.ProductUpdateMetadataTypeStr:
		// productUpdateMetadata.Str is populated
	case components.ProductUpdateMetadataTypeInteger:
		// productUpdateMetadata.Integer is populated
	case components.ProductUpdateMetadataTypeNumber:
		// productUpdateMetadata.Number is populated
	case components.ProductUpdateMetadataTypeBoolean:
		// productUpdateMetadata.Boolean is populated
}
```

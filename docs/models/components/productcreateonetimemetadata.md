# ProductCreateOneTimeMetadata


## Supported Types

### 

```go
productCreateOneTimeMetadata := components.CreateProductCreateOneTimeMetadataStr(string{/* values here */})
```

### 

```go
productCreateOneTimeMetadata := components.CreateProductCreateOneTimeMetadataInteger(int64{/* values here */})
```

### 

```go
productCreateOneTimeMetadata := components.CreateProductCreateOneTimeMetadataNumber(float64{/* values here */})
```

### 

```go
productCreateOneTimeMetadata := components.CreateProductCreateOneTimeMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productCreateOneTimeMetadata.Type {
	case components.ProductCreateOneTimeMetadataTypeStr:
		// productCreateOneTimeMetadata.Str is populated
	case components.ProductCreateOneTimeMetadataTypeInteger:
		// productCreateOneTimeMetadata.Integer is populated
	case components.ProductCreateOneTimeMetadataTypeNumber:
		// productCreateOneTimeMetadata.Number is populated
	case components.ProductCreateOneTimeMetadataTypeBoolean:
		// productCreateOneTimeMetadata.Boolean is populated
}
```

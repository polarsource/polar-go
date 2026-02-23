# ProductCreateRecurringMetadata


## Supported Types

### 

```go
productCreateRecurringMetadata := components.CreateProductCreateRecurringMetadataStr(string{/* values here */})
```

### 

```go
productCreateRecurringMetadata := components.CreateProductCreateRecurringMetadataInteger(int64{/* values here */})
```

### 

```go
productCreateRecurringMetadata := components.CreateProductCreateRecurringMetadataNumber(float64{/* values here */})
```

### 

```go
productCreateRecurringMetadata := components.CreateProductCreateRecurringMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch productCreateRecurringMetadata.Type {
	case components.ProductCreateRecurringMetadataTypeStr:
		// productCreateRecurringMetadata.Str is populated
	case components.ProductCreateRecurringMetadataTypeInteger:
		// productCreateRecurringMetadata.Integer is populated
	case components.ProductCreateRecurringMetadataTypeNumber:
		// productCreateRecurringMetadata.Number is populated
	case components.ProductCreateRecurringMetadataTypeBoolean:
		// productCreateRecurringMetadata.Boolean is populated
}
```

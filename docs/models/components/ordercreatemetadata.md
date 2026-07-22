# OrderCreateMetadata


## Supported Types

### 

```go
orderCreateMetadata := components.CreateOrderCreateMetadataStr(string{/* values here */})
```

### 

```go
orderCreateMetadata := components.CreateOrderCreateMetadataInteger(int64{/* values here */})
```

### 

```go
orderCreateMetadata := components.CreateOrderCreateMetadataNumber(float64{/* values here */})
```

### 

```go
orderCreateMetadata := components.CreateOrderCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch orderCreateMetadata.Type {
	case components.OrderCreateMetadataTypeStr:
		// orderCreateMetadata.Str is populated
	case components.OrderCreateMetadataTypeInteger:
		// orderCreateMetadata.Integer is populated
	case components.OrderCreateMetadataTypeNumber:
		// orderCreateMetadata.Number is populated
	case components.OrderCreateMetadataTypeBoolean:
		// orderCreateMetadata.Boolean is populated
}
```

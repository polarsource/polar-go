# MetadataOutputType


## Supported Types

### 

```go
metadataOutputType := components.CreateMetadataOutputTypeStr(string{/* values here */})
```

### 

```go
metadataOutputType := components.CreateMetadataOutputTypeInteger(int64{/* values here */})
```

### 

```go
metadataOutputType := components.CreateMetadataOutputTypeNumber(float64{/* values here */})
```

### 

```go
metadataOutputType := components.CreateMetadataOutputTypeBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metadataOutputType.Type {
	case components.MetadataOutputTypeTypeStr:
		// metadataOutputType.Str is populated
	case components.MetadataOutputTypeTypeInteger:
		// metadataOutputType.Integer is populated
	case components.MetadataOutputTypeTypeNumber:
		// metadataOutputType.Number is populated
	case components.MetadataOutputTypeTypeBoolean:
		// metadataOutputType.Boolean is populated
}
```

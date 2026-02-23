# Metadata


## Supported Types

### 

```go
metadata := components.CreateMetadataStr(string{/* values here */})
```

### 

```go
metadata := components.CreateMetadataInteger(int64{/* values here */})
```

### 

```go
metadata := components.CreateMetadataNumber(float64{/* values here */})
```

### 

```go
metadata := components.CreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metadata.Type {
	case components.MetadataTypeStr:
		// metadata.Str is populated
	case components.MetadataTypeInteger:
		// metadata.Integer is populated
	case components.MetadataTypeNumber:
		// metadata.Number is populated
	case components.MetadataTypeBoolean:
		// metadata.Boolean is populated
}
```

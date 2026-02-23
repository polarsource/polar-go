# MetadataQuery


## Supported Types

### 

```go
metadataQuery := components.CreateMetadataQueryStr(string{/* values here */})
```

### 

```go
metadataQuery := components.CreateMetadataQueryInteger(int64{/* values here */})
```

### 

```go
metadataQuery := components.CreateMetadataQueryBoolean(bool{/* values here */})
```

### 

```go
metadataQuery := components.CreateMetadataQueryArrayOfStr([]string{/* values here */})
```

### 

```go
metadataQuery := components.CreateMetadataQueryArrayOfInteger([]int64{/* values here */})
```

### 

```go
metadataQuery := components.CreateMetadataQueryArrayOfBoolean([]bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch metadataQuery.Type {
	case components.MetadataQueryTypeStr:
		// metadataQuery.Str is populated
	case components.MetadataQueryTypeInteger:
		// metadataQuery.Integer is populated
	case components.MetadataQueryTypeBoolean:
		// metadataQuery.Boolean is populated
	case components.MetadataQueryTypeArrayOfStr:
		// metadataQuery.ArrayOfStr is populated
	case components.MetadataQueryTypeArrayOfInteger:
		// metadataQuery.ArrayOfInteger is populated
	case components.MetadataQueryTypeArrayOfBoolean:
		// metadataQuery.ArrayOfBoolean is populated
}
```

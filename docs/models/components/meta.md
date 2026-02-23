# Meta


## Supported Types

### 

```go
meta := components.CreateMetaStr(string{/* values here */})
```

### 

```go
meta := components.CreateMetaInteger(int64{/* values here */})
```

### 

```go
meta := components.CreateMetaNumber(float64{/* values here */})
```

### 

```go
meta := components.CreateMetaBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch meta.Type {
	case components.MetaTypeStr:
		// meta.Str is populated
	case components.MetaTypeInteger:
		// meta.Integer is populated
	case components.MetaTypeNumber:
		// meta.Number is populated
	case components.MetaTypeBoolean:
		// meta.Boolean is populated
}
```

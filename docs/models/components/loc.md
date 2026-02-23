# Loc


## Supported Types

### 

```go
loc := components.CreateLocStr(string{/* values here */})
```

### 

```go
loc := components.CreateLocInteger(int64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch loc.Type {
	case components.LocTypeStr:
		// loc.Str is populated
	case components.LocTypeInteger:
		// loc.Integer is populated
}
```

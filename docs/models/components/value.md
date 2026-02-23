# Value


## Supported Types

### 

```go
value := components.CreateValueStr(string{/* values here */})
```

### 

```go
value := components.CreateValueInteger(int64{/* values here */})
```

### 

```go
value := components.CreateValueBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch value.Type {
	case components.ValueTypeStr:
		// value.Str is populated
	case components.ValueTypeInteger:
		// value.Integer is populated
	case components.ValueTypeBoolean:
		// value.Boolean is populated
}
```

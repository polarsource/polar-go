# Conditions


## Supported Types

### 

```go
conditions := components.CreateConditionsStr(string{/* values here */})
```

### 

```go
conditions := components.CreateConditionsInteger(int64{/* values here */})
```

### 

```go
conditions := components.CreateConditionsNumber(float64{/* values here */})
```

### 

```go
conditions := components.CreateConditionsBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch conditions.Type {
	case components.ConditionsTypeStr:
		// conditions.Str is populated
	case components.ConditionsTypeInteger:
		// conditions.Integer is populated
	case components.ConditionsTypeNumber:
		// conditions.Number is populated
	case components.ConditionsTypeBoolean:
		// conditions.Boolean is populated
}
```

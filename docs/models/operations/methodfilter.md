# MethodFilter

Filter by payment method.


## Supported Types

### 

```go
methodFilter := operations.CreateMethodFilterStr(string{/* values here */})
```

### 

```go
methodFilter := operations.CreateMethodFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch methodFilter.Type {
	case operations.MethodFilterTypeStr:
		// methodFilter.Str is populated
	case operations.MethodFilterTypeArrayOfStr:
		// methodFilter.ArrayOfStr is populated
}
```

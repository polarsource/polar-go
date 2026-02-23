# NameFilter

Filter by event name.


## Supported Types

### 

```go
nameFilter := operations.CreateNameFilterStr(string{/* values here */})
```

### 

```go
nameFilter := operations.CreateNameFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch nameFilter.Type {
	case operations.NameFilterTypeStr:
		// nameFilter.Str is populated
	case operations.NameFilterTypeArrayOfStr:
		// nameFilter.ArrayOfStr is populated
}
```

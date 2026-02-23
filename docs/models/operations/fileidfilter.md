# FileIDFilter

Filter by file ID.


## Supported Types

### 

```go
fileIDFilter := operations.CreateFileIDFilterStr(string{/* values here */})
```

### 

```go
fileIDFilter := operations.CreateFileIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch fileIDFilter.Type {
	case operations.FileIDFilterTypeStr:
		// fileIDFilter.Str is populated
	case operations.FileIDFilterTypeArrayOfStr:
		// fileIDFilter.ArrayOfStr is populated
}
```

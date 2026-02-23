# ExcludeIDs

Exclude benefits with these IDs.


## Supported Types

### 

```go
excludeIDs := operations.CreateExcludeIDsStr(string{/* values here */})
```

### 

```go
excludeIDs := operations.CreateExcludeIDsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch excludeIDs.Type {
	case operations.ExcludeIDsTypeStr:
		// excludeIDs.Str is populated
	case operations.ExcludeIDsTypeArrayOfStr:
		// excludeIDs.ArrayOfStr is populated
}
```

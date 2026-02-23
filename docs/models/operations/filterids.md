# FilterIDs

Filter by benefit IDs.


## Supported Types

### 

```go
filterIDs := operations.CreateFilterIDsStr(string{/* values here */})
```

### 

```go
filterIDs := operations.CreateFilterIDsArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch filterIDs.Type {
	case operations.FilterIDsTypeStr:
		// filterIDs.Str is populated
	case operations.FilterIDsTypeArrayOfStr:
		// filterIDs.ArrayOfStr is populated
}
```

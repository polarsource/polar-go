# MemberIDFilter

Filter by member.


## Supported Types

### 

```go
memberIDFilter := operations.CreateMemberIDFilterStr(string{/* values here */})
```

### 

```go
memberIDFilter := operations.CreateMemberIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch memberIDFilter.Type {
	case operations.MemberIDFilterTypeStr:
		// memberIDFilter.Str is populated
	case operations.MemberIDFilterTypeArrayOfStr:
		// memberIDFilter.ArrayOfStr is populated
}
```

# QueryParamMemberIDFilter

Filter by member ID.


## Supported Types

### 

```go
queryParamMemberIDFilter := operations.CreateQueryParamMemberIDFilterStr(string{/* values here */})
```

### 

```go
queryParamMemberIDFilter := operations.CreateQueryParamMemberIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamMemberIDFilter.Type {
	case operations.QueryParamMemberIDFilterTypeStr:
		// queryParamMemberIDFilter.Str is populated
	case operations.QueryParamMemberIDFilterTypeArrayOfStr:
		// queryParamMemberIDFilter.ArrayOfStr is populated
}
```

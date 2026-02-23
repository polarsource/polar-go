# QueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
queryParamProductIDFilter := operations.CreateQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
queryParamProductIDFilter := operations.CreateQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamProductIDFilter.Type {
	case operations.QueryParamProductIDFilterTypeStr:
		// queryParamProductIDFilter.Str is populated
	case operations.QueryParamProductIDFilterTypeArrayOfStr:
		// queryParamProductIDFilter.ArrayOfStr is populated
}
```

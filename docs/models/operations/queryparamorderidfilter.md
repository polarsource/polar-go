# QueryParamOrderIDFilter

Filter by order ID.


## Supported Types

### 

```go
queryParamOrderIDFilter := operations.CreateQueryParamOrderIDFilterStr(string{/* values here */})
```

### 

```go
queryParamOrderIDFilter := operations.CreateQueryParamOrderIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamOrderIDFilter.Type {
	case operations.QueryParamOrderIDFilterTypeStr:
		// queryParamOrderIDFilter.Str is populated
	case operations.QueryParamOrderIDFilterTypeArrayOfStr:
		// queryParamOrderIDFilter.ArrayOfStr is populated
}
```

# QueryParamCheckoutIDFilter

Filter by checkout ID.


## Supported Types

### 

```go
queryParamCheckoutIDFilter := operations.CreateQueryParamCheckoutIDFilterStr(string{/* values here */})
```

### 

```go
queryParamCheckoutIDFilter := operations.CreateQueryParamCheckoutIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamCheckoutIDFilter.Type {
	case operations.QueryParamCheckoutIDFilterTypeStr:
		// queryParamCheckoutIDFilter.Str is populated
	case operations.QueryParamCheckoutIDFilterTypeArrayOfStr:
		// queryParamCheckoutIDFilter.ArrayOfStr is populated
}
```

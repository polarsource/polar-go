# QueryParamSubscriptionIDFilter

Filter by subscription ID.


## Supported Types

### 

```go
queryParamSubscriptionIDFilter := operations.CreateQueryParamSubscriptionIDFilterStr(string{/* values here */})
```

### 

```go
queryParamSubscriptionIDFilter := operations.CreateQueryParamSubscriptionIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamSubscriptionIDFilter.Type {
	case operations.QueryParamSubscriptionIDFilterTypeStr:
		// queryParamSubscriptionIDFilter.Str is populated
	case operations.QueryParamSubscriptionIDFilterTypeArrayOfStr:
		// queryParamSubscriptionIDFilter.ArrayOfStr is populated
}
```

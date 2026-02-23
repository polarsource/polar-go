# CheckoutsListQueryParamProductIDFilter

Filter by product ID.


## Supported Types

### 

```go
checkoutsListQueryParamProductIDFilter := operations.CreateCheckoutsListQueryParamProductIDFilterStr(string{/* values here */})
```

### 

```go
checkoutsListQueryParamProductIDFilter := operations.CreateCheckoutsListQueryParamProductIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutsListQueryParamProductIDFilter.Type {
	case operations.CheckoutsListQueryParamProductIDFilterTypeStr:
		// checkoutsListQueryParamProductIDFilter.Str is populated
	case operations.CheckoutsListQueryParamProductIDFilterTypeArrayOfStr:
		// checkoutsListQueryParamProductIDFilter.ArrayOfStr is populated
}
```

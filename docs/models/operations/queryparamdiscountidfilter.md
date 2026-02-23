# QueryParamDiscountIDFilter

Filter by discount ID.


## Supported Types

### 

```go
queryParamDiscountIDFilter := operations.CreateQueryParamDiscountIDFilterStr(string{/* values here */})
```

### 

```go
queryParamDiscountIDFilter := operations.CreateQueryParamDiscountIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamDiscountIDFilter.Type {
	case operations.QueryParamDiscountIDFilterTypeStr:
		// queryParamDiscountIDFilter.Str is populated
	case operations.QueryParamDiscountIDFilterTypeArrayOfStr:
		// queryParamDiscountIDFilter.ArrayOfStr is populated
}
```

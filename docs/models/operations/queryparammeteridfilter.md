# QueryParamMeterIDFilter

Filter by meter ID.


## Supported Types

### 

```go
queryParamMeterIDFilter := operations.CreateQueryParamMeterIDFilterStr(string{/* values here */})
```

### 

```go
queryParamMeterIDFilter := operations.CreateQueryParamMeterIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamMeterIDFilter.Type {
	case operations.QueryParamMeterIDFilterTypeStr:
		// queryParamMeterIDFilter.Str is populated
	case operations.QueryParamMeterIDFilterTypeArrayOfStr:
		// queryParamMeterIDFilter.ArrayOfStr is populated
}
```

# QueryParamSourceFilter

Filter by event source.


## Supported Types

### EventSource

```go
queryParamSourceFilter := operations.CreateQueryParamSourceFilterEventSource(components.EventSource{/* values here */})
```

### 

```go
queryParamSourceFilter := operations.CreateQueryParamSourceFilterArrayOfEventSource([]components.EventSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamSourceFilter.Type {
	case operations.QueryParamSourceFilterTypeEventSource:
		// queryParamSourceFilter.EventSource is populated
	case operations.QueryParamSourceFilterTypeArrayOfEventSource:
		// queryParamSourceFilter.ArrayOfEventSource is populated
}
```

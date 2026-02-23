# SourceFilter

Filter by event source.


## Supported Types

### EventSource

```go
sourceFilter := operations.CreateSourceFilterEventSource(components.EventSource{/* values here */})
```

### 

```go
sourceFilter := operations.CreateSourceFilterArrayOfEventSource([]components.EventSource{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch sourceFilter.Type {
	case operations.SourceFilterTypeEventSource:
		// sourceFilter.EventSource is populated
	case operations.SourceFilterTypeArrayOfEventSource:
		// sourceFilter.ArrayOfEventSource is populated
}
```

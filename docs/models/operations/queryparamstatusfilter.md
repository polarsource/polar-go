# QueryParamStatusFilter

Filter by dispute status.


## Supported Types

### DisputeStatus

```go
queryParamStatusFilter := operations.CreateQueryParamStatusFilterDisputeStatus(components.DisputeStatus{/* values here */})
```

### 

```go
queryParamStatusFilter := operations.CreateQueryParamStatusFilterArrayOfDisputeStatus([]components.DisputeStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch queryParamStatusFilter.Type {
	case operations.QueryParamStatusFilterTypeDisputeStatus:
		// queryParamStatusFilter.DisputeStatus is populated
	case operations.QueryParamStatusFilterTypeArrayOfDisputeStatus:
		// queryParamStatusFilter.ArrayOfDisputeStatus is populated
}
```

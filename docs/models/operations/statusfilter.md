# StatusFilter

Filter by dispute status.


## Supported Types

### DisputeStatus

```go
statusFilter := operations.CreateStatusFilterDisputeStatus(components.DisputeStatus{/* values here */})
```

### 

```go
statusFilter := operations.CreateStatusFilterArrayOfDisputeStatus([]components.DisputeStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch statusFilter.Type {
	case operations.StatusFilterTypeDisputeStatus:
		// statusFilter.DisputeStatus is populated
	case operations.StatusFilterTypeArrayOfDisputeStatus:
		// statusFilter.ArrayOfDisputeStatus is populated
}
```

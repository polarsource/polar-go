# RefundIDFilter

Filter by refund ID.


## Supported Types

### 

```go
refundIDFilter := operations.CreateRefundIDFilterStr(string{/* values here */})
```

### 

```go
refundIDFilter := operations.CreateRefundIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch refundIDFilter.Type {
	case operations.RefundIDFilterTypeStr:
		// refundIDFilter.Str is populated
	case operations.RefundIDFilterTypeArrayOfStr:
		// refundIDFilter.ArrayOfStr is populated
}
```

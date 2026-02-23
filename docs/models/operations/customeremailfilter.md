# CustomerEmailFilter

Filter by customer email.


## Supported Types

### 

```go
customerEmailFilter := operations.CreateCustomerEmailFilterStr(string{/* values here */})
```

### 

```go
customerEmailFilter := operations.CreateCustomerEmailFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerEmailFilter.Type {
	case operations.CustomerEmailFilterTypeStr:
		// customerEmailFilter.Str is populated
	case operations.CustomerEmailFilterTypeArrayOfStr:
		// customerEmailFilter.ArrayOfStr is populated
}
```

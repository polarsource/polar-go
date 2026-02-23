# CustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
customerIDFilter := operations.CreateCustomerIDFilterStr(string{/* values here */})
```

### 

```go
customerIDFilter := operations.CreateCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerIDFilter.Type {
	case operations.CustomerIDFilterTypeStr:
		// customerIDFilter.Str is populated
	case operations.CustomerIDFilterTypeArrayOfStr:
		// customerIDFilter.ArrayOfStr is populated
}
```

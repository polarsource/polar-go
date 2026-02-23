# ExternalCustomerIDFilter

Filter by customer external ID.


## Supported Types

### 

```go
externalCustomerIDFilter := operations.CreateExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
externalCustomerIDFilter := operations.CreateExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch externalCustomerIDFilter.Type {
	case operations.ExternalCustomerIDFilterTypeStr:
		// externalCustomerIDFilter.Str is populated
	case operations.ExternalCustomerIDFilterTypeArrayOfStr:
		// externalCustomerIDFilter.ArrayOfStr is populated
}
```

# EventTypesListQueryParamExternalCustomerIDFilter

Filter by external customer ID.


## Supported Types

### 

```go
eventTypesListQueryParamExternalCustomerIDFilter := operations.CreateEventTypesListQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
eventTypesListQueryParamExternalCustomerIDFilter := operations.CreateEventTypesListQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventTypesListQueryParamExternalCustomerIDFilter.Type {
	case operations.EventTypesListQueryParamExternalCustomerIDFilterTypeStr:
		// eventTypesListQueryParamExternalCustomerIDFilter.Str is populated
	case operations.EventTypesListQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// eventTypesListQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```

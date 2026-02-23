# EventTypesListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
eventTypesListQueryParamCustomerIDFilter := operations.CreateEventTypesListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
eventTypesListQueryParamCustomerIDFilter := operations.CreateEventTypesListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventTypesListQueryParamCustomerIDFilter.Type {
	case operations.EventTypesListQueryParamCustomerIDFilterTypeStr:
		// eventTypesListQueryParamCustomerIDFilter.Str is populated
	case operations.EventTypesListQueryParamCustomerIDFilterTypeArrayOfStr:
		// eventTypesListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

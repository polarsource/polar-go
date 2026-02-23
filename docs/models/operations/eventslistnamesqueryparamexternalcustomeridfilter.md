# EventsListNamesQueryParamExternalCustomerIDFilter

Filter by external customer ID.


## Supported Types

### 

```go
eventsListNamesQueryParamExternalCustomerIDFilter := operations.CreateEventsListNamesQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
eventsListNamesQueryParamExternalCustomerIDFilter := operations.CreateEventsListNamesQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsListNamesQueryParamExternalCustomerIDFilter.Type {
	case operations.EventsListNamesQueryParamExternalCustomerIDFilterTypeStr:
		// eventsListNamesQueryParamExternalCustomerIDFilter.Str is populated
	case operations.EventsListNamesQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// eventsListNamesQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```

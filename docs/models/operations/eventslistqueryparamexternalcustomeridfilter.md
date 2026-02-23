# EventsListQueryParamExternalCustomerIDFilter

Filter by external customer ID.


## Supported Types

### 

```go
eventsListQueryParamExternalCustomerIDFilter := operations.CreateEventsListQueryParamExternalCustomerIDFilterStr(string{/* values here */})
```

### 

```go
eventsListQueryParamExternalCustomerIDFilter := operations.CreateEventsListQueryParamExternalCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsListQueryParamExternalCustomerIDFilter.Type {
	case operations.EventsListQueryParamExternalCustomerIDFilterTypeStr:
		// eventsListQueryParamExternalCustomerIDFilter.Str is populated
	case operations.EventsListQueryParamExternalCustomerIDFilterTypeArrayOfStr:
		// eventsListQueryParamExternalCustomerIDFilter.ArrayOfStr is populated
}
```

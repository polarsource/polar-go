# EventsListNamesQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
eventsListNamesQueryParamCustomerIDFilter := operations.CreateEventsListNamesQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
eventsListNamesQueryParamCustomerIDFilter := operations.CreateEventsListNamesQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsListNamesQueryParamCustomerIDFilter.Type {
	case operations.EventsListNamesQueryParamCustomerIDFilterTypeStr:
		// eventsListNamesQueryParamCustomerIDFilter.Str is populated
	case operations.EventsListNamesQueryParamCustomerIDFilterTypeArrayOfStr:
		// eventsListNamesQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

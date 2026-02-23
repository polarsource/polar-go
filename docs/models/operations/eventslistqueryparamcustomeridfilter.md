# EventsListQueryParamCustomerIDFilter

Filter by customer ID.


## Supported Types

### 

```go
eventsListQueryParamCustomerIDFilter := operations.CreateEventsListQueryParamCustomerIDFilterStr(string{/* values here */})
```

### 

```go
eventsListQueryParamCustomerIDFilter := operations.CreateEventsListQueryParamCustomerIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsListQueryParamCustomerIDFilter.Type {
	case operations.EventsListQueryParamCustomerIDFilterTypeStr:
		// eventsListQueryParamCustomerIDFilter.Str is populated
	case operations.EventsListQueryParamCustomerIDFilterTypeArrayOfStr:
		// eventsListQueryParamCustomerIDFilter.ArrayOfStr is populated
}
```

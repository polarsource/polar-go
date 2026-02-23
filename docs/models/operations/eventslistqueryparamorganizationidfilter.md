# EventsListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
eventsListQueryParamOrganizationIDFilter := operations.CreateEventsListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
eventsListQueryParamOrganizationIDFilter := operations.CreateEventsListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsListQueryParamOrganizationIDFilter.Type {
	case operations.EventsListQueryParamOrganizationIDFilterTypeStr:
		// eventsListQueryParamOrganizationIDFilter.Str is populated
	case operations.EventsListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// eventsListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

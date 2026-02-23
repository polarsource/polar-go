# EventsListNamesQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
eventsListNamesQueryParamOrganizationIDFilter := operations.CreateEventsListNamesQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
eventsListNamesQueryParamOrganizationIDFilter := operations.CreateEventsListNamesQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsListNamesQueryParamOrganizationIDFilter.Type {
	case operations.EventsListNamesQueryParamOrganizationIDFilterTypeStr:
		// eventsListNamesQueryParamOrganizationIDFilter.Str is populated
	case operations.EventsListNamesQueryParamOrganizationIDFilterTypeArrayOfStr:
		// eventsListNamesQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

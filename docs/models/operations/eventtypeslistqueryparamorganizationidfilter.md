# EventTypesListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
eventTypesListQueryParamOrganizationIDFilter := operations.CreateEventTypesListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
eventTypesListQueryParamOrganizationIDFilter := operations.CreateEventTypesListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventTypesListQueryParamOrganizationIDFilter.Type {
	case operations.EventTypesListQueryParamOrganizationIDFilterTypeStr:
		// eventTypesListQueryParamOrganizationIDFilter.Str is populated
	case operations.EventTypesListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// eventTypesListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

# FilesListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
filesListQueryParamOrganizationIDFilter := operations.CreateFilesListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
filesListQueryParamOrganizationIDFilter := operations.CreateFilesListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch filesListQueryParamOrganizationIDFilter.Type {
	case operations.FilesListQueryParamOrganizationIDFilterTypeStr:
		// filesListQueryParamOrganizationIDFilter.Str is populated
	case operations.FilesListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// filesListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

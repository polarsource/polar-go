# OrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
organizationIDFilter := operations.CreateOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
organizationIDFilter := operations.CreateOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch organizationIDFilter.Type {
	case operations.OrganizationIDFilterTypeStr:
		// organizationIDFilter.Str is populated
	case operations.OrganizationIDFilterTypeArrayOfStr:
		// organizationIDFilter.ArrayOfStr is populated
}
```

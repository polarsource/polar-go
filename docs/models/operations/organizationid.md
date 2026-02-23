# OrganizationID

Filter by organization ID.


## Supported Types

### 

```go
organizationID := operations.CreateOrganizationIDStr(string{/* values here */})
```

### 

```go
organizationID := operations.CreateOrganizationIDArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch organizationID.Type {
	case operations.OrganizationIDTypeStr:
		// organizationID.Str is populated
	case operations.OrganizationIDTypeArrayOfStr:
		// organizationID.ArrayOfStr is populated
}
```

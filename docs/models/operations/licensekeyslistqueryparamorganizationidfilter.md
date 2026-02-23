# LicenseKeysListQueryParamOrganizationIDFilter

Filter by organization ID.


## Supported Types

### 

```go
licenseKeysListQueryParamOrganizationIDFilter := operations.CreateLicenseKeysListQueryParamOrganizationIDFilterStr(string{/* values here */})
```

### 

```go
licenseKeysListQueryParamOrganizationIDFilter := operations.CreateLicenseKeysListQueryParamOrganizationIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch licenseKeysListQueryParamOrganizationIDFilter.Type {
	case operations.LicenseKeysListQueryParamOrganizationIDFilterTypeStr:
		// licenseKeysListQueryParamOrganizationIDFilter.Str is populated
	case operations.LicenseKeysListQueryParamOrganizationIDFilterTypeArrayOfStr:
		// licenseKeysListQueryParamOrganizationIDFilter.ArrayOfStr is populated
}
```

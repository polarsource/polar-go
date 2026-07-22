# LicenseKeyStatusFilter

Filter by license key status.


## Supported Types

### LicenseKeyStatus

```go
licenseKeyStatusFilter := operations.CreateLicenseKeyStatusFilterLicenseKeyStatus(components.LicenseKeyStatus{/* values here */})
```

### 

```go
licenseKeyStatusFilter := operations.CreateLicenseKeyStatusFilterArrayOfLicenseKeyStatus([]components.LicenseKeyStatus{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch licenseKeyStatusFilter.Type {
	case operations.LicenseKeyStatusFilterTypeLicenseKeyStatus:
		// licenseKeyStatusFilter.LicenseKeyStatus is populated
	case operations.LicenseKeyStatusFilterTypeArrayOfLicenseKeyStatus:
		// licenseKeyStatusFilter.ArrayOfLicenseKeyStatus is populated
}
```

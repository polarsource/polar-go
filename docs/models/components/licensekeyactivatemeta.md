# LicenseKeyActivateMeta


## Supported Types

### 

```go
licenseKeyActivateMeta := components.CreateLicenseKeyActivateMetaStr(string{/* values here */})
```

### 

```go
licenseKeyActivateMeta := components.CreateLicenseKeyActivateMetaInteger(int64{/* values here */})
```

### 

```go
licenseKeyActivateMeta := components.CreateLicenseKeyActivateMetaNumber(float64{/* values here */})
```

### 

```go
licenseKeyActivateMeta := components.CreateLicenseKeyActivateMetaBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch licenseKeyActivateMeta.Type {
	case components.LicenseKeyActivateMetaTypeStr:
		// licenseKeyActivateMeta.Str is populated
	case components.LicenseKeyActivateMetaTypeInteger:
		// licenseKeyActivateMeta.Integer is populated
	case components.LicenseKeyActivateMetaTypeNumber:
		// licenseKeyActivateMeta.Number is populated
	case components.LicenseKeyActivateMetaTypeBoolean:
		// licenseKeyActivateMeta.Boolean is populated
}
```

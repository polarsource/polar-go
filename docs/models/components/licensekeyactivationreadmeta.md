# LicenseKeyActivationReadMeta


## Supported Types

### 

```go
licenseKeyActivationReadMeta := components.CreateLicenseKeyActivationReadMetaStr(string{/* values here */})
```

### 

```go
licenseKeyActivationReadMeta := components.CreateLicenseKeyActivationReadMetaInteger(int64{/* values here */})
```

### 

```go
licenseKeyActivationReadMeta := components.CreateLicenseKeyActivationReadMetaNumber(float64{/* values here */})
```

### 

```go
licenseKeyActivationReadMeta := components.CreateLicenseKeyActivationReadMetaBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch licenseKeyActivationReadMeta.Type {
	case components.LicenseKeyActivationReadMetaTypeStr:
		// licenseKeyActivationReadMeta.Str is populated
	case components.LicenseKeyActivationReadMetaTypeInteger:
		// licenseKeyActivationReadMeta.Integer is populated
	case components.LicenseKeyActivationReadMetaTypeNumber:
		// licenseKeyActivationReadMeta.Number is populated
	case components.LicenseKeyActivationReadMetaTypeBoolean:
		// licenseKeyActivationReadMeta.Boolean is populated
}
```

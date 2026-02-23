# LicenseKeyActivateConditions


## Supported Types

### 

```go
licenseKeyActivateConditions := components.CreateLicenseKeyActivateConditionsStr(string{/* values here */})
```

### 

```go
licenseKeyActivateConditions := components.CreateLicenseKeyActivateConditionsInteger(int64{/* values here */})
```

### 

```go
licenseKeyActivateConditions := components.CreateLicenseKeyActivateConditionsNumber(float64{/* values here */})
```

### 

```go
licenseKeyActivateConditions := components.CreateLicenseKeyActivateConditionsBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch licenseKeyActivateConditions.Type {
	case components.LicenseKeyActivateConditionsTypeStr:
		// licenseKeyActivateConditions.Str is populated
	case components.LicenseKeyActivateConditionsTypeInteger:
		// licenseKeyActivateConditions.Integer is populated
	case components.LicenseKeyActivateConditionsTypeNumber:
		// licenseKeyActivateConditions.Number is populated
	case components.LicenseKeyActivateConditionsTypeBoolean:
		// licenseKeyActivateConditions.Boolean is populated
}
```

# LicenseKeyCustomerTaxID


## Supported Types

### 

```go
licenseKeyCustomerTaxID := components.CreateLicenseKeyCustomerTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
licenseKeyCustomerTaxID := components.CreateLicenseKeyCustomerTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch licenseKeyCustomerTaxID.Type {
	case components.LicenseKeyCustomerTaxIDTypeStr:
		// licenseKeyCustomerTaxID.Str is populated
	case components.LicenseKeyCustomerTaxIDTypeTaxIDFormat:
		// licenseKeyCustomerTaxID.TaxIDFormat is populated
}
```

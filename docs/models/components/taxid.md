# TaxID


## Supported Types

### 

```go
taxID := components.CreateTaxIDStr(string{/* values here */})
```

### TaxIDFormat

```go
taxID := components.CreateTaxIDTaxIDFormat(components.TaxIDFormat{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch taxID.Type {
	case components.TaxIDTypeStr:
		// taxID.Str is populated
	case components.TaxIDTypeTaxIDFormat:
		// taxID.TaxIDFormat is populated
}
```

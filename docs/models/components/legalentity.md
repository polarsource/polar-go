# LegalEntity


## Supported Types

### OrganizationCompanyLegalEntitySchema

```go
legalEntity := components.CreateLegalEntityCompany(components.OrganizationCompanyLegalEntitySchema{/* values here */})
```

### OrganizationIndividualLegalEntitySchema

```go
legalEntity := components.CreateLegalEntityIndividual(components.OrganizationIndividualLegalEntitySchema{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch legalEntity.Type {
	case components.LegalEntityTypeCompany:
		// legalEntity.OrganizationCompanyLegalEntitySchema is populated
	case components.LegalEntityTypeIndividual:
		// legalEntity.OrganizationIndividualLegalEntitySchema is populated
}
```

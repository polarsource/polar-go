# CustomerCreate


## Supported Types

### CustomerIndividualCreate

```go
customerCreate := components.CreateCustomerCreateCustomerIndividualCreate(components.CustomerIndividualCreate{/* values here */})
```

### CustomerTeamCreate

```go
customerCreate := components.CreateCustomerCreateCustomerTeamCreate(components.CustomerTeamCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerCreate.Type {
	case components.CustomerCreateTypeCustomerIndividualCreate:
		// customerCreate.CustomerIndividualCreate is populated
	case components.CustomerCreateTypeCustomerTeamCreate:
		// customerCreate.CustomerTeamCreate is populated
}
```

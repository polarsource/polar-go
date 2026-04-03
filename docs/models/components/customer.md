# Customer


## Supported Types

### CustomerIndividual

```go
customer := components.CreateCustomerIndividual(components.CustomerIndividual{/* values here */})
```

### CustomerTeam

```go
customer := components.CreateCustomerTeam(components.CustomerTeam{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customer.Type {
	case components.CustomerUnionTypeIndividual:
		// customer.CustomerIndividual is populated
	case components.CustomerUnionTypeTeam:
		// customer.CustomerTeam is populated
}
```

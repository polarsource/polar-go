# CustomerState


## Supported Types

### CustomerStateIndividual

```go
customerState := components.CreateCustomerStateIndividual(components.CustomerStateIndividual{/* values here */})
```

### CustomerStateTeam

```go
customerState := components.CreateCustomerStateTeam(components.CustomerStateTeam{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerState.Type {
	case components.CustomerStateTypeIndividual:
		// customerState.CustomerStateIndividual is populated
	case components.CustomerStateTypeTeam:
		// customerState.CustomerStateTeam is populated
}
```

# CustomerSessionsCreateCustomerSessionCreate


## Supported Types

### CustomerSessionCustomerIDCreate

```go
customerSessionsCreateCustomerSessionCreate := operations.CreateCustomerSessionsCreateCustomerSessionCreateCustomerSessionCustomerIDCreate(components.CustomerSessionCustomerIDCreate{/* values here */})
```

### CustomerSessionCustomerExternalIDCreate

```go
customerSessionsCreateCustomerSessionCreate := operations.CreateCustomerSessionsCreateCustomerSessionCreateCustomerSessionCustomerExternalIDCreate(components.CustomerSessionCustomerExternalIDCreate{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerSessionsCreateCustomerSessionCreate.Type {
	case operations.CustomerSessionsCreateCustomerSessionCreateTypeCustomerSessionCustomerIDCreate:
		// customerSessionsCreateCustomerSessionCreate.CustomerSessionCustomerIDCreate is populated
	case operations.CustomerSessionsCreateCustomerSessionCreateTypeCustomerSessionCustomerExternalIDCreate:
		// customerSessionsCreateCustomerSessionCreate.CustomerSessionCustomerExternalIDCreate is populated
}
```

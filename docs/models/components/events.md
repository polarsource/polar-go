# Events


## Supported Types

### EventCreateCustomer

```go
events := components.CreateEventsEventCreateCustomer(components.EventCreateCustomer{/* values here */})
```

### EventCreateExternalCustomer

```go
events := components.CreateEventsEventCreateExternalCustomer(components.EventCreateExternalCustomer{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch events.Type {
	case components.EventsTypeEventCreateCustomer:
		// events.EventCreateCustomer is populated
	case components.EventsTypeEventCreateExternalCustomer:
		// events.EventCreateExternalCustomer is populated
}
```

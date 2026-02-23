# Event


## Supported Types

### SystemEvent

```go
event := components.CreateEventSystemEvent(components.SystemEvent{/* values here */})
```

### UserEvent

```go
event := components.CreateEventUserEvent(components.UserEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch event.Type {
	case components.EventUnionTypeSystemEvent:
		// event.SystemEvent is populated
	case components.EventUnionTypeUserEvent:
		// event.UserEvent is populated
}
```

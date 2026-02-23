# EventType

Filter by webhook event type.


## Supported Types

### WebhookEventType

```go
eventType := operations.CreateEventTypeWebhookEventType(components.WebhookEventType{/* values here */})
```

### 

```go
eventType := operations.CreateEventTypeArrayOfWebhookEventType([]components.WebhookEventType{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventType.Type {
	case operations.EventTypeTypeWebhookEventType:
		// eventType.WebhookEventType is populated
	case operations.EventTypeTypeArrayOfWebhookEventType:
		// eventType.ArrayOfWebhookEventType is populated
}
```

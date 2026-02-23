# EventsListResponseEventsList

Successful Response


## Supported Types

### ListResourceEvent

```go
eventsListResponseEventsList := operations.CreateEventsListResponseEventsListListResourceEvent(components.ListResourceEvent{/* values here */})
```

### ListResourceWithCursorPaginationEvent

```go
eventsListResponseEventsList := operations.CreateEventsListResponseEventsListListResourceWithCursorPaginationEvent(components.ListResourceWithCursorPaginationEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventsListResponseEventsList.Type {
	case operations.EventsListResponseEventsListTypeListResourceEvent:
		// eventsListResponseEventsList.ListResourceEvent is populated
	case operations.EventsListResponseEventsListTypeListResourceWithCursorPaginationEvent:
		// eventsListResponseEventsList.ListResourceWithCursorPaginationEvent is populated
}
```

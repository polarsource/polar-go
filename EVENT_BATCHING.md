# Event Batching

Event batching is an optional performance optimization feature that automatically queues event calls and flushes them in batches.

## How It Works

When enabled, the event batcher:
- Queues `Events.Ingest()` calls instead of sending them immediately
- Automatically flushes when **10 events** are queued
- Automatically flushes after **5 seconds** if fewer than 10 events are queued
- Returns synthetic responses with correct `inserted` count for each call

## Enabling Event Batching

Event batching is **disabled by default**. To enable it, set the environment variable:

```bash
export POLAR_ENABLE_EVENT_BATCHING=true
```

Or in your Go code:

```go
import "os"

func main() {
    os.Setenv("POLAR_ENABLE_EVENT_BATCHING", "true")

    client := polargo.New(polargo.WithSecurity("YOUR_ACCESS_TOKEN"))
    // ... use client
}
```

## Usage Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"

    polargo "github.com/polarsource/polar-go"
    "github.com/polarsource/polar-go/internal/hooks"
    "github.com/polarsource/polar-go/models/components"
)

func main() {
    // Enable event batching
    os.Setenv("POLAR_ENABLE_EVENT_BATCHING", "true")

    client := polargo.New(
        polargo.WithSecurity("YOUR_ACCESS_TOKEN"),
    )

    ctx := context.Background()

    // Send events - these will be batched
    for i := 0; i < 15; i++ {
        timestamp := time.Now()
        event := components.CreateEventsEventCreateCustomer(components.EventCreateCustomer{
            Name:       fmt.Sprintf("event-%d", i),
            CustomerID: "customer-123",
            Timestamp:  &timestamp,
        })

        resp, err := client.Events.Ingest(ctx, components.EventsIngest{
            Events: []components.Events{event},
        })
        if err != nil {
            fmt.Printf("Error: %v\n", err)
            continue
        }
        // Response shows inserted count for this individual call
        fmt.Printf("Inserted: %d\n", resp.EventsIngestResponse.Inserted)
    }

    // Flush remaining events before shutdown
    if err := hooks.FlushEvents(); err != nil {
        fmt.Printf("Error flushing events: %v\n", err)
    }
}
```

## Graceful Shutdown

Always call `hooks.FlushEvents()` before your application exits to ensure all queued events are sent:

```go
import "github.com/polarsource/polar-go/internal/hooks"

func main() {
    // ... your application code

    // Flush before shutdown
    if err := hooks.FlushEvents(); err != nil {
        log.Printf("Error flushing events: %v", err)
    }
}
```

## Configuration

| Setting | Value | Description |
|---------|-------|-------------|
| Batch Size | 10 events | Automatic flush when queue reaches this size |
| Flush Interval | 5 seconds | Automatic flush after this time period |
| Environment Variable | `POLAR_ENABLE_EVENT_BATCHING=true` | Required to enable batching |

## Default Behavior (Batching Disabled)

When `POLAR_ENABLE_EVENT_BATCHING` is not set or set to any value other than `"true"`:
- Events are sent immediately without batching
- Each `Events.Ingest()` call results in an actual HTTP request
- `hooks.FlushEvents()` is a safe no-op
- This is the default SDK behavior

## Thread Safety

The event batcher is thread-safe and can be used concurrently from multiple goroutines.

## Troubleshooting

### Events not being sent

Make sure to call `hooks.FlushEvents()` before your application exits, especially if you're sending fewer than 10 events or your application exits within 5 seconds.

### Batching not working

Verify that the environment variable is set correctly:

```go
fmt.Println(os.Getenv("POLAR_ENABLE_EVENT_BATCHING")) // Should print "true"
```

The environment variable must be set **before** creating the Polar client.

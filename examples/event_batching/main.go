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
	os.Setenv("POLAR_ENABLE_EVENT_BATCHING", "true")

	client := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	ctx := context.Background()

	fmt.Println("Example 1: Auto-batching at 10 events")
	for i := 0; i < 12; i++ {
		timestamp := time.Now()
		event := components.CreateEventsEventCreateCustomer(components.EventCreateCustomer{
			Name:       fmt.Sprintf("customer-batch-%d", i),
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
		fmt.Printf("Event %d: Inserted count = %d\n", i, resp.EventsIngestResponse.Inserted)
	}

	fmt.Println("\nExample 2: Manual flush")
	for i := 0; i < 5; i++ {
		timestamp := time.Now()
		event := components.CreateEventsEventCreateCustomer(components.EventCreateCustomer{
			Name:       fmt.Sprintf("customer-manual-%d", i),
			CustomerID: "customer-456",
			Timestamp:  &timestamp,
		})

		resp, err := client.Events.Ingest(ctx, components.EventsIngest{
			Events: []components.Events{event},
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Event %d: Inserted count = %d\n", i, resp.EventsIngestResponse.Inserted)
	}

	fmt.Println("Flushing manually...")
	if err := hooks.FlushEvents(); err != nil {
		fmt.Printf("Error flushing events: %v\n", err)
	}

	fmt.Println("\nExample 3: Timeout flush (wait 6 seconds)")
	for i := 0; i < 3; i++ {
		timestamp := time.Now()
		event := components.CreateEventsEventCreateCustomer(components.EventCreateCustomer{
			Name:       fmt.Sprintf("customer-timeout-%d", i),
			CustomerID: "customer-789",
			Timestamp:  &timestamp,
		})

		resp, err := client.Events.Ingest(ctx, components.EventsIngest{
			Events: []components.Events{event},
		})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Event %d: Inserted count = %d\n", i, resp.EventsIngestResponse.Inserted)
	}

	fmt.Println("Waiting for timeout flush...")
	time.Sleep(6 * time.Second)

	fmt.Println("\nFinal flush before shutdown")
	if err := hooks.FlushEvents(); err != nil {
		fmt.Printf("Error flushing events: %v\n", err)
	}

	fmt.Println("Done!")
}

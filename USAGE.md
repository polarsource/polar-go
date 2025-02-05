<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.ExternalOrganizations.List(ctx, operations.ExternalOrganizationsListRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceExternalOrganization != nil {
		for {
			// handle items

			res, err = res.Next()

			if err != nil {
				// handle error
			}

			if res == nil {
				break
			}
		}
	}
}

```
<!-- End SDK Example Usage [usage] -->
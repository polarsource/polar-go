<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	polargo "github.com/polarsource/polar-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := polargo.New(
		polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
	)

	res, err := s.Organizations.List(ctx, nil, polargo.Pointer[int64](1), polargo.Pointer[int64](10), nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.ListResourceOrganization != nil {
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
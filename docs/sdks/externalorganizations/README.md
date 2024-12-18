# ExternalOrganizations
(*ExternalOrganizations*)

## Overview

### Available Operations

* [List](#list) - List External Organizations

## List

List external organizations.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"polar/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
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

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ExternalOrganizationsListRequest](../../models/operations/externalorganizationslistrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ExternalOrganizationsListResponse](../../models/operations/externalorganizationslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
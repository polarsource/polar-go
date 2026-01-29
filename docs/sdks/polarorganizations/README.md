# CustomerPortal.Organizations

## Overview

### Available Operations

* [Get](#get) - Get Organization

## Get

Get a customer portal's organization by slug.

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:organizations:get" method="get" path="/v1/customer-portal/organizations/{slug}" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.Organizations.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerOrganizationData != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `slug`                                                   | *string*                                                 | :heavy_check_mark:                                       | The organization slug.                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomerPortalOrganizationsGetResponse](../../models/operations/customerportalorganizationsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
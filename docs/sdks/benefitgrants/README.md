# BenefitGrants
(*BenefitGrants*)

## Overview

### Available Operations

* [List](#list) - List Benefit Grants

## List

List benefit grants across all benefits for the authenticated organization.

**Scopes**: `benefits:read` `benefits:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="benefit-grants:list" method="get" path="/v1/benefit-grants/" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.BenefitGrants.List(ctx, operations.BenefitGrantsListRequest{
        OrganizationID: polargo.Pointer(operations.CreateBenefitGrantsListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceBenefitGrant != nil {
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.BenefitGrantsListRequest](../../models/operations/benefitgrantslistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.BenefitGrantsListResponse](../../models/operations/benefitgrantslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
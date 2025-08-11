# Downloadables
(*CustomerPortal.Downloadables*)

## Overview

### Available Operations

* [List](#list) - List Downloadables

## List

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:downloadables:list" method="get" path="/v1/customer-portal/downloadables/" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"os"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.Downloadables.List(ctx, operations.CustomerPortalDownloadablesListSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, polargo.Pointer(operations.CreateCustomerPortalDownloadablesListQueryParamOrganizationIDFilterStr(
        "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
    )), nil, polargo.Int64(1), polargo.Int64(10))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceDownloadableRead != nil {
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

| Parameter                                                                                                                                                             | Type                                                                                                                                                                  | Required                                                                                                                                                              | Description                                                                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                 | :heavy_check_mark:                                                                                                                                                    | The context to use for the request.                                                                                                                                   |
| `security`                                                                                                                                                            | [operations.CustomerPortalDownloadablesListSecurity](../../models/operations/customerportaldownloadableslistsecurity.md)                                              | :heavy_check_mark:                                                                                                                                                    | The security requirements to use for the request.                                                                                                                     |
| `organizationID`                                                                                                                                                      | [*operations.CustomerPortalDownloadablesListQueryParamOrganizationIDFilter](../../models/operations/customerportaldownloadableslistqueryparamorganizationidfilter.md) | :heavy_minus_sign:                                                                                                                                                    | Filter by organization ID.                                                                                                                                            |
| `benefitID`                                                                                                                                                           | [*operations.CustomerPortalDownloadablesListQueryParamBenefitIDFilter](../../models/operations/customerportaldownloadableslistqueryparambenefitidfilter.md)           | :heavy_minus_sign:                                                                                                                                                    | Filter by benefit ID.                                                                                                                                                 |
| `page`                                                                                                                                                                | **int64*                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                    | Page number, defaults to 1.                                                                                                                                           |
| `limit`                                                                                                                                                               | **int64*                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                    | Size of a page, defaults to 10. Maximum is 100.                                                                                                                       |
| `opts`                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                              | :heavy_minus_sign:                                                                                                                                                    | The options for this request.                                                                                                                                         |

### Response

**[*operations.CustomerPortalDownloadablesListResponse](../../models/operations/customerportaldownloadableslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
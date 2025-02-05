# Downloadables
(*CustomerPortal.Downloadables*)

## Overview

### Available Operations

* [List](#list) - List Downloadables
* [Get](#get) - Get Downloadable

## List

List Downloadables

### Example Usage

```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.CustomerPortal.Downloadables.List(ctx, nil, nil, nil, nil)
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

## Get

Get Downloadable

### Example Usage

```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.CustomerPortal.Downloadables.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `token`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomerPortalDownloadablesCustomerPortalDownloadablesGetResponse](../../models/operations/customerportaldownloadablescustomerportaldownloadablesgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
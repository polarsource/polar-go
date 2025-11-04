# PolarWallets
(*CustomerPortal.Wallets*)

## Overview

### Available Operations

* [List](#list) - List Wallets
* [Get](#get) - Get Wallet

## List

List wallets of the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:wallets:list" method="get" path="/v1/customer-portal/wallets/" -->
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

    res, err := s.CustomerPortal.Wallets.List(ctx, operations.CustomerPortalWalletsListSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, polargo.Pointer[int64](1), polargo.Pointer[int64](10), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceCustomerWallet != nil {
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

| Parameter                                                                                                                                                               | Type                                                                                                                                                                    | Required                                                                                                                                                                | Description                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                   | :heavy_check_mark:                                                                                                                                                      | The context to use for the request.                                                                                                                                     |
| `security`                                                                                                                                                              | [operations.CustomerPortalWalletsListSecurity](../../models/operations/customerportalwalletslistsecurity.md)                                                            | :heavy_check_mark:                                                                                                                                                      | The security requirements to use for the request.                                                                                                                       |
| `page`                                                                                                                                                                  | **int64*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | Page number, defaults to 1.                                                                                                                                             |
| `limit`                                                                                                                                                                 | **int64*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | Size of a page, defaults to 10. Maximum is 100.                                                                                                                         |
| `sorting`                                                                                                                                                               | [][components.CustomerWalletSortProperty](../../models/components/customerwalletsortproperty.md)                                                                        | :heavy_minus_sign:                                                                                                                                                      | Sorting criterion. Several criteria can be used simultaneously and will be applied in order. Add a minus sign `-` before the criteria name to sort by descending order. |
| `opts`                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | The options for this request.                                                                                                                                           |

### Response

**[*operations.CustomerPortalWalletsListResponse](../../models/operations/customerportalwalletslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a wallet by ID for the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:wallets:get" method="get" path="/v1/customer-portal/wallets/{id}" -->
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

    res, err := s.CustomerPortal.Wallets.Get(ctx, operations.CustomerPortalWalletsGetSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerWallet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `security`                                                                                                 | [operations.CustomerPortalWalletsGetSecurity](../../models/operations/customerportalwalletsgetsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `id`                                                                                                       | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The wallet ID.                                                                                             |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CustomerPortalWalletsGetResponse](../../models/operations/customerportalwalletsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
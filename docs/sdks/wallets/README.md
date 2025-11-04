# Wallets
(*Wallets*)

## Overview

### Available Operations

* [List](#list) - List Wallets
* [Get](#get) - Get Wallet
* [TopUp](#topup) - Top-Up Wallet

## List

List wallets.

**Scopes**: `wallets:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="wallets:list" method="get" path="/v1/wallets/" -->
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

    res, err := s.Wallets.List(ctx, operations.WalletsListRequest{
        OrganizationID: polargo.Pointer(operations.CreateWalletsListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceWallet != nil {
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.WalletsListRequest](../../models/operations/walletslistrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.WalletsListResponse](../../models/operations/walletslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a wallet by ID.

**Scopes**: `wallets:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="wallets:get" method="get" path="/v1/wallets/{id}" -->
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

    res, err := s.Wallets.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Wallet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The wallet ID.                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WalletsGetResponse](../../models/operations/walletsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## TopUp

Top-up a wallet by adding funds to its balance.

The customer should have a valid payment method on file.

**Scopes**: `wallets:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="wallets:top_up" method="post" path="/v1/wallets/{id}/top-up" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Wallets.TopUp(ctx, "<value>", components.WalletTopUpCreate{
        Amount: 2000,
        Currency: "Venezuelan bolívar",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Wallet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The wallet ID.                                                               |
| `walletTopUpCreate`                                                          | [components.WalletTopUpCreate](../../models/components/wallettopupcreate.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.WalletsTopUpResponse](../../models/operations/walletstopupresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.PaymentIntentFailedError  | 400                                 | application/json                    |
| apierrors.MissingPaymentMethodError | 402                                 | application/json                    |
| apierrors.ResourceNotFound          | 404                                 | application/json                    |
| apierrors.HTTPValidationError       | 422                                 | application/json                    |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |
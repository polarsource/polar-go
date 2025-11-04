# Subscriptions
(*Subscriptions*)

## Overview

### Available Operations

* [List](#list) - List Subscriptions
* [Create](#create) - Create Subscription
* [Export](#export) - Export Subscriptions
* [Get](#get) - Get Subscription
* [Update](#update) - Update Subscription
* [Revoke](#revoke) - Revoke Subscription

## List

List subscriptions.

**Scopes**: `subscriptions:read` `subscriptions:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="subscriptions:list" method="get" path="/v1/subscriptions/" -->
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

    res, err := s.Subscriptions.List(ctx, operations.SubscriptionsListRequest{
        OrganizationID: polargo.Pointer(operations.CreateOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceSubscription != nil {
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
| `request`                                                                                  | [operations.SubscriptionsListRequest](../../models/operations/subscriptionslistrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.SubscriptionsListResponse](../../models/operations/subscriptionslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a subscription programmatically.

This endpoint only allows to create subscription on free products.
For paid products, use the checkout flow.

No initial order will be created and no confirmation email will be sent.

**Scopes**: `subscriptions:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="subscriptions:create" method="post" path="/v1/subscriptions/" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Subscriptions.Create(ctx, operations.CreateSubscriptionsCreateSubscriptionCreateSubscriptionCreateCustomer(
        components.SubscriptionCreateCustomer{
            ProductID: "d8dd2de1-21b7-4a41-8bc3-ce909c0cfe23",
            CustomerID: "992fae2a-2a17-4b7a-8d9e-e287cf90131b",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.SubscriptionsCreateSubscriptionCreate](../../models/operations/subscriptionscreatesubscriptioncreate.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.SubscriptionsCreateResponse](../../models/operations/subscriptionscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Export

Export subscriptions as a CSV file.

**Scopes**: `subscriptions:read` `subscriptions:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="subscriptions:export" method="get" path="/v1/subscriptions/export" -->
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

    res, err := s.Subscriptions.Export(ctx, polargo.Pointer(operations.CreateOrganizationIDStr(
        "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
    )))
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `organizationID`                                                        | [*operations.OrganizationID](../../models/operations/organizationid.md) | :heavy_minus_sign:                                                      | Filter by organization ID.                                              |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |

### Response

**[*operations.SubscriptionsExportResponse](../../models/operations/subscriptionsexportresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a subscription by ID.

**Scopes**: `subscriptions:read` `subscriptions:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="subscriptions:get" method="get" path="/v1/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The subscription ID.                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscriptionsGetResponse](../../models/operations/subscriptionsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a subscription.

**Scopes**: `subscriptions:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="subscriptions:update" method="patch" path="/v1/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.Update(ctx, "<value>", components.CreateSubscriptionUpdateSubscriptionUpdateProduct(
        components.SubscriptionUpdateProduct{
            ProductID: "<value>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The subscription ID.                                                           |
| `subscriptionUpdate`                                                           | [components.SubscriptionUpdate](../../models/components/subscriptionupdate.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.SubscriptionsUpdateResponse](../../models/operations/subscriptionsupdateresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.AlreadyCanceledSubscription | 403                                   | application/json                      |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.SubscriptionLocked          | 409                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Revoke

Revoke a subscription, i.e cancel immediately.

**Scopes**: `subscriptions:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="subscriptions:revoke" method="delete" path="/v1/subscriptions/{id}" -->
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

    res, err := s.Subscriptions.Revoke(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Subscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The subscription ID.                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.SubscriptionsRevokeResponse](../../models/operations/subscriptionsrevokeresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.AlreadyCanceledSubscription | 403                                   | application/json                      |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.SubscriptionLocked          | 409                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |
# PolarSubscriptions
(*CustomerPortal.Subscriptions*)

## Overview

### Available Operations

* [List](#list) - List Subscriptions
* [Get](#get) - Get Subscription
* [Update](#update) - Update Subscription
* [Cancel](#cancel) - Cancel Subscription

## List

List subscriptions of the authenticated customer or user.

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

    res, err := s.CustomerPortal.Subscriptions.List(ctx, operations.CustomerPortalSubscriptionsListRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceCustomerSubscription != nil {
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.CustomerPortalSubscriptionsListRequest](../../models/operations/customerportalsubscriptionslistrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CustomerPortalSubscriptionsListResponse](../../models/operations/customerportalsubscriptionslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a subscription for the authenticated customer or user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.CustomerPortal.Subscriptions.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSubscription != nil {
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

**[*operations.CustomerPortalSubscriptionsGetResponse](../../models/operations/customerportalsubscriptionsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a subscription of the authenticated customer or user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"polar/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.CustomerPortal.Subscriptions.Update(ctx, "<value>", components.CustomerSubscriptionUpdate{
        ProductPriceID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | The subscription ID.                                                                           |
| `customerSubscriptionUpdate`                                                                   | [components.CustomerSubscriptionUpdate](../../models/components/customersubscriptionupdate.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CustomerPortalSubscriptionsUpdateResponse](../../models/operations/customerportalsubscriptionsupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Cancel

Cancel a subscription of the authenticated customer or user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"polar"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := polar.New(
        polar.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.CustomerPortal.Subscriptions.Cancel(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSubscription != nil {
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

**[*operations.CustomerPortalSubscriptionsCancelResponse](../../models/operations/customerportalsubscriptionscancelresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.AlreadyCanceledSubscription | 403                                   | application/json                      |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |
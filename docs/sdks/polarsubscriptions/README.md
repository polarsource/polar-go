# PolarSubscriptions
(*CustomerPortal.Subscriptions*)

## Overview

### Available Operations

* [List](#list) - List Subscriptions
* [Get](#get) - Get Subscription
* [Update](#update) - Update Subscription
* [Cancel](#cancel) - Cancel Subscription

## List

List subscriptions of the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:subscriptions:list" method="get" path="/v1/customer-portal/subscriptions/" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/operations"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.Subscriptions.List(ctx, operations.CustomerPortalSubscriptionsListRequest{
        OrganizationID: polargo.Pointer(operations.CreateCustomerPortalSubscriptionsListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    }, operations.CustomerPortalSubscriptionsListSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.CustomerPortalSubscriptionsListRequest](../../models/operations/customerportalsubscriptionslistrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.CustomerPortalSubscriptionsListSecurity](../../models/operations/customerportalsubscriptionslistsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CustomerPortalSubscriptionsListResponse](../../models/operations/customerportalsubscriptionslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a subscription for the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:subscriptions:get" method="get" path="/v1/customer-portal/subscriptions/{id}" -->
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

    res, err := s.CustomerPortal.Subscriptions.Get(ctx, operations.CustomerPortalSubscriptionsGetSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `security`                                                                                                             | [operations.CustomerPortalSubscriptionsGetSecurity](../../models/operations/customerportalsubscriptionsgetsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The subscription ID.                                                                                                   |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CustomerPortalSubscriptionsGetResponse](../../models/operations/customerportalsubscriptionsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a subscription of the authenticated customer.

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:subscriptions:update" method="patch" path="/v1/customer-portal/subscriptions/{id}" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"os"
	"github.com/polarsource/polar-go/models/operations"
	"github.com/polarsource/polar-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.Subscriptions.Update(ctx, operations.CustomerPortalSubscriptionsUpdateSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>", components.CreateCustomerSubscriptionUpdateCustomerSubscriptionCancel(
        components.CustomerSubscriptionCancel{},
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `security`                                                                                                                   | [operations.CustomerPortalSubscriptionsUpdateSecurity](../../models/operations/customerportalsubscriptionsupdatesecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |
| `id`                                                                                                                         | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The subscription ID.                                                                                                         |
| `customerSubscriptionUpdate`                                                                                                 | [components.CustomerSubscriptionUpdate](../../models/components/customersubscriptionupdate.md)                               | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CustomerPortalSubscriptionsUpdateResponse](../../models/operations/customerportalsubscriptionsupdateresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.AlreadyCanceledSubscription | 403                                   | application/json                      |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Cancel

Cancel a subscription of the authenticated customer.

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:subscriptions:cancel" method="delete" path="/v1/customer-portal/subscriptions/{id}" -->
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

    res, err := s.CustomerPortal.Subscriptions.Cancel(ctx, operations.CustomerPortalSubscriptionsCancelSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `security`                                                                                                                   | [operations.CustomerPortalSubscriptionsCancelSecurity](../../models/operations/customerportalsubscriptionscancelsecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |
| `id`                                                                                                                         | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The subscription ID.                                                                                                         |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CustomerPortalSubscriptionsCancelResponse](../../models/operations/customerportalsubscriptionscancelresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.AlreadyCanceledSubscription | 403                                   | application/json                      |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |
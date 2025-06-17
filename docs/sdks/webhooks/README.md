# Webhooks
(*Webhooks*)

## Overview

### Available Operations

* [ListWebhookEndpoints](#listwebhookendpoints) - List Webhook Endpoints
* [CreateWebhookEndpoint](#createwebhookendpoint) - Create Webhook Endpoint
* [GetWebhookEndpoint](#getwebhookendpoint) - Get Webhook Endpoint
* [UpdateWebhookEndpoint](#updatewebhookendpoint) - Update Webhook Endpoint
* [DeleteWebhookEndpoint](#deletewebhookendpoint) - Delete Webhook Endpoint
* [ListWebhookDeliveries](#listwebhookdeliveries) - List Webhook Deliveries
* [RedeliverWebhookEvent](#redeliverwebhookevent) - Redeliver Webhook Event

## ListWebhookEndpoints

List webhook endpoints.

**Scopes**: `webhooks:read` `webhooks:write`

### Example Usage

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

    res, err := s.Webhooks.ListWebhookEndpoints(ctx, polargo.Pointer(operations.CreateQueryParamOrganizationIDStr(
        "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
    )), polargo.Int64(1), polargo.Int64(10))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceWebhookEndpoint != nil {
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

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `organizationID`                                                                            | [*operations.QueryParamOrganizationID](../../models/operations/queryparamorganizationid.md) | :heavy_minus_sign:                                                                          | Filter by organization ID.                                                                  |
| `page`                                                                                      | **int64*                                                                                    | :heavy_minus_sign:                                                                          | Page number, defaults to 1.                                                                 |
| `limit`                                                                                     | **int64*                                                                                    | :heavy_minus_sign:                                                                          | Size of a page, defaults to 10. Maximum is 100.                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.WebhooksListWebhookEndpointsResponse](../../models/operations/webhookslistwebhookendpointsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## CreateWebhookEndpoint

Create a webhook endpoint.

**Scopes**: `webhooks:write`

### Example Usage

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

    res, err := s.Webhooks.CreateWebhookEndpoint(ctx, components.WebhookEndpointCreate{
        URL: "https://webhook.site/cb791d80-f26e-4f8c-be88-6e56054192b0",
        Format: components.WebhookFormatSlack,
        Secret: "f_z6mfSpxkjogyw3FkA2aH2gYE5huxruNf34MpdWMcA",
        Events: []components.WebhookEventType{
            components.WebhookEventTypeSubscriptionUncanceled,
        },
        OrganizationID: polargo.String("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookEndpoint != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.WebhookEndpointCreate](../../models/components/webhookendpointcreate.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.WebhooksCreateWebhookEndpointResponse](../../models/operations/webhookscreatewebhookendpointresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetWebhookEndpoint

Get a webhook endpoint by ID.

**Scopes**: `webhooks:read` `webhooks:write`

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

    res, err := s.Webhooks.GetWebhookEndpoint(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookEndpoint != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The webhook endpoint ID.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WebhooksGetWebhookEndpointResponse](../../models/operations/webhooksgetwebhookendpointresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## UpdateWebhookEndpoint

Update a webhook endpoint.

**Scopes**: `webhooks:write`

### Example Usage

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

    res, err := s.Webhooks.UpdateWebhookEndpoint(ctx, "<value>", components.WebhookEndpointUpdate{
        URL: polargo.String("https://webhook.site/cb791d80-f26e-4f8c-be88-6e56054192b0"),
        Secret: polargo.String("f_z6mfSpxkjogyw3FkA2aH2gYE5huxruNf34MpdWMcA"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WebhookEndpoint != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | The webhook endpoint ID.                                                             |
| `webhookEndpointUpdate`                                                              | [components.WebhookEndpointUpdate](../../models/components/webhookendpointupdate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.WebhooksUpdateWebhookEndpointResponse](../../models/operations/webhooksupdatewebhookendpointresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## DeleteWebhookEndpoint

Delete a webhook endpoint.

**Scopes**: `webhooks:write`

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

    res, err := s.Webhooks.DeleteWebhookEndpoint(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The webhook endpoint ID.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WebhooksDeleteWebhookEndpointResponse](../../models/operations/webhooksdeletewebhookendpointresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ListWebhookDeliveries

List webhook deliveries.

Deliveries are all the attempts to deliver a webhook event to an endpoint.

**Scopes**: `webhooks:read` `webhooks:write`

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

    res, err := s.Webhooks.ListWebhookDeliveries(ctx, nil, polargo.Int64(1), polargo.Int64(10))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceWebhookDelivery != nil {
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

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |
| `endpointID`                                                    | [*operations.EndpointID](../../models/operations/endpointid.md) | :heavy_minus_sign:                                              | Filter by webhook endpoint ID.                                  |
| `page`                                                          | **int64*                                                        | :heavy_minus_sign:                                              | Page number, defaults to 1.                                     |
| `limit`                                                         | **int64*                                                        | :heavy_minus_sign:                                              | Size of a page, defaults to 10. Maximum is 100.                 |
| `opts`                                                          | [][operations.Option](../../models/operations/option.md)        | :heavy_minus_sign:                                              | The options for this request.                                   |

### Response

**[*operations.WebhooksListWebhookDeliveriesResponse](../../models/operations/webhookslistwebhookdeliveriesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## RedeliverWebhookEvent

Schedule the re-delivery of a webhook event.

**Scopes**: `webhooks:write`

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

    res, err := s.Webhooks.RedeliverWebhookEvent(ctx, "<value>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The webhook event ID.                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.WebhooksRedeliverWebhookEventResponse](../../models/operations/webhooksredeliverwebhookeventresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
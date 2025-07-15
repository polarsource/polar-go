# Checkouts
(*Checkouts*)

## Overview

### Available Operations

* [List](#list) - List Checkout Sessions
* [Create](#create) - Create Checkout Session
* [Get](#get) - Get Checkout Session
* [Update](#update) - Update Checkout Session
* [ClientGet](#clientget) - Get Checkout Session from Client
* [ClientUpdate](#clientupdate) - Update Checkout Session from Client
* [ClientConfirm](#clientconfirm) - Confirm Checkout Session from Client

## List

List checkout sessions.

**Scopes**: `checkouts:read` `checkouts:write`

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

    res, err := s.Checkouts.List(ctx, operations.CheckoutsListRequest{
        OrganizationID: polargo.Pointer(operations.CreateCheckoutsListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceCheckout != nil {
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CheckoutsListRequest](../../models/operations/checkoutslistrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CheckoutsListResponse](../../models/operations/checkoutslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a checkout session.

**Scopes**: `checkouts:write`

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

    res, err := s.Checkouts.Create(ctx, components.CheckoutCreate{
        CustomerBillingAddress: &components.Address{
            Country: "US",
        },
        Products: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Checkout != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.CheckoutCreate](../../models/components/checkoutcreate.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CheckoutsCreateResponse](../../models/operations/checkoutscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a checkout session by ID.

**Scopes**: `checkouts:read` `checkouts:write`

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

    res, err := s.Checkouts.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Checkout != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The checkout session ID.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CheckoutsGetResponse](../../models/operations/checkoutsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a checkout session.

**Scopes**: `checkouts:write`

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

    res, err := s.Checkouts.Update(ctx, "<value>", components.CheckoutUpdate{
        CustomerBillingAddress: &components.Address{
            Country: "US",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Checkout != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The checkout session ID.                                               |
| `checkoutUpdate`                                                       | [components.CheckoutUpdate](../../models/components/checkoutupdate.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CheckoutsUpdateResponse](../../models/operations/checkoutsupdateresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.CheckoutForbiddenError | 403                              | application/json                 |
| apierrors.ResourceNotFound       | 404                              | application/json                 |
| apierrors.HTTPValidationError    | 422                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |

## ClientGet

Get a checkout session by client secret.

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

    res, err := s.Checkouts.ClientGet(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutPublic != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `clientSecret`                                           | *string*                                                 | :heavy_check_mark:                                       | The checkout session client secret.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CheckoutsClientGetResponse](../../models/operations/checkoutsclientgetresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.ResourceNotFound     | 404                            | application/json               |
| apierrors.ExpiredCheckoutError | 410                            | application/json               |
| apierrors.HTTPValidationError  | 422                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## ClientUpdate

Update a checkout session by client secret.

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

    res, err := s.Checkouts.ClientUpdate(ctx, "<value>", components.CheckoutUpdatePublic{
        CustomerBillingAddress: &components.Address{
            Country: "US",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutPublic != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `clientSecret`                                                                     | *string*                                                                           | :heavy_check_mark:                                                                 | The checkout session client secret.                                                |
| `checkoutUpdatePublic`                                                             | [components.CheckoutUpdatePublic](../../models/components/checkoutupdatepublic.md) | :heavy_check_mark:                                                                 | N/A                                                                                |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CheckoutsClientUpdateResponse](../../models/operations/checkoutsclientupdateresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.CheckoutForbiddenError | 403                              | application/json                 |
| apierrors.ResourceNotFound       | 404                              | application/json                 |
| apierrors.ExpiredCheckoutError   | 410                              | application/json                 |
| apierrors.HTTPValidationError    | 422                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |

## ClientConfirm

Confirm a checkout session by client secret.

Orders and subscriptions will be processed.

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

    res, err := s.Checkouts.ClientConfirm(ctx, "<value>", components.CheckoutConfirmStripe{
        CustomerBillingAddress: &components.Address{
            Country: "US",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutPublicConfirmed != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `clientSecret`                                                                       | *string*                                                                             | :heavy_check_mark:                                                                   | The checkout session client secret.                                                  |
| `checkoutConfirmStripe`                                                              | [components.CheckoutConfirmStripe](../../models/components/checkoutconfirmstripe.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CheckoutsClientConfirmResponse](../../models/operations/checkoutsclientconfirmresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.PaymentError           | 400                              | application/json                 |
| apierrors.CheckoutForbiddenError | 403                              | application/json                 |
| apierrors.ResourceNotFound       | 404                              | application/json                 |
| apierrors.ExpiredCheckoutError   | 410                              | application/json                 |
| apierrors.HTTPValidationError    | 422                              | application/json                 |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |
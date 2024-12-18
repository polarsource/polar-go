# Checkouts
(*Checkouts*)

## Overview

### Available Operations

* [~~Create~~](#create) - Create Checkout :warning: **Deprecated** Use [Create](docs/sdks/custom/README.md#create) instead.
* [~~Get~~](#get) - Get Checkout :warning: **Deprecated**

## ~~Create~~

Create a checkout session.

> :warning: **DEPRECATED**: This API is deprecated. We recommend you to use the new custom checkout API, which is more flexible and powerful. Please refer to the documentation for more information.. Use `Create` instead.

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

    res, err := s.Checkouts.Create(ctx, components.CheckoutLegacyCreate{
        ProductPriceID: "<value>",
        SuccessURL: "https://probable-heating.com/",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutLegacy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.CheckoutLegacyCreate](../../models/components/checkoutlegacycreate.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CheckoutsCreateResponse](../../models/operations/checkoutscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ~~Get~~

Get an active checkout session by ID.

> :warning: **DEPRECATED**: This API is deprecated. We recommend you to use the new custom checkout API, which is more flexible and powerful. Please refer to the documentation for more information..

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

    res, err := s.Checkouts.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckoutLegacy != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CheckoutsGetResponse](../../models/operations/checkoutsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
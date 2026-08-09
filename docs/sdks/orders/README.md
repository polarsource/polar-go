# Orders

## Overview

### Available Operations

* [List](#list) - List Orders
* [Create](#create) - Create Order
* [Export](#export) - Export Orders
* [Get](#get) - Get Order
* [Update](#update) - Update Order
* [Finalize](#finalize) - Finalize Order
* [Invoice](#invoice) - Get Order Invoice
* [GenerateInvoice](#generateinvoice) - Generate Order Invoice
* [Receipt](#receipt) - Get Order Receipt

## List

List orders.

**Scopes**: `orders:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:list" method="get" path="/v1/orders/" -->
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

    res, err := s.Orders.List(ctx, operations.OrdersListRequest{
        OrganizationID: polargo.Pointer(operations.CreateOrdersListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceOrder != nil {
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.OrdersListRequest](../../models/operations/orderslistrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.OrdersListResponse](../../models/operations/orderslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a draft order for an off-session charge against a saved payment
method. The order is created with `status=draft` and no invoice number;
call `POST /v1/orders/{id}/finalize` to attempt the charge.

The organization must have the `off_session_charges_enabled` feature flag.

**Scopes**: `orders:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:create" method="post" path="/v1/orders/" -->
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

    res, err := s.Orders.Create(ctx, components.OrderCreate{
        OrganizationID: polargo.Pointer("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
        CustomerID: "<value>",
        ProductID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        switch res.Order.Discount.Type {
            case components.OrderDiscountTypeDiscountFixedOnceForeverDurationBase:
                // res.Order.Discount.DiscountFixedOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountFixedRepeatDurationBase:
                // res.Order.Discount.DiscountFixedRepeatDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageOnceForeverDurationBase:
                // res.Order.Discount.DiscountPercentageOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageRepeatDurationBase:
                // res.Order.Discount.DiscountPercentageRepeatDurationBase is populated
        }

    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.OrderCreate](../../models/components/ordercreate.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.OrdersCreateResponse](../../models/operations/orderscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Export

Export orders as a CSV file.

**Scopes**: `orders:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:export" method="get" path="/v1/orders/export" -->
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

    res, err := s.Orders.Export(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |
| `organizationID`                                                                                                                | [*operations.OrdersExportQueryParamOrganizationIDFilter](../../models/operations/ordersexportqueryparamorganizationidfilter.md) | :heavy_minus_sign:                                                                                                              | Filter by organization ID.                                                                                                      |
| `productID`                                                                                                                     | [*operations.OrdersExportQueryParamProductIDFilter](../../models/operations/ordersexportqueryparamproductidfilter.md)           | :heavy_minus_sign:                                                                                                              | Filter by product ID.                                                                                                           |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |

### Response

**[*operations.OrdersExportResponse](../../models/operations/ordersexportresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get an order by ID.

**Scopes**: `orders:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:get" method="get" path="/v1/orders/{id}" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"log"
	"github.com/polarsource/polar-go/models/components"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Orders.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        switch res.Order.Discount.Type {
            case components.OrderDiscountTypeDiscountFixedOnceForeverDurationBase:
                // res.Order.Discount.DiscountFixedOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountFixedRepeatDurationBase:
                // res.Order.Discount.DiscountFixedRepeatDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageOnceForeverDurationBase:
                // res.Order.Discount.DiscountPercentageOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageRepeatDurationBase:
                // res.Order.Discount.DiscountPercentageRepeatDurationBase is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The order ID.                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OrdersGetResponse](../../models/operations/ordersgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update an order.

**Scopes**: `orders:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:update" method="patch" path="/v1/orders/{id}" -->
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

    res, err := s.Orders.Update(ctx, "<value>", components.OrderUpdate{
        BillingAddress: &components.AddressInput{
            Country: components.AddressInputCountryAlpha2InputUs,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        switch res.Order.Discount.Type {
            case components.OrderDiscountTypeDiscountFixedOnceForeverDurationBase:
                // res.Order.Discount.DiscountFixedOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountFixedRepeatDurationBase:
                // res.Order.Discount.DiscountFixedRepeatDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageOnceForeverDurationBase:
                // res.Order.Discount.DiscountPercentageOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageRepeatDurationBase:
                // res.Order.Discount.DiscountPercentageRepeatDurationBase is populated
        }

    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `id`                                                             | `string`                                                         | :heavy_check_mark:                                               | The order ID.                                                    |
| `orderUpdate`                                                    | [components.OrderUpdate](../../models/components/orderupdate.md) | :heavy_check_mark:                                               | N/A                                                              |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.OrdersUpdateResponse](../../models/operations/ordersupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Finalize

Finalize a draft order and synchronously attempt an off-session charge.

On success, the order transitions to `paid` and benefit grants fire
before the response returns. On failure (decline, missing payment method,
SCA challenge), the order stays in `draft` and a 4xx error is returned.

The request fails with 412 if the order is not in `draft` status.

**Scopes**: `orders:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:finalize" method="post" path="/v1/orders/{id}/finalize" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"log"
	"github.com/polarsource/polar-go/models/components"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Orders.Finalize(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        switch res.Order.Discount.Type {
            case components.OrderDiscountTypeDiscountFixedOnceForeverDurationBase:
                // res.Order.Discount.DiscountFixedOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountFixedRepeatDurationBase:
                // res.Order.Discount.DiscountFixedRepeatDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageOnceForeverDurationBase:
                // res.Order.Discount.DiscountPercentageOnceForeverDurationBase is populated
            case components.OrderDiscountTypeDiscountPercentageRepeatDurationBase:
                // res.Order.Discount.DiscountPercentageRepeatDurationBase is populated
        }

    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `id`                                                                  | `string`                                                              | :heavy_check_mark:                                                    | The order ID.                                                         |
| `orderFinalize`                                                       | [*components.OrderFinalize](../../models/components/orderfinalize.md) | :heavy_minus_sign:                                                    | N/A                                                                   |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.OrdersFinalizeResponse](../../models/operations/ordersfinalizeresponse.md), error**

### Errors

| Error Type                                        | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| apierrors.OrdersFinalizeResponse402OrdersFinalize | 402                                               | application/json                                  |
| apierrors.OrdersFinalizeResponse403OrdersFinalize | 403                                               | application/json                                  |
| apierrors.ResourceNotFound                        | 404                                               | application/json                                  |
| apierrors.OrderNotDraft                           | 412                                               | application/json                                  |
| apierrors.HTTPValidationError                     | 422                                               | application/json                                  |
| apierrors.APIError                                | 4XX, 5XX                                          | \*/\*                                             |

## Invoice

Get an order's invoice data.

**Scopes**: `orders:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:invoice" method="get" path="/v1/orders/{id}/invoice" -->
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

    res, err := s.Orders.Invoice(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrderInvoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The order ID.                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OrdersInvoiceResponse](../../models/operations/ordersinvoiceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GenerateInvoice

Trigger generation of an order's invoice.

**Scopes**: `orders:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:generate_invoice" method="post" path="/v1/orders/{id}/invoice" -->
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

    res, err := s.Orders.GenerateInvoice(ctx, "<value>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The order ID.                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OrdersGenerateInvoiceResponse](../../models/operations/ordersgenerateinvoiceresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ResourceNotFound             | 404                                    | application/json                       |
| apierrors.OrderNotEligibleForInvoice   | 409                                    | application/json                       |
| apierrors.MissingInvoiceBillingDetails | 422                                    | application/json                       |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

## Receipt

Get a presigned URL to download an order's receipt PDF.

**Scopes**: `orders:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="orders:receipt" method="get" path="/v1/orders/{id}/receipt" -->
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

    res, err := s.Orders.Receipt(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrderReceipt != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The order ID.                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OrdersReceiptResponse](../../models/operations/ordersreceiptresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
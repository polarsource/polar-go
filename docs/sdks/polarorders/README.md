# PolarOrders
(*CustomerPortal.Orders*)

## Overview

### Available Operations

* [List](#list) - List Orders
* [Get](#get) - Get Order
* [Update](#update) - Update Order
* [GenerateInvoice](#generateinvoice) - Generate Order Invoice
* [Invoice](#invoice) - Get Order Invoice
* [GetPaymentStatus](#getpaymentstatus) - Get Order Payment Status
* [ConfirmRetryPayment](#confirmretrypayment) - Confirm Retry Payment

## List

List orders of the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:orders:list" method="get" path="/v1/customer-portal/orders/" -->
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

    res, err := s.CustomerPortal.Orders.List(ctx, operations.CustomerPortalOrdersListRequest{}, operations.CustomerPortalOrdersListSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceCustomerOrder != nil {
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.CustomerPortalOrdersListRequest](../../models/operations/customerportalorderslistrequest.md)   | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `security`                                                                                                 | [operations.CustomerPortalOrdersListSecurity](../../models/operations/customerportalorderslistsecurity.md) | :heavy_check_mark:                                                                                         | The security requirements to use for the request.                                                          |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CustomerPortalOrdersListResponse](../../models/operations/customerportalorderslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get an order by ID for the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:orders:get" method="get" path="/v1/customer-portal/orders/{id}" -->
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

    res, err := s.CustomerPortal.Orders.Get(ctx, operations.CustomerPortalOrdersGetSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `security`                                                                                               | [operations.CustomerPortalOrdersGetSecurity](../../models/operations/customerportalordersgetsecurity.md) | :heavy_check_mark:                                                                                       | The security requirements to use for the request.                                                        |
| `id`                                                                                                     | *string*                                                                                                 | :heavy_check_mark:                                                                                       | The order ID.                                                                                            |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.CustomerPortalOrdersGetResponse](../../models/operations/customerportalordersgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update an order for the authenticated customer.

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:orders:update" method="patch" path="/v1/customer-portal/orders/{id}" -->
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

    res, err := s.CustomerPortal.Orders.Update(ctx, operations.CustomerPortalOrdersUpdateSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>", components.CustomerOrderUpdate{
        BillingName: polargo.Pointer("<value>"),
        BillingAddress: &components.AddressInput{
            Country: components.CountryAlpha2InputUs,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerOrder != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `security`                                                                                                     | [operations.CustomerPortalOrdersUpdateSecurity](../../models/operations/customerportalordersupdatesecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |
| `id`                                                                                                           | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The order ID.                                                                                                  |
| `customerOrderUpdate`                                                                                          | [components.CustomerOrderUpdate](../../models/components/customerorderupdate.md)                               | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CustomerPortalOrdersUpdateResponse](../../models/operations/customerportalordersupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GenerateInvoice

Trigger generation of an order's invoice.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:orders:generate_invoice" method="post" path="/v1/customer-portal/orders/{id}/invoice" -->
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

    res, err := s.CustomerPortal.Orders.GenerateInvoice(ctx, operations.CustomerPortalOrdersGenerateInvoiceSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `security`                                                                                                                       | [operations.CustomerPortalOrdersGenerateInvoiceSecurity](../../models/operations/customerportalordersgenerateinvoicesecurity.md) | :heavy_check_mark:                                                                                                               | The security requirements to use for the request.                                                                                |
| `id`                                                                                                                             | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The order ID.                                                                                                                    |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.CustomerPortalOrdersGenerateInvoiceResponse](../../models/operations/customerportalordersgenerateinvoiceresponse.md), error**

### Errors

| Error Type                                                                                  | Status Code                                                                                 | Content Type                                                                                |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| apierrors.CustomerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoice | 422                                                                                         | application/json                                                                            |
| apierrors.APIError                                                                          | 4XX, 5XX                                                                                    | \*/\*                                                                                       |

## Invoice

Get an order's invoice data.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:orders:invoice" method="get" path="/v1/customer-portal/orders/{id}/invoice" -->
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

    res, err := s.CustomerPortal.Orders.Invoice(ctx, operations.CustomerPortalOrdersInvoiceSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerOrderInvoice != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `security`                                                                                                       | [operations.CustomerPortalOrdersInvoiceSecurity](../../models/operations/customerportalordersinvoicesecurity.md) | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `id`                                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The order ID.                                                                                                    |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CustomerPortalOrdersInvoiceResponse](../../models/operations/customerportalordersinvoiceresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetPaymentStatus

Get the current payment status for an order.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:orders:get_payment_status" method="get" path="/v1/customer-portal/orders/{id}/payment-status" -->
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

    res, err := s.CustomerPortal.Orders.GetPaymentStatus(ctx, operations.CustomerPortalOrdersGetPaymentStatusSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerOrderPaymentStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `security`                                                                                                                         | [operations.CustomerPortalOrdersGetPaymentStatusSecurity](../../models/operations/customerportalordersgetpaymentstatussecurity.md) | :heavy_check_mark:                                                                                                                 | The security requirements to use for the request.                                                                                  |
| `id`                                                                                                                               | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | The order ID.                                                                                                                      |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CustomerPortalOrdersGetPaymentStatusResponse](../../models/operations/customerportalordersgetpaymentstatusresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ConfirmRetryPayment

Confirm a retry payment using a Stripe confirmation token.

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:orders:confirm_retry_payment" method="post" path="/v1/customer-portal/orders/{id}/confirm-payment" -->
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

    res, err := s.CustomerPortal.Orders.ConfirmRetryPayment(ctx, operations.CustomerPortalOrdersConfirmRetryPaymentSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>", components.CustomerOrderConfirmPayment{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerOrderPaymentConfirmation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `security`                                                                                                                               | [operations.CustomerPortalOrdersConfirmRetryPaymentSecurity](../../models/operations/customerportalordersconfirmretrypaymentsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |
| `id`                                                                                                                                     | *string*                                                                                                                                 | :heavy_check_mark:                                                                                                                       | The order ID.                                                                                                                            |
| `customerOrderConfirmPayment`                                                                                                            | [components.CustomerOrderConfirmPayment](../../models/components/customerorderconfirmpayment.md)                                         | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.CustomerPortalOrdersConfirmRetryPaymentResponse](../../models/operations/customerportalordersconfirmretrypaymentresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ResourceNotFound         | 404                                | application/json                   |
| apierrors.PaymentAlreadyInProgress | 409                                | application/json                   |
| apierrors.OrderNotEligibleForRetry | 422                                | application/json                   |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |
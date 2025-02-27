# PolarCustomers
(*CustomerPortal.Customers*)

## Overview

### Available Operations

* [Get](#get) - Get Customer
* [Update](#update) - Update Customer
* [GetPaymentMethods](#getpaymentmethods) - Get Customer Payment Methods
* [AddPaymentMethod](#addpaymentmethod) - Add Customer Payment Method
* [DeletePaymentMethod](#deletepaymentmethod) - Delete Customer Payment Method

## Get

Get authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

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

    res, err := s.CustomerPortal.Customers.Get(ctx, operations.CustomerPortalCustomersGetSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerPortalCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `security`                                                                                                     | [operations.CustomerPortalCustomersGetSecurity](../../models/operations/customerportalcustomersgetsecurity.md) | :heavy_check_mark:                                                                                             | The security requirements to use for the request.                                                              |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CustomerPortalCustomersGetResponse](../../models/operations/customerportalcustomersgetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update authenticated customer.

**Scopes**: `customer_portal:write`

### Example Usage

```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"os"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.Customers.Update(ctx, components.CustomerPortalCustomerUpdate{}, operations.CustomerPortalCustomersUpdateSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerPortalCustomer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [components.CustomerPortalCustomerUpdate](../../models/components/customerportalcustomerupdate.md)                   | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.CustomerPortalCustomersUpdateSecurity](../../models/operations/customerportalcustomersupdatesecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CustomerPortalCustomersUpdateResponse](../../models/operations/customerportalcustomersupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetPaymentMethods

Get saved payment methods of the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

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

    res, err := s.CustomerPortal.Customers.GetPaymentMethods(ctx, operations.CustomerPortalCustomersGetPaymentMethodsSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceUnionPaymentMethodCardPaymentMethodGeneric != nil {
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `security`                                                                                                                                 | [operations.CustomerPortalCustomersGetPaymentMethodsSecurity](../../models/operations/customerportalcustomersgetpaymentmethodssecurity.md) | :heavy_check_mark:                                                                                                                         | The security requirements to use for the request.                                                                                          |
| `page`                                                                                                                                     | **int64*                                                                                                                                   | :heavy_minus_sign:                                                                                                                         | Page number, defaults to 1.                                                                                                                |
| `limit`                                                                                                                                    | **int64*                                                                                                                                   | :heavy_minus_sign:                                                                                                                         | Size of a page, defaults to 10. Maximum is 100.                                                                                            |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.CustomerPortalCustomersGetPaymentMethodsResponse](../../models/operations/customerportalcustomersgetpaymentmethodsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## AddPaymentMethod

Add a payment method to the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"os"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.Customers.AddPaymentMethod(ctx, components.CustomerPaymentMethodCreate{
        ConfirmationTokenID: "<id>",
        SetDefault: false,
        ReturnURL: "https://slight-digestive.info",
    }, operations.CustomerPortalCustomersAddPaymentMethodSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseCustomerPortalCustomersAddPaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [components.CustomerPaymentMethodCreate](../../models/components/customerpaymentmethodcreate.md)                                         | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `security`                                                                                                                               | [operations.CustomerPortalCustomersAddPaymentMethodSecurity](../../models/operations/customerportalcustomersaddpaymentmethodsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.CustomerPortalCustomersAddPaymentMethodResponse](../../models/operations/customerportalcustomersaddpaymentmethodresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## DeletePaymentMethod

Delete a payment method from the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

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

    res, err := s.CustomerPortal.Customers.DeletePaymentMethod(ctx, operations.CustomerPortalCustomersDeletePaymentMethodSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                      | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                                          | :heavy_check_mark:                                                                                                                             | The context to use for the request.                                                                                                            |
| `security`                                                                                                                                     | [operations.CustomerPortalCustomersDeletePaymentMethodSecurity](../../models/operations/customerportalcustomersdeletepaymentmethodsecurity.md) | :heavy_check_mark:                                                                                                                             | The security requirements to use for the request.                                                                                              |
| `id`                                                                                                                                           | *string*                                                                                                                                       | :heavy_check_mark:                                                                                                                             | N/A                                                                                                                                            |
| `opts`                                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                                       | :heavy_minus_sign:                                                                                                                             | The options for this request.                                                                                                                  |

### Response

**[*operations.CustomerPortalCustomersDeletePaymentMethodResponse](../../models/operations/customerportalcustomersdeletepaymentmethodresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
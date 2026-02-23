# Customers

## Overview

### Available Operations

* [List](#list) - List Customers
* [Create](#create) - Create Customer
* [Export](#export) - Export Customers
* [Get](#get) - Get Customer
* [Delete](#delete) - Delete Customer
* [Update](#update) - Update Customer
* [GetExternal](#getexternal) - Get Customer by External ID
* [DeleteExternal](#deleteexternal) - Delete Customer by External ID
* [UpdateExternal](#updateexternal) - Update Customer by External ID
* [GetState](#getstate) - Get Customer State
* [GetStateExternal](#getstateexternal) - Get Customer State by External ID

## List

List customers.

**Scopes**: `customers:read` `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:list" method="get" path="/v1/customers/" -->
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

    res, err := s.Customers.List(ctx, operations.CustomersListRequest{
        OrganizationID: polargo.Pointer(operations.CreateCustomersListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceCustomerWithMembers != nil {
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
| `request`                                                                          | [operations.CustomersListRequest](../../models/operations/customerslistrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CustomersListResponse](../../models/operations/customerslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a customer.

**Scopes**: `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:create" method="post" path="/v1/customers/" -->
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

    res, err := s.Customers.Create(ctx, components.CustomerCreate{
        ExternalID: polargo.Pointer("usr_1337"),
        Email: "customer@example.com",
        Name: polargo.Pointer("John Doe"),
        BillingAddress: &components.AddressInput{
            Country: components.CountryAlpha2InputUs,
        },
        TaxID: []*components.CustomerCreateTaxID{
            polargo.Pointer(components.CreateCustomerCreateTaxIDStr(
                "911144442",
            )),
            polargo.Pointer(components.CreateCustomerCreateTaxIDStr(
                "us_ein",
            )),
        },
        Locale: polargo.Pointer("en"),
        Type: components.CustomerTypeIndividual.ToPointer(),
        OrganizationID: polargo.Pointer("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
        Owner: &components.OwnerCreate{
            Email: polargo.Pointer("member@example.com"),
            Name: polargo.Pointer("Jane Doe"),
            ExternalID: polargo.Pointer("usr_1337"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerWithMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.CustomerCreate](../../models/components/customercreate.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CustomersCreateResponse](../../models/operations/customerscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Export

Export customers as a CSV file.

**Scopes**: `customers:read` `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:export" method="get" path="/v1/customers/export" -->
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

    res, err := s.Customers.Export(ctx, polargo.Pointer(operations.CreateCustomersExportQueryParamOrganizationIDStr(
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

| Parameter                                                                                                                 | Type                                                                                                                      | Required                                                                                                                  | Description                                                                                                               |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                     | :heavy_check_mark:                                                                                                        | The context to use for the request.                                                                                       |
| `organizationID`                                                                                                          | [*operations.CustomersExportQueryParamOrganizationID](../../models/operations/customersexportqueryparamorganizationid.md) | :heavy_minus_sign:                                                                                                        | Filter by organization ID.                                                                                                |
| `opts`                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                  | :heavy_minus_sign:                                                                                                        | The options for this request.                                                                                             |

### Response

**[*operations.CustomersExportResponse](../../models/operations/customersexportresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a customer by ID.

**Scopes**: `customers:read` `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:get" method="get" path="/v1/customers/{id}" -->
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

    res, err := s.Customers.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerWithMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The customer ID.                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersGetResponse](../../models/operations/customersgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a customer.

This action cannot be undone and will immediately:
- Cancel any active subscriptions for the customer
- Revoke all their benefits
- Clear any `external_id`

Use it only in the context of deleting a user within your
own service. Otherwise, use more granular API endpoints to cancel
a specific subscription or revoke certain benefits.

Note: The customers information will nonetheless be retained for historic
orders and subscriptions.

Set `anonymize=true` to also anonymize PII for GDPR compliance.

**Scopes**: `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:delete" method="delete" path="/v1/customers/{id}" -->
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

    res, err := s.Customers.Delete(ctx, "<value>", polargo.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                    |
| `id`                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                     | The customer ID.                                                                                                                                                                                                                                       |
| `anonymize`                                                                                                                                                                                                                                            | **bool*                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                     | If true, also anonymize the customer's personal data for GDPR compliance. This replaces email with a hashed version, hashes name and billing name (name preserved for businesses with tax_id), clears billing address, and removes OAuth account data. |
| `opts`                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                          |

### Response

**[*operations.CustomersDeleteResponse](../../models/operations/customersdeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a customer.

**Scopes**: `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:update" method="patch" path="/v1/customers/{id}" -->
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

    res, err := s.Customers.Update(ctx, "<value>", components.CustomerUpdate{
        Email: polargo.Pointer("customer@example.com"),
        Name: polargo.Pointer("John Doe"),
        BillingAddress: &components.AddressInput{
            Country: components.CountryAlpha2InputUs,
        },
        TaxID: []*components.CustomerUpdateTaxID{
            polargo.Pointer(components.CreateCustomerUpdateTaxIDStr(
                "911144442",
            )),
            polargo.Pointer(components.CreateCustomerUpdateTaxIDStr(
                "us_ein",
            )),
        },
        Locale: polargo.Pointer("en"),
        ExternalID: polargo.Pointer("usr_1337"),
        Type: components.CustomerTypeIndividual.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerWithMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The customer ID.                                                       |
| `customerUpdate`                                                       | [components.CustomerUpdate](../../models/components/customerupdate.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CustomersUpdateResponse](../../models/operations/customersupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetExternal

Get a customer by external ID.

**Scopes**: `customers:read` `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:get_external" method="get" path="/v1/customers/external/{external_id}" -->
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

    res, err := s.Customers.GetExternal(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerWithMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `externalID`                                             | *string*                                                 | :heavy_check_mark:                                       | The customer external ID.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersGetExternalResponse](../../models/operations/customersgetexternalresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## DeleteExternal

Delete a customer by external ID.

Immediately cancels any active subscriptions and revokes any active benefits.

Set `anonymize=true` to also anonymize PII for GDPR compliance.

**Scopes**: `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:delete_external" method="delete" path="/v1/customers/external/{external_id}" -->
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

    res, err := s.Customers.DeleteExternal(ctx, "<id>", polargo.Pointer(false))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `externalID`                                                              | *string*                                                                  | :heavy_check_mark:                                                        | The customer external ID.                                                 |
| `anonymize`                                                               | **bool*                                                                   | :heavy_minus_sign:                                                        | If true, also anonymize the customer's personal data for GDPR compliance. |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |

### Response

**[*operations.CustomersDeleteExternalResponse](../../models/operations/customersdeleteexternalresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## UpdateExternal

Update a customer by external ID.

**Scopes**: `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:update_external" method="patch" path="/v1/customers/external/{external_id}" -->
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

    res, err := s.Customers.UpdateExternal(ctx, "<id>", components.CustomerUpdateExternalID{
        Email: polargo.Pointer("customer@example.com"),
        Name: polargo.Pointer("John Doe"),
        BillingAddress: &components.AddressInput{
            Country: components.CountryAlpha2InputUs,
        },
        TaxID: []*components.CustomerUpdateExternalIDTaxID{
            polargo.Pointer(components.CreateCustomerUpdateExternalIDTaxIDStr(
                "911144442",
            )),
            polargo.Pointer(components.CreateCustomerUpdateExternalIDTaxIDStr(
                "us_ein",
            )),
        },
        Locale: polargo.Pointer("en"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerWithMembers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `externalID`                                                                               | *string*                                                                                   | :heavy_check_mark:                                                                         | The customer external ID.                                                                  |
| `customerUpdateExternalID`                                                                 | [components.CustomerUpdateExternalID](../../models/components/customerupdateexternalid.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CustomersUpdateExternalResponse](../../models/operations/customersupdateexternalresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetState

Get a customer state by ID.

The customer state includes information about
the customer's active subscriptions and benefits.

It's the ideal endpoint to use when you need to get a full overview
of a customer's status.

**Scopes**: `customers:read` `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:get_state" method="get" path="/v1/customers/{id}/state" -->
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

    res, err := s.Customers.GetState(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerState != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The customer ID.                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersGetStateResponse](../../models/operations/customersgetstateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetStateExternal

Get a customer state by external ID.

The customer state includes information about
the customer's active subscriptions and benefits.

It's the ideal endpoint to use when you need to get a full overview
of a customer's status.

**Scopes**: `customers:read` `customers:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:get_state_external" method="get" path="/v1/customers/external/{external_id}/state" -->
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

    res, err := s.Customers.GetStateExternal(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerState != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `externalID`                                             | *string*                                                 | :heavy_check_mark:                                       | The customer external ID.                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersGetStateExternalResponse](../../models/operations/customersgetstateexternalresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
# Discounts

## Overview

### Available Operations

* [List](#list) - List Discounts
* [Create](#create) - Create Discount
* [Get](#get) - Get Discount
* [Delete](#delete) - Delete Discount
* [Update](#update) - Update Discount

## List

List discounts.

**Scopes**: `discounts:read` `discounts:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="discounts:list" method="get" path="/v1/discounts/" -->
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

    res, err := s.Discounts.List(ctx, operations.DiscountsListRequest{
        OrganizationID: polargo.Pointer(operations.CreateDiscountsListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceDiscount != nil {
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
| `request`                                                                          | [operations.DiscountsListRequest](../../models/operations/discountslistrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.DiscountsListResponse](../../models/operations/discountslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a discount.

**Scopes**: `discounts:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="discounts:create" method="post" path="/v1/discounts/" -->
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

    res, err := s.Discounts.Create(ctx, components.CreateDiscountCreateDiscountPercentageOnceForeverDurationCreate(
        components.DiscountPercentageOnceForeverDurationCreate{
            Duration: components.DiscountDurationOnce,
            Type: components.DiscountTypeFixed,
            BasisPoints: 449604,
            Name: "<value>",
            OrganizationID: polargo.Pointer("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Discount != nil {
        switch res.Discount.Type {
            case components.DiscountUnionTypeDiscountFixedOnceForeverDuration:
                // res.Discount.DiscountFixedOnceForeverDuration is populated
            case components.DiscountUnionTypeDiscountFixedRepeatDuration:
                // res.Discount.DiscountFixedRepeatDuration is populated
            case components.DiscountUnionTypeDiscountPercentageOnceForeverDuration:
                // res.Discount.DiscountPercentageOnceForeverDuration is populated
            case components.DiscountUnionTypeDiscountPercentageRepeatDuration:
                // res.Discount.DiscountPercentageRepeatDuration is populated
        }

    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.DiscountCreate](../../models/components/discountcreate.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.DiscountsCreateResponse](../../models/operations/discountscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a discount by ID.

**Scopes**: `discounts:read` `discounts:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="discounts:get" method="get" path="/v1/discounts/{id}" -->
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

    res, err := s.Discounts.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Discount != nil {
        switch res.Discount.Type {
            case components.DiscountUnionTypeDiscountFixedOnceForeverDuration:
                // res.Discount.DiscountFixedOnceForeverDuration is populated
            case components.DiscountUnionTypeDiscountFixedRepeatDuration:
                // res.Discount.DiscountFixedRepeatDuration is populated
            case components.DiscountUnionTypeDiscountPercentageOnceForeverDuration:
                // res.Discount.DiscountPercentageOnceForeverDuration is populated
            case components.DiscountUnionTypeDiscountPercentageRepeatDuration:
                // res.Discount.DiscountPercentageRepeatDuration is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The discount ID.                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DiscountsGetResponse](../../models/operations/discountsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a discount.

**Scopes**: `discounts:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="discounts:delete" method="delete" path="/v1/discounts/{id}" -->
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

    res, err := s.Discounts.Delete(ctx, "<value>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The discount ID.                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DiscountsDeleteResponse](../../models/operations/discountsdeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a discount.

**Scopes**: `discounts:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="discounts:update" method="patch" path="/v1/discounts/{id}" -->
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

    res, err := s.Discounts.Update(ctx, "<value>", components.DiscountUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Discount != nil {
        switch res.Discount.Type {
            case components.DiscountUnionTypeDiscountFixedOnceForeverDuration:
                // res.Discount.DiscountFixedOnceForeverDuration is populated
            case components.DiscountUnionTypeDiscountFixedRepeatDuration:
                // res.Discount.DiscountFixedRepeatDuration is populated
            case components.DiscountUnionTypeDiscountPercentageOnceForeverDuration:
                // res.Discount.DiscountPercentageOnceForeverDuration is populated
            case components.DiscountUnionTypeDiscountPercentageRepeatDuration:
                // res.Discount.DiscountPercentageRepeatDuration is populated
        }

    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | The discount ID.                                                       |
| `discountUpdate`                                                       | [components.DiscountUpdate](../../models/components/discountupdate.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.DiscountsUpdateResponse](../../models/operations/discountsupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
# Customers.Members

## Overview

### Available Operations

* [Create](#create) - Create Member
* [CreateExternal](#createexternal) - Create Member by Customer External ID
* [Get](#get) - Get Member
* [Delete](#delete) - Delete Member
* [Update](#update) - Update Member
* [GetExternal](#getexternal) - Get Member by External ID
* [DeleteExternal](#deleteexternal) - Delete Member by External ID
* [UpdateExternal](#updateexternal) - Update Member by External ID

## Create

Create a new member for a customer.

Only B2B customers with the member management feature enabled can add members.
The authenticated user or organization must have access to the customer's organization.

**Scopes**: `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:create" method="post" path="/v1/customers/{id}/members" -->
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

    res, err := s.Customers.Members.Create(ctx, "<value>", components.MemberCreateFromCustomer{
        Email: "member@example.com",
        Name: polargo.Pointer("Jane Doe"),
        ExternalID: polargo.Pointer("usr_1337"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Member != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `id`                                                                                       | `string`                                                                                   | :heavy_check_mark:                                                                         | The customer ID.                                                                           |
| `memberCreateFromCustomer`                                                                 | [components.MemberCreateFromCustomer](../../models/components/membercreatefromcustomer.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CustomersMembersCreateResponse](../../models/operations/customersmemberscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotPermitted        | 403                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## CreateExternal

Create a new member for a customer identified by its external ID.

**Scopes**: `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:create_external" method="post" path="/v1/customers/external/{external_id}/members" -->
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

    res, err := s.Customers.Members.CreateExternal(ctx, "<id>", components.MemberCreateFromCustomer{
        Email: "member@example.com",
        Name: polargo.Pointer("Jane Doe"),
        ExternalID: polargo.Pointer("usr_1337"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Member != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `externalID`                                                                               | `string`                                                                                   | :heavy_check_mark:                                                                         | The customer external ID.                                                                  |
| `memberCreateFromCustomer`                                                                 | [components.MemberCreateFromCustomer](../../models/components/membercreatefromcustomer.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CustomersMembersCreateExternalResponse](../../models/operations/customersmemberscreateexternalresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.NotPermitted                | 403                                   | application/json                      |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.AmbiguousExternalCustomerID | 409                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## Get

Get a member of a customer by its ID.

**Scopes**: `members:read` `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:get" method="get" path="/v1/customers/{id}/members/{member_id}" -->
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

    res, err := s.Customers.Members.Get(ctx, "<value>", "a794a9c8-dc43-40b4-b2f5-ed16145e28ac")
    if err != nil {
        log.Fatal(err)
    }
    if res.Member != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The customer ID.                                         |
| `memberID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersMembersGetResponse](../../models/operations/customersmembersgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Delete

Delete a member of a customer.

**Scopes**: `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:delete" method="delete" path="/v1/customers/{id}/members/{member_id}" -->
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

    res, err := s.Customers.Members.Delete(ctx, "<value>", "a6d6f519-f76e-49a0-9868-b346c98100a6")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The customer ID.                                         |
| `memberID`                                               | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersMembersDeleteResponse](../../models/operations/customersmembersdeleteresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a member of a customer.

Only name, email and role can be updated.

**Scopes**: `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:update" method="patch" path="/v1/customers/{id}/members/{member_id}" -->
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

    res, err := s.Customers.Members.Update(ctx, "<value>", "f48ea05d-6a60-4bb1-b3d9-4b3cd7194f3a", components.MemberUpdate{
        Name: polargo.Pointer("Jane Doe"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Member != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `id`                                                               | `string`                                                           | :heavy_check_mark:                                                 | The customer ID.                                                   |
| `memberID`                                                         | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `memberUpdate`                                                     | [components.MemberUpdate](../../models/components/memberupdate.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CustomersMembersUpdateResponse](../../models/operations/customersmembersupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetExternal

Get a member by external ID for a customer identified by its external ID.

**Scopes**: `members:read` `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:get_external" method="get" path="/v1/customers/external/{external_id}/members/{member_external_id}" -->
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

    res, err := s.Customers.Members.GetExternal(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Member != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `externalID`                                             | `string`                                                 | :heavy_check_mark:                                       | The customer external ID.                                |
| `memberExternalID`                                       | `string`                                                 | :heavy_check_mark:                                       | The member external ID.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersMembersGetExternalResponse](../../models/operations/customersmembersgetexternalresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.AmbiguousExternalCustomerID | 409                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## DeleteExternal

Delete a member by external ID for a customer identified by its external ID.

**Scopes**: `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:delete_external" method="delete" path="/v1/customers/external/{external_id}/members/{member_external_id}" -->
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

    res, err := s.Customers.Members.DeleteExternal(ctx, "<id>", "<id>")
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
| `externalID`                                             | `string`                                                 | :heavy_check_mark:                                       | The customer external ID.                                |
| `memberExternalID`                                       | `string`                                                 | :heavy_check_mark:                                       | The member external ID.                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomersMembersDeleteExternalResponse](../../models/operations/customersmembersdeleteexternalresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.AmbiguousExternalCustomerID | 409                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |

## UpdateExternal

Update a member by external ID for a customer identified by its external ID.

**Scopes**: `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customers:members:update_external" method="patch" path="/v1/customers/external/{external_id}/members/{member_external_id}" -->
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

    res, err := s.Customers.Members.UpdateExternal(ctx, "<id>", "<id>", components.MemberUpdate{
        Name: polargo.Pointer("Jane Doe"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Member != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `externalID`                                                       | `string`                                                           | :heavy_check_mark:                                                 | The customer external ID.                                          |
| `memberExternalID`                                                 | `string`                                                           | :heavy_check_mark:                                                 | The member external ID.                                            |
| `memberUpdate`                                                     | [components.MemberUpdate](../../models/components/memberupdate.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CustomersMembersUpdateExternalResponse](../../models/operations/customersmembersupdateexternalresponse.md), error**

### Errors

| Error Type                            | Status Code                           | Content Type                          |
| ------------------------------------- | ------------------------------------- | ------------------------------------- |
| apierrors.ResourceNotFound            | 404                                   | application/json                      |
| apierrors.AmbiguousExternalCustomerID | 409                                   | application/json                      |
| apierrors.HTTPValidationError         | 422                                   | application/json                      |
| apierrors.APIError                    | 4XX, 5XX                              | \*/\*                                 |
# CustomFields
(*CustomFields*)

## Overview

### Available Operations

* [List](#list) - List Custom Fields
* [Create](#create) - Create Custom Field
* [Get](#get) - Get Custom Field
* [Update](#update) - Update Custom Field
* [Delete](#delete) - Delete Custom Field

## List

List custom fields.

**Scopes**: `custom_fields:read` `custom_fields:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="custom-fields:list" method="get" path="/v1/custom-fields/" -->
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

    res, err := s.CustomFields.List(ctx, operations.CustomFieldsListRequest{
        OrganizationID: polargo.Pointer(operations.CreateCustomFieldsListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceCustomField != nil {
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CustomFieldsListRequest](../../models/operations/customfieldslistrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CustomFieldsListResponse](../../models/operations/customfieldslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a custom field.

**Scopes**: `custom_fields:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="custom-fields:create" method="post" path="/v1/custom-fields/" -->
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

    res, err := s.CustomFields.Create(ctx, components.CreateCustomFieldCreateSelect(
        components.CustomFieldCreateSelect{
            Slug: "<value>",
            Name: "<value>",
            OrganizationID: polargo.Pointer("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
            Properties: components.CustomFieldSelectProperties{
                Options: []components.CustomFieldSelectOption{},
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomField != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.CustomFieldCreate](../../models/components/customfieldcreate.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CustomFieldsCreateResponse](../../models/operations/customfieldscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a custom field by ID.

**Scopes**: `custom_fields:read` `custom_fields:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="custom-fields:get" method="get" path="/v1/custom-fields/{id}" -->
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

    res, err := s.CustomFields.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomField != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The custom field ID.                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomFieldsGetResponse](../../models/operations/customfieldsgetresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.PolarExceptionsResourceNotFound | 404                                       | application/json                          |
| apierrors.HTTPValidationError             | 422                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## Update

Update a custom field.

**Scopes**: `custom_fields:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="custom-fields:update" method="patch" path="/v1/custom-fields/{id}" -->
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

    res, err := s.CustomFields.Update(ctx, "<value>", components.CreateCustomFieldUpdateDate(
        components.CustomFieldUpdateDate{},
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomField != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The custom field ID.                                                         |
| `customFieldUpdate`                                                          | [components.CustomFieldUpdate](../../models/components/customfieldupdate.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CustomFieldsUpdateResponse](../../models/operations/customfieldsupdateresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.PolarExceptionsResourceNotFound | 404                                       | application/json                          |
| apierrors.HTTPValidationError             | 422                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |

## Delete

Delete a custom field.

**Scopes**: `custom_fields:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="custom-fields:delete" method="delete" path="/v1/custom-fields/{id}" -->
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

    res, err := s.CustomFields.Delete(ctx, "<value>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The custom field ID.                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomFieldsDeleteResponse](../../models/operations/customfieldsdeleteresponse.md), error**

### Errors

| Error Type                                | Status Code                               | Content Type                              |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| apierrors.PolarExceptionsResourceNotFound | 404                                       | application/json                          |
| apierrors.HTTPValidationError             | 422                                       | application/json                          |
| apierrors.APIError                        | 4XX, 5XX                                  | \*/\*                                     |
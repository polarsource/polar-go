# LicenseKeys
(*LicenseKeys*)

## Overview

### Available Operations

* [List](#list) - List License Keys
* [Get](#get) - Get License Key
* [Update](#update) - Update License Key
* [GetActivation](#getactivation) - Get Activation
* [Validate](#validate) - Validate License Key
* [Activate](#activate) - Activate License Key
* [Deactivate](#deactivate) - Deactivate License Key

## List

Get license keys connected to the given organization & filters.

**Scopes**: `license_keys:read` `license_keys:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="license_keys:list" method="get" path="/v1/license-keys/" -->
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

    res, err := s.LicenseKeys.List(ctx, polargo.Pointer(operations.CreateLicenseKeysListQueryParamOrganizationIDFilterStr(
        "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
    )), nil, polargo.Int64(1), polargo.Int64(10))
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceLicenseKeyRead != nil {
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

| Parameter                                                                                                                             | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                 | :heavy_check_mark:                                                                                                                    | The context to use for the request.                                                                                                   |
| `organizationID`                                                                                                                      | [*operations.LicenseKeysListQueryParamOrganizationIDFilter](../../models/operations/licensekeyslistqueryparamorganizationidfilter.md) | :heavy_minus_sign:                                                                                                                    | Filter by organization ID.                                                                                                            |
| `benefitID`                                                                                                                           | [*operations.QueryParamBenefitIDFilter](../../models/operations/queryparambenefitidfilter.md)                                         | :heavy_minus_sign:                                                                                                                    | Filter by benefit ID.                                                                                                                 |
| `page`                                                                                                                                | **int64*                                                                                                                              | :heavy_minus_sign:                                                                                                                    | Page number, defaults to 1.                                                                                                           |
| `limit`                                                                                                                               | **int64*                                                                                                                              | :heavy_minus_sign:                                                                                                                    | Size of a page, defaults to 10. Maximum is 100.                                                                                       |
| `opts`                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                              | :heavy_minus_sign:                                                                                                                    | The options for this request.                                                                                                         |

### Response

**[*operations.LicenseKeysListResponse](../../models/operations/licensekeyslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.Unauthorized        | 401                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a license key.

**Scopes**: `license_keys:read` `license_keys:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="license_keys:get" method="get" path="/v1/license-keys/{id}" -->
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

    res, err := s.LicenseKeys.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LicenseKeyWithActivations != nil {
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

**[*operations.LicenseKeysGetResponse](../../models/operations/licensekeysgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.Unauthorized        | 401                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a license key.

**Scopes**: `license_keys:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="license_keys:update" method="patch" path="/v1/license-keys/{id}" -->
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

    res, err := s.LicenseKeys.Update(ctx, "<value>", components.LicenseKeyUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.LicenseKeyRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `licenseKeyUpdate`                                                         | [components.LicenseKeyUpdate](../../models/components/licensekeyupdate.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.LicenseKeysUpdateResponse](../../models/operations/licensekeysupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.Unauthorized        | 401                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetActivation

Get a license key activation.

**Scopes**: `license_keys:read` `license_keys:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="license_keys:get_activation" method="get" path="/v1/license-keys/{id}/activations/{activation_id}" -->
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

    res, err := s.LicenseKeys.GetActivation(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LicenseKeyActivationRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `activationID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LicenseKeysGetActivationResponse](../../models/operations/licensekeysgetactivationresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.Unauthorized        | 401                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Validate

Validate a license key.

**Scopes**: `license_keys:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="license_keys:validate" method="post" path="/v1/license-keys/validate" -->
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

    res, err := s.LicenseKeys.Validate(ctx, components.LicenseKeyValidate{
        Key: "<key>",
        OrganizationID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidatedLicenseKey != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.LicenseKeyValidate](../../models/components/licensekeyvalidate.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.LicenseKeysValidateResponse](../../models/operations/licensekeysvalidateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Activate

Activate a license key instance.

**Scopes**: `license_keys:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="license_keys:activate" method="post" path="/v1/license-keys/activate" -->
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

    res, err := s.LicenseKeys.Activate(ctx, components.LicenseKeyActivate{
        Key: "<key>",
        OrganizationID: "<value>",
        Label: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LicenseKeyActivationRead != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.LicenseKeyActivate](../../models/components/licensekeyactivate.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.LicenseKeysActivateResponse](../../models/operations/licensekeysactivateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotPermitted        | 403                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Deactivate

Deactivate a license key instance.

**Scopes**: `license_keys:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="license_keys:deactivate" method="post" path="/v1/license-keys/deactivate" -->
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

    res, err := s.LicenseKeys.Deactivate(ctx, components.LicenseKeyDeactivate{
        Key: "<key>",
        OrganizationID: "<value>",
        ActivationID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.LicenseKeyDeactivate](../../models/components/licensekeydeactivate.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.LicenseKeysDeactivateResponse](../../models/operations/licensekeysdeactivateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
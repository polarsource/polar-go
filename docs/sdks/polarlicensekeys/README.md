# PolarLicenseKeys
(*CustomerPortal.LicenseKeys*)

## Overview

### Available Operations

* [List](#list) - List License Keys
* [Get](#get) - Get License Key
* [Validate](#validate) - Validate License Key
* [Activate](#activate) - Activate License Key
* [Deactivate](#deactivate) - Deactivate License Key

## List

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:license_keys:list" method="get" path="/v1/customer-portal/license-keys/" -->
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

    res, err := s.CustomerPortal.LicenseKeys.List(ctx, operations.CustomerPortalLicenseKeysListSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, polargo.Pointer(operations.CreateCustomerPortalLicenseKeysListQueryParamOrganizationIDFilterStr(
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

| Parameter                                                                                                                                                         | Type                                                                                                                                                              | Required                                                                                                                                                          | Description                                                                                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                             | :heavy_check_mark:                                                                                                                                                | The context to use for the request.                                                                                                                               |
| `security`                                                                                                                                                        | [operations.CustomerPortalLicenseKeysListSecurity](../../models/operations/customerportallicensekeyslistsecurity.md)                                              | :heavy_check_mark:                                                                                                                                                | The security requirements to use for the request.                                                                                                                 |
| `organizationID`                                                                                                                                                  | [*operations.CustomerPortalLicenseKeysListQueryParamOrganizationIDFilter](../../models/operations/customerportallicensekeyslistqueryparamorganizationidfilter.md) | :heavy_minus_sign:                                                                                                                                                | Filter by organization ID.                                                                                                                                        |
| `benefitID`                                                                                                                                                       | **string*                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                | Filter by a specific benefit                                                                                                                                      |
| `page`                                                                                                                                                            | **int64*                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                | Page number, defaults to 1.                                                                                                                                       |
| `limit`                                                                                                                                                           | **int64*                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                | Size of a page, defaults to 10. Maximum is 100.                                                                                                                   |
| `opts`                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                          | :heavy_minus_sign:                                                                                                                                                | The options for this request.                                                                                                                                     |

### Response

**[*operations.CustomerPortalLicenseKeysListResponse](../../models/operations/customerportallicensekeyslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.Unauthorized        | 401                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a license key.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:license_keys:get" method="get" path="/v1/customer-portal/license-keys/{id}" -->
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

    res, err := s.CustomerPortal.LicenseKeys.Get(ctx, operations.CustomerPortalLicenseKeysGetSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LicenseKeyWithActivations != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `security`                                                                                                         | [operations.CustomerPortalLicenseKeysGetSecurity](../../models/operations/customerportallicensekeysgetsecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |
| `id`                                                                                                               | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CustomerPortalLicenseKeysGetResponse](../../models/operations/customerportallicensekeysgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Validate

Validate a license key.

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:license_keys:validate" method="post" path="/v1/customer-portal/license-keys/validate" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.LicenseKeys.Validate(ctx, components.LicenseKeyValidate{
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

**[*operations.CustomerPortalLicenseKeysValidateResponse](../../models/operations/customerportallicensekeysvalidateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Activate

Activate a license key instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:license_keys:activate" method="post" path="/v1/customer-portal/license-keys/activate" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.LicenseKeys.Activate(ctx, components.LicenseKeyActivate{
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

**[*operations.CustomerPortalLicenseKeysActivateResponse](../../models/operations/customerportallicensekeysactivateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotPermitted        | 403                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Deactivate

Deactivate a license key instance.

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:license_keys:deactivate" method="post" path="/v1/customer-portal/license-keys/deactivate" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.LicenseKeys.Deactivate(ctx, components.LicenseKeyDeactivate{
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

**[*operations.CustomerPortalLicenseKeysDeactivateResponse](../../models/operations/customerportallicensekeysdeactivateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
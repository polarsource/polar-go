# CustomerPortal.BenefitGrants

## Overview

### Available Operations

* [List](#list) - List Benefit Grants
* [Get](#get) - Get Benefit Grant
* [Update](#update) - Update Benefit Grant

## List

List benefits grants of the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:benefit-grants:list" method="get" path="/v1/customer-portal/benefit-grants/" -->
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

    res, err := s.CustomerPortal.BenefitGrants.List(ctx, operations.CustomerPortalBenefitGrantsListRequest{}, operations.CustomerPortalBenefitGrantsListSecurity{
        CustomerSession: polargo.Pointer(os.Getenv("POLAR_CUSTOMER_SESSION")),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceCustomerBenefitGrant != nil {
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

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.CustomerPortalBenefitGrantsListRequest](../../models/operations/customerportalbenefitgrantslistrequest.md)   | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `security`                                                                                                               | [operations.CustomerPortalBenefitGrantsListSecurity](../../models/operations/customerportalbenefitgrantslistsecurity.md) | :heavy_check_mark:                                                                                                       | The security requirements to use for the request.                                                                        |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.CustomerPortalBenefitGrantsListResponse](../../models/operations/customerportalbenefitgrantslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a benefit grant by ID for the authenticated customer.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:benefit-grants:get" method="get" path="/v1/customer-portal/benefit-grants/{id}" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"os"
	"github.com/polarsource/polar-go/models/operations"
	"log"
	"github.com/polarsource/polar-go/models/components"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.BenefitGrants.Get(ctx, operations.CustomerPortalBenefitGrantsGetSecurity{
        CustomerSession: polargo.Pointer(os.Getenv("POLAR_CUSTOMER_SESSION")),
    }, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerBenefitGrant != nil {
        switch res.CustomerBenefitGrant.Type {
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantDiscord:
                // res.CustomerBenefitGrant.CustomerBenefitGrantDiscord is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantGitHubRepository:
                // res.CustomerBenefitGrant.CustomerBenefitGrantGitHubRepository is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantDownloadables:
                // res.CustomerBenefitGrant.CustomerBenefitGrantDownloadables is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantLicenseKeys:
                // res.CustomerBenefitGrant.CustomerBenefitGrantLicenseKeys is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantCustom:
                // res.CustomerBenefitGrant.CustomerBenefitGrantCustom is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantMeterCredit:
                // res.CustomerBenefitGrant.CustomerBenefitGrantMeterCredit is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantFeatureFlag:
                // res.CustomerBenefitGrant.CustomerBenefitGrantFeatureFlag is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `security`                                                                                                             | [operations.CustomerPortalBenefitGrantsGetSecurity](../../models/operations/customerportalbenefitgrantsgetsecurity.md) | :heavy_check_mark:                                                                                                     | The security requirements to use for the request.                                                                      |
| `id`                                                                                                                   | `string`                                                                                                               | :heavy_check_mark:                                                                                                     | The benefit grant ID.                                                                                                  |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CustomerPortalBenefitGrantsGetResponse](../../models/operations/customerportalbenefitgrantsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a benefit grant for the authenticated customer.

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:benefit-grants:update" method="patch" path="/v1/customer-portal/benefit-grants/{id}" -->
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

    res, err := s.CustomerPortal.BenefitGrants.Update(ctx, operations.CustomerPortalBenefitGrantsUpdateSecurity{
        CustomerSession: polargo.Pointer(os.Getenv("POLAR_CUSTOMER_SESSION")),
    }, "<value>", components.CreateCustomerBenefitGrantUpdateLicenseKeys(
        components.CustomerBenefitGrantLicenseKeysUpdate{},
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerBenefitGrant != nil {
        switch res.CustomerBenefitGrant.Type {
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantDiscord:
                // res.CustomerBenefitGrant.CustomerBenefitGrantDiscord is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantGitHubRepository:
                // res.CustomerBenefitGrant.CustomerBenefitGrantGitHubRepository is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantDownloadables:
                // res.CustomerBenefitGrant.CustomerBenefitGrantDownloadables is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantLicenseKeys:
                // res.CustomerBenefitGrant.CustomerBenefitGrantLicenseKeys is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantCustom:
                // res.CustomerBenefitGrant.CustomerBenefitGrantCustom is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantMeterCredit:
                // res.CustomerBenefitGrant.CustomerBenefitGrantMeterCredit is populated
            case components.CustomerBenefitGrantTypeCustomerBenefitGrantFeatureFlag:
                // res.CustomerBenefitGrant.CustomerBenefitGrantFeatureFlag is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `security`                                                                                                                   | [operations.CustomerPortalBenefitGrantsUpdateSecurity](../../models/operations/customerportalbenefitgrantsupdatesecurity.md) | :heavy_check_mark:                                                                                                           | The security requirements to use for the request.                                                                            |
| `id`                                                                                                                         | `string`                                                                                                                     | :heavy_check_mark:                                                                                                           | The benefit grant ID.                                                                                                        |
| `customerBenefitGrantUpdate`                                                                                                 | [components.CustomerBenefitGrantUpdate](../../models/components/customerbenefitgrantupdate.md)                               | :heavy_check_mark:                                                                                                           | N/A                                                                                                                          |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.CustomerPortalBenefitGrantsUpdateResponse](../../models/operations/customerportalbenefitgrantsupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.NotPermitted        | 403                           | application/json              |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
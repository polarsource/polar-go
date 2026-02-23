# Meters

## Overview

### Available Operations

* [List](#list) - List Meters
* [Create](#create) - Create Meter
* [Get](#get) - Get Meter
* [Update](#update) - Update Meter
* [Quantities](#quantities) - Get Meter Quantities

## List

List meters.

**Scopes**: `meters:read` `meters:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="meters:list" method="get" path="/v1/meters/" -->
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

    res, err := s.Meters.List(ctx, operations.MetersListRequest{
        OrganizationID: polargo.Pointer(operations.CreateMetersListQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceMeter != nil {
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
| `request`                                                                    | [operations.MetersListRequest](../../models/operations/meterslistrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.MetersListResponse](../../models/operations/meterslistresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Create

Create a meter.

**Scopes**: `meters:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="meters:create" method="post" path="/v1/meters/" -->
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

    res, err := s.Meters.Create(ctx, components.MeterCreate{
        Name: "<value>",
        Filter: components.Filter{
            Conjunction: components.FilterConjunctionOr,
            Clauses: []components.Clauses{
                components.CreateClausesFilterClause(
                    components.FilterClause{
                        Property: "<value>",
                        Operator: components.FilterOperatorNe,
                        Value: components.CreateValueStr(
                            "<value>",
                        ),
                    },
                ),
            },
        },
        Aggregation: components.CreateMeterCreateAggregationAvg(
            components.PropertyAggregation{
                Func: components.FuncMax,
                Property: "<value>",
            },
        ),
        OrganizationID: polargo.Pointer("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Meter != nil {
        switch res.Meter.Aggregation.Type {
            case components.MeterAggregationTypeAvg:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeCount:
                // res.Meter.Aggregation.CountAggregation is populated
            case components.MeterAggregationTypeMax:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeMin:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeSum:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeUnique:
                // res.Meter.Aggregation.UniqueAggregation is populated
        }

    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.MeterCreate](../../models/components/metercreate.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.MetersCreateResponse](../../models/operations/meterscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Get

Get a meter by ID.

**Scopes**: `meters:read` `meters:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="meters:get" method="get" path="/v1/meters/{id}" -->
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

    res, err := s.Meters.Get(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Meter != nil {
        switch res.Meter.Aggregation.Type {
            case components.MeterAggregationTypeAvg:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeCount:
                // res.Meter.Aggregation.CountAggregation is populated
            case components.MeterAggregationTypeMax:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeMin:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeSum:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeUnique:
                // res.Meter.Aggregation.UniqueAggregation is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The meter ID.                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MetersGetResponse](../../models/operations/metersgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Update

Update a meter.

**Scopes**: `meters:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="meters:update" method="patch" path="/v1/meters/{id}" -->
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

    res, err := s.Meters.Update(ctx, "<value>", components.MeterUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Meter != nil {
        switch res.Meter.Aggregation.Type {
            case components.MeterAggregationTypeAvg:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeCount:
                // res.Meter.Aggregation.CountAggregation is populated
            case components.MeterAggregationTypeMax:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeMin:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeSum:
                // res.Meter.Aggregation.PropertyAggregation is populated
            case components.MeterAggregationTypeUnique:
                // res.Meter.Aggregation.UniqueAggregation is populated
        }

    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `id`                                                             | *string*                                                         | :heavy_check_mark:                                               | The meter ID.                                                    |
| `meterUpdate`                                                    | [components.MeterUpdate](../../models/components/meterupdate.md) | :heavy_check_mark:                                               | N/A                                                              |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.MetersUpdateResponse](../../models/operations/metersupdateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Quantities

Get quantities of a meter over a time period.

**Scopes**: `meters:read` `meters:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="meters:quantities" method="get" path="/v1/meters/{id}/quantities" -->
```go
package main

import(
	"context"
	"os"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/types"
	"github.com/polarsource/polar-go/models/components"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New(
        polargo.WithSecurity(os.Getenv("POLAR_ACCESS_TOKEN")),
    )

    res, err := s.Meters.Quantities(ctx, operations.MetersQuantitiesRequest{
        ID: "<value>",
        StartTimestamp: types.MustTimeFromString("2025-11-25T04:37:16.823Z"),
        EndTimestamp: types.MustTimeFromString("2025-11-26T17:06:00.727Z"),
        Interval: components.TimeIntervalDay,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeterQuantities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.MetersQuantitiesRequest](../../models/operations/metersquantitiesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.MetersQuantitiesResponse](../../models/operations/metersquantitiesresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.ResourceNotFound    | 404                           | application/json              |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
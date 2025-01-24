# Metrics
(*Metrics*)

## Overview

### Available Operations

* [Get](#get) - Get Metrics
* [Limits](#limits) - Get Metrics Limits

## Get

Get metrics about your orders and subscriptions.

### Example Usage

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

    res, err := s.Metrics.Get(ctx, operations.MetricsGetRequest{
        StartDate: types.MustDateFromString("2025-02-06"),
        EndDate: types.MustDateFromString("2024-09-04"),
        Interval: components.IntervalWeek,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.MetricsGetRequest](../../models/operations/metricsgetrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.MetricsGetResponse](../../models/operations/metricsgetresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## Limits

Get the interval limits for the metrics endpoint.

### Example Usage

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

    res, err := s.Metrics.Limits(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsLimits != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MetricsLimitsResponse](../../models/operations/metricslimitsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |
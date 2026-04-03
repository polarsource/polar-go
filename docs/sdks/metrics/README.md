# Metrics

## Overview

### Available Operations

* [Get](#get) - Get Metrics
* [Limits](#limits) - Get Metrics Limits
* [ListDashboards](#listdashboards) - List Metric Dashboards
* [CreateDashboard](#createdashboard) - Create Metric Dashboard
* [GetDashboard](#getdashboard) - Get Metric Dashboard
* [DeleteDashboard](#deletedashboard) - Delete Metric Dashboard
* [UpdateDashboard](#updatedashboard) - Update Metric Dashboard

## Get

Get metrics about your orders and subscriptions.

Currency values are output in cents.

**Scopes**: `metrics:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="metrics:get" method="get" path="/v1/metrics/" -->
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
        StartDate: types.MustDateFromString("2025-03-14"),
        EndDate: types.MustDateFromString("2025-03-18"),
        Interval: components.TimeIntervalHour,
        OrganizationID: polargo.Pointer(operations.CreateMetricsGetQueryParamOrganizationIDFilterStr(
            "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
        )),
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

**Scopes**: `metrics:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="metrics:limits" method="get" path="/v1/metrics/limits" -->
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

## ListDashboards

List user-defined metric dashboards.

**Scopes**: `metrics:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="metrics:list_dashboards" method="get" path="/v1/metrics/dashboards" -->
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

    res, err := s.Metrics.ListDashboards(ctx, polargo.Pointer(operations.CreateMetricsListDashboardsQueryParamOrganizationIDFilterStr(
        "1dbfc517-0bbf-4301-9ba8-555ca42b9737",
    )))
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseMetricsListDashboards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                         | Type                                                                                                                                              | Required                                                                                                                                          | Description                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                             | :heavy_check_mark:                                                                                                                                | The context to use for the request.                                                                                                               |
| `organizationID`                                                                                                                                  | [*operations.MetricsListDashboardsQueryParamOrganizationIDFilter](../../models/operations/metricslistdashboardsqueryparamorganizationidfilter.md) | :heavy_minus_sign:                                                                                                                                | Filter by organization ID.                                                                                                                        |
| `opts`                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                          | :heavy_minus_sign:                                                                                                                                | The options for this request.                                                                                                                     |

### Response

**[*operations.MetricsListDashboardsResponse](../../models/operations/metricslistdashboardsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## CreateDashboard

Create a user-defined metric dashboard.

**Scopes**: `metrics:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="metrics:create_dashboard" method="post" path="/v1/metrics/dashboards" -->
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

    res, err := s.Metrics.CreateDashboard(ctx, components.MetricDashboardCreate{
        Name: "<value>",
        OrganizationID: polargo.Pointer("1dbfc517-0bbf-4301-9ba8-555ca42b9737"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricDashboardSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.MetricDashboardCreate](../../models/components/metricdashboardcreate.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.MetricsCreateDashboardResponse](../../models/operations/metricscreatedashboardresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetDashboard

Get a user-defined metric dashboard by ID.

**Scopes**: `metrics:read`

### Example Usage

<!-- UsageSnippet language="go" operationID="metrics:get_dashboard" method="get" path="/v1/metrics/dashboards/{id}" -->
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

    res, err := s.Metrics.GetDashboard(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricDashboardSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The metric dashboard ID.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MetricsGetDashboardResponse](../../models/operations/metricsgetdashboardresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## DeleteDashboard

Delete a user-defined metric dashboard.

**Scopes**: `metrics:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="metrics:delete_dashboard" method="delete" path="/v1/metrics/dashboards/{id}" -->
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

    res, err := s.Metrics.DeleteDashboard(ctx, "<value>")
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
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The metric dashboard ID.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MetricsDeleteDashboardResponse](../../models/operations/metricsdeletedashboardresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## UpdateDashboard

Update a user-defined metric dashboard.

**Scopes**: `metrics:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="metrics:update_dashboard" method="patch" path="/v1/metrics/dashboards/{id}" -->
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

    res, err := s.Metrics.UpdateDashboard(ctx, "<value>", components.MetricDashboardUpdate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricDashboardSchema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | `string`                                                                             | :heavy_check_mark:                                                                   | The metric dashboard ID.                                                             |
| `metricDashboardUpdate`                                                              | [components.MetricDashboardUpdate](../../models/components/metricdashboardupdate.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.MetricsUpdateDashboardResponse](../../models/operations/metricsupdatedashboardresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
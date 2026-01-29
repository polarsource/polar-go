# MemberSessions

## Overview

### Available Operations

* [Create](#create) - Create Member Session

## Create

Create a member session.

This endpoint is only available for organizations with `member_model_enabled`
and `seat_based_pricing_enabled` feature flags enabled.

**Scopes**: `member_sessions:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="member-sessions:create" method="post" path="/v1/member-sessions/" -->
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

    res, err := s.MemberSessions.Create(ctx, components.MemberSessionCreate{
        MemberID: "<value>",
        ReturnURL: polargo.Pointer("https://example.com/account"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MemberSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.MemberSessionCreate](../../models/components/membersessioncreate.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.MemberSessionsCreateResponse](../../models/operations/membersessionscreateresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
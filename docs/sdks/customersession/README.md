# CustomerSession
(*CustomerPortal.CustomerSession*)

## Overview

### Available Operations

* [Introspect](#introspect) - Introspect Customer Session

## Introspect

Introspect the current session and return its information.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:customer-session:introspect" method="get" path="/v1/customer-portal/customer-session/introspect" -->
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

    res, err := s.CustomerPortal.CustomerSession.Introspect(ctx, operations.CustomerPortalCustomerSessionIntrospectSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerCustomerSession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `security`                                                                                                                               | [operations.CustomerPortalCustomerSessionIntrospectSecurity](../../models/operations/customerportalcustomersessionintrospectsecurity.md) | :heavy_check_mark:                                                                                                                       | The security requirements to use for the request.                                                                                        |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.CustomerPortalCustomerSessionIntrospectResponse](../../models/operations/customerportalcustomersessionintrospectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |
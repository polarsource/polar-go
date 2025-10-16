# Seats
(*CustomerPortal.Seats*)

## Overview

### Available Operations

* [ListSeats](#listseats) - List Seats
* [AssignSeat](#assignseat) - Assign Seat
* [RevokeSeat](#revokeseat) - Revoke Seat
* [ResendInvitation](#resendinvitation) - Resend Invitation
* [ListClaimedSubscriptions](#listclaimedsubscriptions) - List Claimed Subscriptions

## ListSeats

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:seats:list_seats" method="get" path="/v1/customer-portal/seats" -->
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

    res, err := s.CustomerPortal.Seats.ListSeats(ctx, operations.CustomerPortalSeatsListSeatsSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SeatsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `security`                                                                                                         | [operations.CustomerPortalSeatsListSeatsSecurity](../../models/operations/customerportalseatslistseatssecurity.md) | :heavy_check_mark:                                                                                                 | The security requirements to use for the request.                                                                  |
| `subscriptionID`                                                                                                   | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | Subscription ID                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CustomerPortalSeatsListSeatsResponse](../../models/operations/customerportalseatslistseatsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## AssignSeat

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:seats:assign_seat" method="post" path="/v1/customer-portal/seats" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"github.com/polarsource/polar-go/models/components"
	"os"
	"github.com/polarsource/polar-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerPortal.Seats.AssignSeat(ctx, components.SeatAssign{}, operations.CustomerPortalSeatsAssignSeatSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSeat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [components.SeatAssign](../../models/components/seatassign.md)                                                       | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `security`                                                                                                           | [operations.CustomerPortalSeatsAssignSeatSecurity](../../models/operations/customerportalseatsassignseatsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CustomerPortalSeatsAssignSeatResponse](../../models/operations/customerportalseatsassignseatresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## RevokeSeat

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:seats:revoke_seat" method="delete" path="/v1/customer-portal/seats/{seat_id}" -->
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

    res, err := s.CustomerPortal.Seats.RevokeSeat(ctx, operations.CustomerPortalSeatsRevokeSeatSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "4b3d74b3-01ff-4aac-bd03-320535cd5ce4")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSeat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `security`                                                                                                           | [operations.CustomerPortalSeatsRevokeSeatSecurity](../../models/operations/customerportalseatsrevokeseatsecurity.md) | :heavy_check_mark:                                                                                                   | The security requirements to use for the request.                                                                    |
| `seatID`                                                                                                             | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.CustomerPortalSeatsRevokeSeatResponse](../../models/operations/customerportalseatsrevokeseatresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ResendInvitation

**Scopes**: `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:seats:resend_invitation" method="post" path="/v1/customer-portal/seats/{seat_id}/resend" -->
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

    res, err := s.CustomerPortal.Seats.ResendInvitation(ctx, operations.CustomerPortalSeatsResendInvitationSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    }, "e3817437-8d53-4578-88d2-1dc256825965")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSeat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `security`                                                                                                                       | [operations.CustomerPortalSeatsResendInvitationSecurity](../../models/operations/customerportalseatsresendinvitationsecurity.md) | :heavy_check_mark:                                                                                                               | The security requirements to use for the request.                                                                                |
| `seatID`                                                                                                                         | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.CustomerPortalSeatsResendInvitationResponse](../../models/operations/customerportalseatsresendinvitationresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ListClaimedSubscriptions

List all subscriptions where the authenticated customer has claimed a seat.

**Scopes**: `customer_portal:read` `customer_portal:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer_portal:seats:list_claimed_subscriptions" method="get" path="/v1/customer-portal/seats/subscriptions" -->
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

    res, err := s.CustomerPortal.Seats.ListClaimedSubscriptions(ctx, operations.CustomerPortalSeatsListClaimedSubscriptionsSecurity{
        CustomerSession: os.Getenv("POLAR_CUSTOMER_SESSION"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseCustomerPortalSeatsListClaimedSubscriptions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `security`                                                                                                                                       | [operations.CustomerPortalSeatsListClaimedSubscriptionsSecurity](../../models/operations/customerportalseatslistclaimedsubscriptionssecurity.md) | :heavy_check_mark:                                                                                                                               | The security requirements to use for the request.                                                                                                |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.CustomerPortalSeatsListClaimedSubscriptionsResponse](../../models/operations/customerportalseatslistclaimedsubscriptionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |
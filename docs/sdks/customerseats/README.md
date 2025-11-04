# CustomerSeats
(*CustomerSeats*)

## Overview

### Available Operations

* [AssignSeat](#assignseat) - Assign Seat
* [ListSeats](#listseats) - List Seats
* [RevokeSeat](#revokeseat) - Revoke Seat
* [ResendInvitation](#resendinvitation) - Resend Invitation
* [GetClaimInfo](#getclaiminfo) - Get Claim Info
* [ClaimSeat](#claimseat) - Claim Seat

## AssignSeat

**Scopes**: `customer_seats:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer-seats:assign_seat" method="post" path="/v1/customer-seats" -->
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

    res, err := s.CustomerSeats.AssignSeat(ctx, components.SeatAssign{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSeat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.SeatAssign](../../models/components/seatassign.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.CustomerSeatsAssignSeatResponse](../../models/operations/customerseatsassignseatresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ListSeats

**Scopes**: `customer_seats:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer-seats:list_seats" method="get" path="/v1/customer-seats" -->
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

    res, err := s.CustomerSeats.ListSeats(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.SeatsList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `subscriptionID`                                         | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `orderID`                                                | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomerSeatsListSeatsResponse](../../models/operations/customerseatslistseatsresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## RevokeSeat

**Scopes**: `customer_seats:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer-seats:revoke_seat" method="delete" path="/v1/customer-seats/{seat_id}" -->
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

    res, err := s.CustomerSeats.RevokeSeat(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSeat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `seatID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomerSeatsRevokeSeatResponse](../../models/operations/customerseatsrevokeseatresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ResendInvitation

**Scopes**: `customer_seats:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="customer-seats:resend_invitation" method="post" path="/v1/customer-seats/{seat_id}/resend" -->
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

    res, err := s.CustomerSeats.ResendInvitation(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSeat != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `seatID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomerSeatsResendInvitationResponse](../../models/operations/customerseatsresendinvitationresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## GetClaimInfo

Get Claim Info

### Example Usage

<!-- UsageSnippet language="go" operationID="customer-seats:get_claim_info" method="get" path="/v1/customer-seats/claim/{invitation_token}" -->
```go
package main

import(
	"context"
	polargo "github.com/polarsource/polar-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := polargo.New()

    res, err := s.CustomerSeats.GetClaimInfo(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SeatClaimInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `invitationToken`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CustomerSeatsGetClaimInfoResponse](../../models/operations/customerseatsgetclaiminforesponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |

## ClaimSeat

Claim Seat

### Example Usage

<!-- UsageSnippet language="go" operationID="customer-seats:claim_seat" method="post" path="/v1/customer-seats/claim" -->
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

    res, err := s.CustomerSeats.ClaimSeat(ctx, components.SeatClaim{
        InvitationToken: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerSeatClaimResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [components.SeatClaim](../../models/components/seatclaim.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.CustomerSeatsClaimSeatResponse](../../models/operations/customerseatsclaimseatresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
# Members
(*Members*)

## Overview

### Available Operations

* [ListMembers](#listmembers) - List Members

## ListMembers

List members with optional customer ID filter.

**Scopes**: `members:read` `members:write`

### Example Usage

<!-- UsageSnippet language="go" operationID="members:list_members" method="get" path="/v1/members/" -->
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

    res, err := s.Members.ListMembers(ctx, nil, polargo.Pointer[int64](1), polargo.Pointer[int64](10), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListResourceMember != nil {
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

| Parameter                                                                                                                                                               | Type                                                                                                                                                                    | Required                                                                                                                                                                | Description                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                   | :heavy_check_mark:                                                                                                                                                      | The context to use for the request.                                                                                                                                     |
| `customerID`                                                                                                                                                            | **string*                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                      | Filter by customer ID.                                                                                                                                                  |
| `page`                                                                                                                                                                  | **int64*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | Page number, defaults to 1.                                                                                                                                             |
| `limit`                                                                                                                                                                 | **int64*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | Size of a page, defaults to 10. Maximum is 100.                                                                                                                         |
| `sorting`                                                                                                                                                               | [][components.MemberSortProperty](../../models/components/membersortproperty.md)                                                                                        | :heavy_minus_sign:                                                                                                                                                      | Sorting criterion. Several criteria can be used simultaneously and will be applied in order. Add a minus sign `-` before the criteria name to sort by descending order. |
| `opts`                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | The options for this request.                                                                                                                                           |

### Response

**[*operations.MembersListMembersResponse](../../models/operations/memberslistmembersresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| apierrors.HTTPValidationError | 422                           | application/json              |
| apierrors.APIError            | 4XX, 5XX                      | \*/\*                         |
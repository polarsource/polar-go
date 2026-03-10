# SeatClaimInfo

Read-only information about a seat claim invitation.
Safe for email scanners - no side effects when fetched.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `ProductName`                               | `string`                                    | :heavy_check_mark:                          | Name of the product                         |
| `ProductID`                                 | `string`                                    | :heavy_check_mark:                          | ID of the product                           |
| `OrganizationName`                          | `string`                                    | :heavy_check_mark:                          | Name of the organization                    |
| `OrganizationSlug`                          | `string`                                    | :heavy_check_mark:                          | Slug of the organization                    |
| `CustomerEmail`                             | `string`                                    | :heavy_check_mark:                          | Email of the customer assigned to this seat |
| `CanClaim`                                  | `bool`                                      | :heavy_check_mark:                          | Whether the seat can be claimed             |
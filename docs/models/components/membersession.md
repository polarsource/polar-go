# MemberSession

A member session that can be used to authenticate as a member in the customer portal.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `CreatedAt`                                                | [time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                         | Creation timestamp of the object.                          |
| `ModifiedAt`                                               | [time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                         | Last modification timestamp of the object.                 |
| `ID`                                                       | *string*                                                   | :heavy_check_mark:                                         | The ID of the object.                                      |
| `Token`                                                    | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `ExpiresAt`                                                | [time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                         | N/A                                                        |
| `ReturnURL`                                                | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `MemberPortalURL`                                          | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `MemberID`                                                 | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `Member`                                                   | [components.Member](../../models/components/member.md)     | :heavy_check_mark:                                         | A member of a customer.                                    |
| `CustomerID`                                               | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `Customer`                                                 | [components.Customer](../../models/components/customer.md) | :heavy_check_mark:                                         | A customer in an organization.                             |
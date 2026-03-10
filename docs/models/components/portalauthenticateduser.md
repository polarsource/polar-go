# PortalAuthenticatedUser

Information about the authenticated portal user.


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Type`                                                              | `string`                                                            | :heavy_check_mark:                                                  | Type of authenticated user: 'customer' or 'member'                  |
| `Name`                                                              | `*string`                                                           | :heavy_check_mark:                                                  | User's name, if available.                                          |
| `Email`                                                             | `string`                                                            | :heavy_check_mark:                                                  | User's email address.                                               |
| `CustomerID`                                                        | `string`                                                            | :heavy_check_mark:                                                  | Associated customer ID.                                             |
| `MemberID`                                                          | `*string`                                                           | :heavy_minus_sign:                                                  | Member ID. Only set for members.                                    |
| `Role`                                                              | `*string`                                                           | :heavy_minus_sign:                                                  | Member role (owner, billing_manager, member). Only set for members. |
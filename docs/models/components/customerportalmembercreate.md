# CustomerPortalMemberCreate

Schema for adding a new member to the customer's team.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Email`                                                         | `string`                                                        | :heavy_check_mark:                                              | The email address of the new member.                            |
| `Name`                                                          | `*string`                                                       | :heavy_minus_sign:                                              | The name of the new member (optional).                          |
| `Role`                                                          | [*components.MemberRole](../../models/components/memberrole.md) | :heavy_minus_sign:                                              | N/A                                                             |
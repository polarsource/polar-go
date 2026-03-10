# CustomerPortalMember

A member of the customer's team as seen in the customer portal.


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `CreatedAt`                                                    | [time.Time](https://pkg.go.dev/time#Time)                      | :heavy_check_mark:                                             | Creation timestamp of the object.                              |
| `ModifiedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                             | Last modification timestamp of the object.                     |
| `ID`                                                           | `string`                                                       | :heavy_check_mark:                                             | The ID of the object.                                          |
| `Email`                                                        | `string`                                                       | :heavy_check_mark:                                             | The email address of the member.                               |
| `Name`                                                         | `*string`                                                      | :heavy_check_mark:                                             | The name of the member.                                        |
| `Role`                                                         | [components.MemberRole](../../models/components/memberrole.md) | :heavy_check_mark:                                             | N/A                                                            |
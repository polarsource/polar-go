# BenefitDiscordCreateProperties

Properties to create a benefit of type `discord`.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `GuildToken`                                                      | `string`                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `RoleID`                                                          | `string`                                                          | :heavy_check_mark:                                                | The ID of the Discord role to grant.                              |
| `KickMember`                                                      | `bool`                                                            | :heavy_check_mark:                                                | Whether to kick the member from the Discord server on revocation. |
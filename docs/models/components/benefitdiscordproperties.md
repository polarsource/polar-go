# BenefitDiscordProperties

Properties for a benefit of type `discord`.


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `GuildID`                                                         | `string`                                                          | :heavy_check_mark:                                                | The ID of the Discord server.                                     |
| `RoleID`                                                          | `string`                                                          | :heavy_check_mark:                                                | The ID of the Discord role to grant.                              |
| `KickMember`                                                      | `bool`                                                            | :heavy_check_mark:                                                | Whether to kick the member from the Discord server on revocation. |
| `GuildToken`                                                      | `string`                                                          | :heavy_check_mark:                                                | N/A                                                               |
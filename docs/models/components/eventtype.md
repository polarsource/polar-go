# EventType


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CreatedAt`                                                 | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | Creation timestamp of the object.                           |
| `ModifiedAt`                                                | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                          | Last modification timestamp of the object.                  |
| `ID`                                                        | `string`                                                    | :heavy_check_mark:                                          | The ID of the object.                                       |
| `Name`                                                      | `string`                                                    | :heavy_check_mark:                                          | The name of the event type.                                 |
| `Label`                                                     | `string`                                                    | :heavy_check_mark:                                          | The label for the event type.                               |
| `LabelPropertySelector`                                     | `*string`                                                   | :heavy_minus_sign:                                          | Property path to extract dynamic label from event metadata. |
| `OrganizationID`                                            | `string`                                                    | :heavy_check_mark:                                          | The ID of the organization owning the event type.           |
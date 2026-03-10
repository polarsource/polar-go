# CustomerSubscriptionMeterMeter


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `CreatedAt`                                                            | [time.Time](https://pkg.go.dev/time#Time)                              | :heavy_check_mark:                                                     | Creation timestamp of the object.                                      |
| `ModifiedAt`                                                           | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_check_mark:                                                     | Last modification timestamp of the object.                             |
| `ID`                                                                   | `string`                                                               | :heavy_check_mark:                                                     | The ID of the object.                                                  |
| `Name`                                                                 | `string`                                                               | :heavy_check_mark:                                                     | The name of the meter. Will be shown on customer's invoices and usage. |
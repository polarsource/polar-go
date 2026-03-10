# CustomerStateSubscriptionMeter

Current consumption and spending for a subscription meter.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CreatedAt`                                                 | [time.Time](https://pkg.go.dev/time#Time)                   | :heavy_check_mark:                                          | Creation timestamp of the object.                           |                                                             |
| `ModifiedAt`                                                | [*time.Time](https://pkg.go.dev/time#Time)                  | :heavy_check_mark:                                          | Last modification timestamp of the object.                  |                                                             |
| `ID`                                                        | `string`                                                    | :heavy_check_mark:                                          | The ID of the object.                                       |                                                             |
| `ConsumedUnits`                                             | `float64`                                                   | :heavy_check_mark:                                          | The number of consumed units so far in this billing period. | 25                                                          |
| `CreditedUnits`                                             | `int64`                                                     | :heavy_check_mark:                                          | The number of credited units so far in this billing period. | 100                                                         |
| `Amount`                                                    | `int64`                                                     | :heavy_check_mark:                                          | The amount due in cents so far in this billing period.      | 0                                                           |
| `MeterID`                                                   | `string`                                                    | :heavy_check_mark:                                          | The ID of the meter.                                        | d498a884-e2cd-4d3e-8002-f536468a8b22                        |
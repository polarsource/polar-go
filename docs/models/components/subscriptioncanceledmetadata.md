# SubscriptionCanceledMetadata


## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `SubscriptionID`              | `string`                      | :heavy_check_mark:            | N/A                           |
| `ProductID`                   | `*string`                     | :heavy_minus_sign:            | N/A                           |
| `Amount`                      | `int64`                       | :heavy_check_mark:            | N/A                           |
| `Currency`                    | `string`                      | :heavy_check_mark:            | N/A                           |
| `RecurringInterval`           | `string`                      | :heavy_check_mark:            | N/A                           |
| `RecurringIntervalCount`      | `int64`                       | :heavy_check_mark:            | N/A                           |
| `CustomerCancellationReason`  | `*string`                     | :heavy_minus_sign:            | N/A                           |
| `CustomerCancellationComment` | `*string`                     | :heavy_minus_sign:            | N/A                           |
| `CanceledAt`                  | `string`                      | :heavy_check_mark:            | N/A                           |
| `EndsAt`                      | `*string`                     | :heavy_minus_sign:            | N/A                           |
| `CancelAtPeriodEnd`           | `*bool`                       | :heavy_minus_sign:            | N/A                           |
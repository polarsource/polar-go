# MetricPeriod


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `Timestamp`                               | [time.Time](https://pkg.go.dev/time#Time) | :heavy_check_mark:                        | Timestamp of this period data.            |
| `Orders`                                  | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `Revenue`                                 | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `CumulativeRevenue`                       | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `AverageOrderValue`                       | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `OneTimeProducts`                         | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `OneTimeProductsRevenue`                  | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `NewSubscriptions`                        | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `NewSubscriptionsRevenue`                 | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `RenewedSubscriptions`                    | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `RenewedSubscriptionsRevenue`             | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `ActiveSubscriptions`                     | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
| `MonthlyRecurringRevenue`                 | *int64*                                   | :heavy_check_mark:                        | N/A                                       |
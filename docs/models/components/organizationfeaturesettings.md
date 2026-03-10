# OrganizationFeatureSettings


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `IssueFundingEnabled`                                         | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has issue funding enabled                |
| `SeatBasedPricingEnabled`                                     | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has seat-based pricing enabled           |
| `RevopsEnabled`                                               | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has RevOps enabled                       |
| `WalletsEnabled`                                              | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has Wallets enabled                      |
| `MemberModelEnabled`                                          | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has the Member model enabled             |
| `TinybirdRead`                                                | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization reads from Tinybird                      |
| `TinybirdCompare`                                             | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization compares Tinybird results with database  |
| `CheckoutLocalizationEnabled`                                 | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has checkout localization enabled        |
| `OverviewMetrics`                                             | []`string`                                                    | :heavy_minus_sign:                                            | Ordered list of metric slugs shown on the dashboard overview. |
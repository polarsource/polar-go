# OrganizationFeatureSettingsUpdate

Feature settings that organizations can update themselves.

Other feature settings are managed by Polar staff: they're ignored if
provided and keep their current value.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `SeatBasedPricingEnabled`                                     | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has seat-based pricing enabled           |
| `MemberModelEnabled`                                          | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has the Member model enabled             |
| `CheckoutLocalizationEnabled`                                 | `*bool`                                                       | :heavy_minus_sign:                                            | If this organization has checkout localization enabled        |
| `OverviewMetrics`                                             | []`string`                                                    | :heavy_minus_sign:                                            | Ordered list of metric slugs shown on the dashboard overview. |
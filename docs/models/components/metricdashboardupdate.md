# MetricDashboardUpdate

Schema for updating a metrics dashboard.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Name`                                             | `*string`                                          | :heavy_minus_sign:                                 | Display name for the dashboard.                    |
| `Metrics`                                          | []`string`                                         | :heavy_minus_sign:                                 | List of metric slugs to display in this dashboard. |
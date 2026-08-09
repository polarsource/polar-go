# OrganizationCapabilities


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `CheckoutPayments`                                          | `bool`                                                      | :heavy_check_mark:                                          | Whether the organization can accept new checkout payments.  |
| `SubscriptionRenewals`                                      | `bool`                                                      | :heavy_check_mark:                                          | Whether the organization can process subscription renewals. |
| `Payouts`                                                   | `bool`                                                      | :heavy_check_mark:                                          | Whether the organization can withdraw its balance.          |
| `Refunds`                                                   | `bool`                                                      | :heavy_check_mark:                                          | Whether the organization can issue refunds.                 |
| `APIAccess`                                                 | `bool`                                                      | :heavy_check_mark:                                          | Whether the organization can access the API.                |
| `DashboardAccess`                                           | `bool`                                                      | :heavy_check_mark:                                          | Whether the organization can access the dashboard.          |
# WebhookCustomerDeletedPayload

Sent when a customer is deleted.

**Discord & Slack support:** Basic


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Type`                                                     | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        | customer.deleted                                           |
| `Data`                                                     | [components.Customer](../../models/components/customer.md) | :heavy_check_mark:                                         | A customer in an organization.                             |                                                            |
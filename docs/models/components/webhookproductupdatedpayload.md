# WebhookProductUpdatedPayload

Sent when a product is updated.

**Discord & Slack support:** Basic


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Type`                                                   | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | product.updated                                          |
| `Timestamp`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `Data`                                                   | [components.Product](../../models/components/product.md) | :heavy_check_mark:                                       | A product.                                               |                                                          |
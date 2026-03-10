# WebhookOrderPaidPayload

Sent when an order is paid.

When you receive this event, the order is fully processed and payment has been received.

**Discord & Slack support:** Full


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          | Example                                              |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `Type`                                               | `string`                                             | :heavy_check_mark:                                   | N/A                                                  | order.paid                                           |
| `Timestamp`                                          | [time.Time](https://pkg.go.dev/time#Time)            | :heavy_check_mark:                                   | N/A                                                  |                                                      |
| `Data`                                               | [components.Order](../../models/components/order.md) | :heavy_check_mark:                                   | N/A                                                  |                                                      |
# WebhookRefundCreatedPayload

Sent when a refund is created regardless of status.

**Discord & Slack support:** Full


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Type`                                                 | `string`                                               | :heavy_check_mark:                                     | N/A                                                    | refund.created                                         |
| `Timestamp`                                            | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |                                                        |
| `Data`                                                 | [components.Refund](../../models/components/refund.md) | :heavy_check_mark:                                     | N/A                                                    |                                                        |
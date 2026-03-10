# WebhookBenefitCreatedPayload

Sent when a new benefit is created.

**Discord & Slack support:** Basic


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Type`                                                   | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | benefit.created                                          |
| `Timestamp`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | N/A                                                      |                                                          |
| `Data`                                                   | [components.Benefit](../../models/components/benefit.md) | :heavy_check_mark:                                       | N/A                                                      |                                                          |
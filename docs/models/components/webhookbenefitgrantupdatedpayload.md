# WebhookBenefitGrantUpdatedPayload

Sent when a benefit grant is updated.

**Discord & Slack support:** Basic


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Type`                                                                           | `string`                                                                         | :heavy_check_mark:                                                               | N/A                                                                              | benefit_grant.updated                                                            |
| `Timestamp`                                                                      | [time.Time](https://pkg.go.dev/time#Time)                                        | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `Data`                                                                           | [components.BenefitGrantWebhook](../../models/components/benefitgrantwebhook.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
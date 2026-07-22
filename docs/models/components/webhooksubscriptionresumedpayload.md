# WebhookSubscriptionResumedPayload

Sent when a paused subscription resumes, restoring the customer's access.

Resuming starts a new billing period and charges the customer immediately.

**Discord & Slack support:** Full


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Type`                                                             | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                | subscription.resumed                                               |
| `Timestamp`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `Data`                                                             | [components.Subscription](../../models/components/subscription.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
# WebhookSubscriptionRevokedPayload

Sent when a subscription is revoked and the user loses access immediately.
Happens when the subscription is canceled or payment retries are exhausted (status becomes `unpaid`).

For payment failures that can still be recovered, see `subscription.past_due`.

**Discord & Slack support:** Full


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Type`                                                             | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                | subscription.revoked                                               |
| `Timestamp`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `Data`                                                             | [components.Subscription](../../models/components/subscription.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
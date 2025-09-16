# WebhookSubscriptionRevokedPayload

Sent when a subscription is revoked, the user loses access immediately.
Happens when the subscription is canceled, or payment is past due.

**Discord & Slack support:** Full


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Type`                                                             | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                | subscription.revoked                                               |
| `Timestamp`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `Data`                                                             | [components.Subscription](../../models/components/subscription.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
# WebhookCustomerSeatRevokedPayload

Sent when a customer seat is revoked.

This event is triggered when access to a seat is revoked, either manually by the organization or automatically when a subscription is canceled.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Type`                                                             | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                | customer_seat.revoked                                              |
| `Timestamp`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `Data`                                                             | [components.CustomerSeat](../../models/components/customerseat.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
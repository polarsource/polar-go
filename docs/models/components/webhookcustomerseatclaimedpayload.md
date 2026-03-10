# WebhookCustomerSeatClaimedPayload

Sent when a customer seat is claimed.

This event is triggered when a customer accepts the seat invitation and claims their access.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `Type`                                                             | `string`                                                           | :heavy_check_mark:                                                 | N/A                                                                | customer_seat.claimed                                              |
| `Timestamp`                                                        | [time.Time](https://pkg.go.dev/time#Time)                          | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `Data`                                                             | [components.CustomerSeat](../../models/components/customerseat.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
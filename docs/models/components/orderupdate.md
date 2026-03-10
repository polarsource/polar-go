# OrderUpdate

Schema to update an order.


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `BillingName`                                                                                              | `*string`                                                                                                  | :heavy_minus_sign:                                                                                         | The name of the customer that should appear on the invoice.                                                |
| `BillingAddress`                                                                                           | [*components.AddressInput](../../models/components/addressinput.md)                                        | :heavy_minus_sign:                                                                                         | The address of the customer that should appear on the invoice. Country and state fields cannot be updated. |
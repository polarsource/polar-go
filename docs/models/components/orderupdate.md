# OrderUpdate

Schema to update an order.


## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `BillingName`                                                                                                   | *string*                                                                                                        | :heavy_check_mark:                                                                                              | The name of the customer that should appear on the invoice. Can't be updated after the invoice is generated.    |
| `BillingAddress`                                                                                                | [components.AddressInput](../../models/components/addressinput.md)                                              | :heavy_check_mark:                                                                                              | The address of the customer that should appear on the invoice. Can't be updated after the invoice is generated. |
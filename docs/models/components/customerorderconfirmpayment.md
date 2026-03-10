# CustomerOrderConfirmPayment

Schema to confirm a retry payment using either a saved payment method or a new confirmation token.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ConfirmationTokenID`                                                       | `*string`                                                                   | :heavy_minus_sign:                                                          | ID of the Stripe confirmation token for new payment methods.                |
| `PaymentMethodID`                                                           | `*string`                                                                   | :heavy_minus_sign:                                                          | ID of an existing saved payment method.                                     |
| `PaymentProcessor`                                                          | [*components.PaymentProcessor](../../models/components/paymentprocessor.md) | :heavy_minus_sign:                                                          | N/A                                                                         |
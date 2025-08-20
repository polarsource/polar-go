# CustomerOrderConfirmPayment

Schema to confirm a retry payment using a Stripe confirmation token.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ConfirmationTokenID`                                                       | *string*                                                                    | :heavy_check_mark:                                                          | ID of the Stripe confirmation token.                                        |
| `PaymentProcessor`                                                          | [*components.PaymentProcessor](../../models/components/paymentprocessor.md) | :heavy_minus_sign:                                                          | N/A                                                                         |
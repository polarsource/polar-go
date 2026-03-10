# CardPaymentMetadata

Additional metadata for a card payment method.


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Brand`                                     | `string`                                    | :heavy_check_mark:                          | The brand of the card used for the payment. | **Example 1:** visa<br/>**Example 2:** amex |
| `Last4`                                     | `string`                                    | :heavy_check_mark:                          | The last 4 digits of the card number.       | 4242                                        |
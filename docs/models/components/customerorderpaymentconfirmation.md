# CustomerOrderPaymentConfirmation

Response after confirming a retry payment.


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Status`                                       | `string`                                       | :heavy_check_mark:                             | Payment status after confirmation.             |
| `ClientSecret`                                 | `*string`                                      | :heavy_minus_sign:                             | Client secret for handling additional actions. |
| `Error`                                        | `*string`                                      | :heavy_minus_sign:                             | Error message if confirmation failed.          |
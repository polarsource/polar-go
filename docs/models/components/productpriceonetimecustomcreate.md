# ProductPriceOneTimeCustomCreate

Schema to create a pay-what-you-want price for a one-time product.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Type`                                            | *string*                                          | :heavy_check_mark:                                | N/A                                               |
| `AmountType`                                      | *string*                                          | :heavy_check_mark:                                | N/A                                               |
| `PriceCurrency`                                   | **string*                                         | :heavy_minus_sign:                                | The currency. Currently, only `usd` is supported. |
| `MinimumAmount`                                   | **int64*                                          | :heavy_minus_sign:                                | The minimum amount the customer can pay.          |
| `MaximumAmount`                                   | **int64*                                          | :heavy_minus_sign:                                | The maximum amount the customer can pay.          |
| `PresetAmount`                                    | **int64*                                          | :heavy_minus_sign:                                | The initial amount shown to the customer.         |
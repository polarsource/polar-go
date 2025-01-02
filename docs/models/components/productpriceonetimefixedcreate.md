# ProductPriceOneTimeFixedCreate

Schema to create a one-time product price.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `Type`                                            | *string*                                          | :heavy_check_mark:                                | N/A                                               |
| `AmountType`                                      | *string*                                          | :heavy_check_mark:                                | N/A                                               |
| `PriceAmount`                                     | *int64*                                           | :heavy_check_mark:                                | The price in cents.                               |
| `PriceCurrency`                                   | **string*                                         | :heavy_minus_sign:                                | The currency. Currently, only `usd` is supported. |
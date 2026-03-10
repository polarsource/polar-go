# ProductPriceFixedCreate

Schema to create a fixed price.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AmountType`                                                                      | `string`                                                                          | :heavy_check_mark:                                                                | N/A                                                                               |
| `PriceCurrency`                                                                   | [*components.PresentmentCurrency](../../models/components/presentmentcurrency.md) | :heavy_minus_sign:                                                                | N/A                                                                               |
| `PriceAmount`                                                                     | `int64`                                                                           | :heavy_check_mark:                                                                | The price in cents.                                                               |
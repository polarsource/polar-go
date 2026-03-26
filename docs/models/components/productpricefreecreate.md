# ProductPriceFreeCreate

Schema to create a free price.


## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `AmountType`                                                                                           | `string`                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `PriceCurrency`                                                                                        | [*components.PresentmentCurrency](../../models/components/presentmentcurrency.md)                      | :heavy_minus_sign:                                                                                     | N/A                                                                                                    |
| `TaxBehavior`                                                                                          | [*components.TaxBehaviorOption](../../models/components/taxbehavioroption.md)                          | :heavy_minus_sign:                                                                                     | The tax behavior of the price. If not set, it will default to the organization's default tax behavior. |
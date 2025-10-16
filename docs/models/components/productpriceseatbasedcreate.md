# ProductPriceSeatBasedCreate

Schema to create a seat-based price with volume-based tiers.


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `AmountType`                                                                         | *string*                                                                             | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `PriceCurrency`                                                                      | **string*                                                                            | :heavy_minus_sign:                                                                   | The currency. Currently, only `usd` is supported.                                    |
| `SeatTiers`                                                                          | [components.ProductPriceSeatTiers](../../models/components/productpriceseattiers.md) | :heavy_check_mark:                                                                   | List of pricing tiers for seat-based pricing.                                        |
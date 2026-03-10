# ProductPriceSeatTiersOutput

List of pricing tiers for seat-based pricing.

The minimum and maximum seat limits are derived from the tiers:
- minimum_seats = first tier's min_seats
- maximum_seats = last tier's max_seats (None for unlimited)


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Tiers`                                                                                   | [][components.ProductPriceSeatTier](../../models/components/productpriceseattier.md)      | :heavy_check_mark:                                                                        | List of pricing tiers                                                                     |
| `MinimumSeats`                                                                            | `int64`                                                                                   | :heavy_check_mark:                                                                        | Minimum number of seats required for purchase, derived from first tier.                   |
| `MaximumSeats`                                                                            | `*int64`                                                                                  | :heavy_check_mark:                                                                        | Maximum number of seats allowed for purchase, derived from last tier. None for unlimited. |
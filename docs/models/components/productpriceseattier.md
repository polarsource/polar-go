# ProductPriceSeatTier

A pricing tier for seat-based pricing.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `MinSeats`                                               | `int64`                                                  | :heavy_check_mark:                                       | Minimum number of seats (inclusive)                      |
| `MaxSeats`                                               | `*int64`                                                 | :heavy_minus_sign:                                       | Maximum number of seats (inclusive). None for unlimited. |
| `PricePerSeat`                                           | `int64`                                                  | :heavy_check_mark:                                       | Price per seat in cents for this tier                    |
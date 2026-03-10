# BalanceOrderMetadata


## Fields

| Field                 | Type                  | Required              | Description           |
| --------------------- | --------------------- | --------------------- | --------------------- |
| `TransactionID`       | `string`              | :heavy_check_mark:    | N/A                   |
| `OrderID`             | `string`              | :heavy_check_mark:    | N/A                   |
| `ProductID`           | `*string`             | :heavy_minus_sign:    | N/A                   |
| `SubscriptionID`      | `*string`             | :heavy_minus_sign:    | N/A                   |
| `Amount`              | `int64`               | :heavy_check_mark:    | N/A                   |
| `NetAmount`           | `*int64`              | :heavy_minus_sign:    | N/A                   |
| `Currency`            | `string`              | :heavy_check_mark:    | N/A                   |
| `PresentmentAmount`   | `int64`               | :heavy_check_mark:    | N/A                   |
| `PresentmentCurrency` | `string`              | :heavy_check_mark:    | N/A                   |
| `TaxAmount`           | `int64`               | :heavy_check_mark:    | N/A                   |
| `TaxState`            | `*string`             | :heavy_minus_sign:    | N/A                   |
| `TaxCountry`          | `*string`             | :heavy_minus_sign:    | N/A                   |
| `Fee`                 | `int64`               | :heavy_check_mark:    | N/A                   |
| `ExchangeRate`        | `*float64`            | :heavy_minus_sign:    | N/A                   |
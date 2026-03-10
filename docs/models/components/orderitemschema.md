# OrderItemSchema

An order line item.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `CreatedAt`                                  | [time.Time](https://pkg.go.dev/time#Time)    | :heavy_check_mark:                           | Creation timestamp of the object.            |                                              |
| `ModifiedAt`                                 | [*time.Time](https://pkg.go.dev/time#Time)   | :heavy_check_mark:                           | Last modification timestamp of the object.   |                                              |
| `ID`                                         | `string`                                     | :heavy_check_mark:                           | The ID of the object.                        |                                              |
| `Label`                                      | `string`                                     | :heavy_check_mark:                           | Description of the line item charge.         | Pro Plan                                     |
| `Amount`                                     | `int64`                                      | :heavy_check_mark:                           | Amount in cents, before discounts and taxes. | 10000                                        |
| `TaxAmount`                                  | `int64`                                      | :heavy_check_mark:                           | Sales tax amount in cents.                   | 720                                          |
| `Proration`                                  | `bool`                                       | :heavy_check_mark:                           | Whether this charge is due to a proration.   | false                                        |
| `ProductPriceID`                             | `*string`                                    | :heavy_check_mark:                           | Associated price ID, if any.                 |                                              |
# CustomerWallet

A wallet represents your balance with an organization.

You can top-up your wallet and use the balance to pay for usage.


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  | Example                                      |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `ID`                                         | `string`                                     | :heavy_check_mark:                           | The ID of the object.                        |                                              |
| `CreatedAt`                                  | [time.Time](https://pkg.go.dev/time#Time)    | :heavy_check_mark:                           | Creation timestamp of the object.            |                                              |
| `ModifiedAt`                                 | [*time.Time](https://pkg.go.dev/time#Time)   | :heavy_check_mark:                           | Last modification timestamp of the object.   |                                              |
| `CustomerID`                                 | `string`                                     | :heavy_check_mark:                           | The ID of the customer that owns the wallet. | 992fae2a-2a17-4b7a-8d9e-e287cf90131b         |
| `Balance`                                    | `int64`                                      | :heavy_check_mark:                           | The current balance of the wallet, in cents. | 5000                                         |
| `Currency`                                   | `string`                                     | :heavy_check_mark:                           | The currency of the wallet.                  | usd                                          |
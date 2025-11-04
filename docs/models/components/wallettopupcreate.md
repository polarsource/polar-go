# WalletTopUpCreate

Request schema to top-up a wallet.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Amount`                                                                                 | *int64*                                                                                  | :heavy_check_mark:                                                                       | The amount to top-up the wallet by, in cents.                                            | 2000                                                                                     |
| `Currency`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | The currency. Currently, only `usd` is supported. It should match the wallet's currency. |                                                                                          |
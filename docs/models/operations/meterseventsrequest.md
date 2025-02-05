# MetersEventsRequest


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `ID`                                            | *string*                                        | :heavy_check_mark:                              | The meter ID.                                   |
| `Page`                                          | **int64*                                        | :heavy_minus_sign:                              | Page number, defaults to 1.                     |
| `Limit`                                         | **int64*                                        | :heavy_minus_sign:                              | Size of a page, defaults to 10. Maximum is 100. |
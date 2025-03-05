# FilesListRequest


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `OrganizationID`                                | **string*                                       | :heavy_minus_sign:                              | N/A                                             | 1dbfc517-0bbf-4301-9ba8-555ca42b9737            |
| `Ids`                                           | []*string*                                      | :heavy_minus_sign:                              | List of file IDs to get.                        |                                                 |
| `Page`                                          | **int64*                                        | :heavy_minus_sign:                              | Page number, defaults to 1.                     |                                                 |
| `Limit`                                         | **int64*                                        | :heavy_minus_sign:                              | Size of a page, defaults to 10. Maximum is 100. |                                                 |
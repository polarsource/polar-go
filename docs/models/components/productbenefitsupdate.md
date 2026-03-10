# ProductBenefitsUpdate

Schema to update the benefits granted by a product.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Benefits`                                                                     | []`string`                                                                     | :heavy_check_mark:                                                             | List of benefit IDs. Each one must be on the same organization as the product. |
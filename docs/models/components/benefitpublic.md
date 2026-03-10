# BenefitPublic


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | `string`                                                         | :heavy_check_mark:                                               | The ID of the benefit.                                           |
| `CreatedAt`                                                      | [time.Time](https://pkg.go.dev/time#Time)                        | :heavy_check_mark:                                               | Creation timestamp of the object.                                |
| `ModifiedAt`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                       | :heavy_check_mark:                                               | Last modification timestamp of the object.                       |
| `Type`                                                           | [components.BenefitType](../../models/components/benefittype.md) | :heavy_check_mark:                                               | N/A                                                              |
| `Description`                                                    | `string`                                                         | :heavy_check_mark:                                               | The description of the benefit.                                  |
| `Selectable`                                                     | `bool`                                                           | :heavy_check_mark:                                               | Whether the benefit is selectable when creating a product.       |
| `Deletable`                                                      | `bool`                                                           | :heavy_check_mark:                                               | Whether the benefit is deletable.                                |
| `OrganizationID`                                                 | `string`                                                         | :heavy_check_mark:                                               | The ID of the organization owning the benefit.                   |
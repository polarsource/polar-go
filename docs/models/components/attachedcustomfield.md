# AttachedCustomField

Schema of a custom field attached to a resource.


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `CustomFieldID`                                                  | *string*                                                         | :heavy_check_mark:                                               | ID of the custom field.                                          |
| `CustomField`                                                    | [components.CustomField](../../models/components/customfield.md) | :heavy_check_mark:                                               | N/A                                                              |
| `Order`                                                          | *int64*                                                          | :heavy_check_mark:                                               | Order of the custom field in the resource.                       |
| `Required`                                                       | *bool*                                                           | :heavy_check_mark:                                               | Whether the value is required for this custom field.             |
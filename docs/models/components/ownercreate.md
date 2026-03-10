# OwnerCreate

Schema for creating an owner member during customer creation.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Email`                                                                        | `*string`                                                                      | :heavy_minus_sign:                                                             | The email address of the member.                                               | member@example.com                                                             |
| `Name`                                                                         | `*string`                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            | Jane Doe                                                                       |
| `ExternalID`                                                                   | `*string`                                                                      | :heavy_minus_sign:                                                             | The ID of the member in your system. This must be unique within the customer.  | usr_1337                                                                       |
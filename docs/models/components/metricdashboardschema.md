# MetricDashboardSchema

A user-defined metrics dashboard.


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `CreatedAt`                                       | [time.Time](https://pkg.go.dev/time#Time)         | :heavy_check_mark:                                | Creation timestamp of the object.                 |
| `ModifiedAt`                                      | [*time.Time](https://pkg.go.dev/time#Time)        | :heavy_check_mark:                                | Last modification timestamp of the object.        |
| `ID`                                              | `string`                                          | :heavy_check_mark:                                | The ID of the object.                             |
| `Name`                                            | `string`                                          | :heavy_check_mark:                                | Display name for the dashboard.                   |
| `Metrics`                                         | []`string`                                        | :heavy_check_mark:                                | List of metric slugs displayed in this dashboard. |
| `OrganizationID`                                  | `string`                                          | :heavy_check_mark:                                | The ID of the organization owning this dashboard. |
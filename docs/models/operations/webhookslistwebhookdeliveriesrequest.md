# WebhooksListWebhookDeliveriesRequest


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `EndpointID`                                                    | [*operations.EndpointID](../../models/operations/endpointid.md) | :heavy_minus_sign:                                              | Filter by webhook endpoint ID.                                  |
| `Page`                                                          | **int64*                                                        | :heavy_minus_sign:                                              | Page number, defaults to 1.                                     |
| `Limit`                                                         | **int64*                                                        | :heavy_minus_sign:                                              | Size of a page, defaults to 10. Maximum is 100.                 |
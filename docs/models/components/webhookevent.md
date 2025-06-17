# WebhookEvent

A webhook event.

An event represent something that happened in the system
that should be sent to the webhook endpoint.

It can be delivered multiple times until it's marked as succeeded,
each one creating a new delivery.


## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `CreatedAt`                                                                                                      | [time.Time](https://pkg.go.dev/time#Time)                                                                        | :heavy_check_mark:                                                                                               | Creation timestamp of the object.                                                                                |
| `ModifiedAt`                                                                                                     | [time.Time](https://pkg.go.dev/time#Time)                                                                        | :heavy_check_mark:                                                                                               | Last modification timestamp of the object.                                                                       |
| `ID`                                                                                                             | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The ID of the object.                                                                                            |
| `LastHTTPCode`                                                                                                   | **int64*                                                                                                         | :heavy_minus_sign:                                                                                               | Last HTTP code returned by the URL. `null` if no delviery has been attempted or if the endpoint was unreachable. |
| `Succeeded`                                                                                                      | **bool*                                                                                                          | :heavy_minus_sign:                                                                                               | Whether this event was successfully delivered. `null` if no delivery has been attempted.                         |
| `Payload`                                                                                                        | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The payload of the webhook event.                                                                                |
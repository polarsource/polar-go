# WebhookMemberCreatedPayload

Sent when a new member is created.

A member represents an individual within a customer (team).
This event is triggered when a member is added to a customer,
either programmatically via the API or when an owner is automatically
created for a new customer.

**Discord & Slack support:** Basic


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Type`                                                 | `string`                                               | :heavy_check_mark:                                     | N/A                                                    | member.created                                         |
| `Timestamp`                                            | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |                                                        |
| `Data`                                                 | [components.Member](../../models/components/member.md) | :heavy_check_mark:                                     | A member of a customer.                                |                                                        |
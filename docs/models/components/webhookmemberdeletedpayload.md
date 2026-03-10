# WebhookMemberDeletedPayload

Sent when a member is deleted.

This event is triggered when a member is removed from a customer.
Any active seats assigned to the member will be automatically revoked.

**Discord & Slack support:** Basic


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Type`                                                 | `string`                                               | :heavy_check_mark:                                     | N/A                                                    | member.deleted                                         |
| `Timestamp`                                            | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |                                                        |
| `Data`                                                 | [components.Member](../../models/components/member.md) | :heavy_check_mark:                                     | A member of a customer.                                |                                                        |
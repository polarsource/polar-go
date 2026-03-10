# WebhookMemberUpdatedPayload

Sent when a member is updated.

This event is triggered when member details are updated,
such as their name or role within the customer.

**Discord & Slack support:** Basic


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Type`                                                 | `string`                                               | :heavy_check_mark:                                     | N/A                                                    | member.updated                                         |
| `Timestamp`                                            | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | N/A                                                    |                                                        |
| `Data`                                                 | [components.Member](../../models/components/member.md) | :heavy_check_mark:                                     | A member of a customer.                                |                                                        |
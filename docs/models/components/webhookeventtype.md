# WebhookEventType

## Example Usage

```go
import (
	"github.com/polarsource/polar-go/models/components"
)

value := components.WebhookEventTypeCheckoutCreated
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `WebhookEventTypeCheckoutCreated`        | checkout.created                         |
| `WebhookEventTypeCheckoutUpdated`        | checkout.updated                         |
| `WebhookEventTypeCheckoutExpired`        | checkout.expired                         |
| `WebhookEventTypeCustomerCreated`        | customer.created                         |
| `WebhookEventTypeCustomerUpdated`        | customer.updated                         |
| `WebhookEventTypeCustomerDeleted`        | customer.deleted                         |
| `WebhookEventTypeCustomerStateChanged`   | customer.state_changed                   |
| `WebhookEventTypeCustomerSeatAssigned`   | customer_seat.assigned                   |
| `WebhookEventTypeCustomerSeatClaimed`    | customer_seat.claimed                    |
| `WebhookEventTypeCustomerSeatRevoked`    | customer_seat.revoked                    |
| `WebhookEventTypeMemberCreated`          | member.created                           |
| `WebhookEventTypeMemberUpdated`          | member.updated                           |
| `WebhookEventTypeMemberDeleted`          | member.deleted                           |
| `WebhookEventTypeOrderCreated`           | order.created                            |
| `WebhookEventTypeOrderUpdated`           | order.updated                            |
| `WebhookEventTypeOrderPaid`              | order.paid                               |
| `WebhookEventTypeOrderRefunded`          | order.refunded                           |
| `WebhookEventTypeSubscriptionCreated`    | subscription.created                     |
| `WebhookEventTypeSubscriptionUpdated`    | subscription.updated                     |
| `WebhookEventTypeSubscriptionActive`     | subscription.active                      |
| `WebhookEventTypeSubscriptionCanceled`   | subscription.canceled                    |
| `WebhookEventTypeSubscriptionUncanceled` | subscription.uncanceled                  |
| `WebhookEventTypeSubscriptionRevoked`    | subscription.revoked                     |
| `WebhookEventTypeSubscriptionPastDue`    | subscription.past_due                    |
| `WebhookEventTypeRefundCreated`          | refund.created                           |
| `WebhookEventTypeRefundUpdated`          | refund.updated                           |
| `WebhookEventTypeProductCreated`         | product.created                          |
| `WebhookEventTypeProductUpdated`         | product.updated                          |
| `WebhookEventTypeBenefitCreated`         | benefit.created                          |
| `WebhookEventTypeBenefitUpdated`         | benefit.updated                          |
| `WebhookEventTypeBenefitGrantCreated`    | benefit_grant.created                    |
| `WebhookEventTypeBenefitGrantCycled`     | benefit_grant.cycled                     |
| `WebhookEventTypeBenefitGrantUpdated`    | benefit_grant.updated                    |
| `WebhookEventTypeBenefitGrantRevoked`    | benefit_grant.revoked                    |
| `WebhookEventTypeOrganizationUpdated`    | organization.updated                     |
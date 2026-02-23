# SystemEvent


## Supported Types

### BalanceCreditOrderEvent

```go
systemEvent := components.CreateSystemEventBalanceCreditOrder(components.BalanceCreditOrderEvent{/* values here */})
```

### BalanceDisputeEvent

```go
systemEvent := components.CreateSystemEventBalanceDispute(components.BalanceDisputeEvent{/* values here */})
```

### BalanceDisputeReversalEvent

```go
systemEvent := components.CreateSystemEventBalanceDisputeReversal(components.BalanceDisputeReversalEvent{/* values here */})
```

### BalanceOrderEvent

```go
systemEvent := components.CreateSystemEventBalanceOrder(components.BalanceOrderEvent{/* values here */})
```

### BalanceRefundEvent

```go
systemEvent := components.CreateSystemEventBalanceRefund(components.BalanceRefundEvent{/* values here */})
```

### BalanceRefundReversalEvent

```go
systemEvent := components.CreateSystemEventBalanceRefundReversal(components.BalanceRefundReversalEvent{/* values here */})
```

### BenefitCycledEvent

```go
systemEvent := components.CreateSystemEventBenefitCycled(components.BenefitCycledEvent{/* values here */})
```

### BenefitGrantedEvent

```go
systemEvent := components.CreateSystemEventBenefitGranted(components.BenefitGrantedEvent{/* values here */})
```

### BenefitRevokedEvent

```go
systemEvent := components.CreateSystemEventBenefitRevoked(components.BenefitRevokedEvent{/* values here */})
```

### BenefitUpdatedEvent

```go
systemEvent := components.CreateSystemEventBenefitUpdated(components.BenefitUpdatedEvent{/* values here */})
```

### CheckoutCreatedEvent

```go
systemEvent := components.CreateSystemEventCheckoutCreated(components.CheckoutCreatedEvent{/* values here */})
```

### CustomerCreatedEvent

```go
systemEvent := components.CreateSystemEventCustomerCreated(components.CustomerCreatedEvent{/* values here */})
```

### CustomerDeletedEvent

```go
systemEvent := components.CreateSystemEventCustomerDeleted(components.CustomerDeletedEvent{/* values here */})
```

### CustomerUpdatedEvent

```go
systemEvent := components.CreateSystemEventCustomerUpdated(components.CustomerUpdatedEvent{/* values here */})
```

### MeterCreditEvent

```go
systemEvent := components.CreateSystemEventMeterCredited(components.MeterCreditEvent{/* values here */})
```

### MeterResetEvent

```go
systemEvent := components.CreateSystemEventMeterReset(components.MeterResetEvent{/* values here */})
```

### OrderPaidEvent

```go
systemEvent := components.CreateSystemEventOrderPaid(components.OrderPaidEvent{/* values here */})
```

### OrderRefundedEvent

```go
systemEvent := components.CreateSystemEventOrderRefunded(components.OrderRefundedEvent{/* values here */})
```

### SubscriptionBillingPeriodUpdatedEvent

```go
systemEvent := components.CreateSystemEventSubscriptionBillingPeriodUpdated(components.SubscriptionBillingPeriodUpdatedEvent{/* values here */})
```

### SubscriptionCanceledEvent

```go
systemEvent := components.CreateSystemEventSubscriptionCanceled(components.SubscriptionCanceledEvent{/* values here */})
```

### SubscriptionCreatedEvent

```go
systemEvent := components.CreateSystemEventSubscriptionCreated(components.SubscriptionCreatedEvent{/* values here */})
```

### SubscriptionCycledEvent

```go
systemEvent := components.CreateSystemEventSubscriptionCycled(components.SubscriptionCycledEvent{/* values here */})
```

### SubscriptionProductUpdatedEvent

```go
systemEvent := components.CreateSystemEventSubscriptionProductUpdated(components.SubscriptionProductUpdatedEvent{/* values here */})
```

### SubscriptionRevokedEvent

```go
systemEvent := components.CreateSystemEventSubscriptionRevoked(components.SubscriptionRevokedEvent{/* values here */})
```

### SubscriptionSeatsUpdatedEvent

```go
systemEvent := components.CreateSystemEventSubscriptionSeatsUpdated(components.SubscriptionSeatsUpdatedEvent{/* values here */})
```

### SubscriptionUncanceledEvent

```go
systemEvent := components.CreateSystemEventSubscriptionUncanceled(components.SubscriptionUncanceledEvent{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch systemEvent.Type {
	case components.SystemEventTypeBalanceCreditOrder:
		// systemEvent.BalanceCreditOrderEvent is populated
	case components.SystemEventTypeBalanceDispute:
		// systemEvent.BalanceDisputeEvent is populated
	case components.SystemEventTypeBalanceDisputeReversal:
		// systemEvent.BalanceDisputeReversalEvent is populated
	case components.SystemEventTypeBalanceOrder:
		// systemEvent.BalanceOrderEvent is populated
	case components.SystemEventTypeBalanceRefund:
		// systemEvent.BalanceRefundEvent is populated
	case components.SystemEventTypeBalanceRefundReversal:
		// systemEvent.BalanceRefundReversalEvent is populated
	case components.SystemEventTypeBenefitCycled:
		// systemEvent.BenefitCycledEvent is populated
	case components.SystemEventTypeBenefitGranted:
		// systemEvent.BenefitGrantedEvent is populated
	case components.SystemEventTypeBenefitRevoked:
		// systemEvent.BenefitRevokedEvent is populated
	case components.SystemEventTypeBenefitUpdated:
		// systemEvent.BenefitUpdatedEvent is populated
	case components.SystemEventTypeCheckoutCreated:
		// systemEvent.CheckoutCreatedEvent is populated
	case components.SystemEventTypeCustomerCreated:
		// systemEvent.CustomerCreatedEvent is populated
	case components.SystemEventTypeCustomerDeleted:
		// systemEvent.CustomerDeletedEvent is populated
	case components.SystemEventTypeCustomerUpdated:
		// systemEvent.CustomerUpdatedEvent is populated
	case components.SystemEventTypeMeterCredited:
		// systemEvent.MeterCreditEvent is populated
	case components.SystemEventTypeMeterReset:
		// systemEvent.MeterResetEvent is populated
	case components.SystemEventTypeOrderPaid:
		// systemEvent.OrderPaidEvent is populated
	case components.SystemEventTypeOrderRefunded:
		// systemEvent.OrderRefundedEvent is populated
	case components.SystemEventTypeSubscriptionBillingPeriodUpdated:
		// systemEvent.SubscriptionBillingPeriodUpdatedEvent is populated
	case components.SystemEventTypeSubscriptionCanceled:
		// systemEvent.SubscriptionCanceledEvent is populated
	case components.SystemEventTypeSubscriptionCreated:
		// systemEvent.SubscriptionCreatedEvent is populated
	case components.SystemEventTypeSubscriptionCycled:
		// systemEvent.SubscriptionCycledEvent is populated
	case components.SystemEventTypeSubscriptionProductUpdated:
		// systemEvent.SubscriptionProductUpdatedEvent is populated
	case components.SystemEventTypeSubscriptionRevoked:
		// systemEvent.SubscriptionRevokedEvent is populated
	case components.SystemEventTypeSubscriptionSeatsUpdated:
		// systemEvent.SubscriptionSeatsUpdatedEvent is populated
	case components.SystemEventTypeSubscriptionUncanceled:
		// systemEvent.SubscriptionUncanceledEvent is populated
}
```

# OrdersFinalizeResponse403OrdersFinalize

Off-session charges are not enabled for this organization, or its account can't currently accept payments.


## Supported Types

### OffSessionChargesNotEnabled

```go
ordersFinalizeResponse403OrdersFinalize := apierrors.CreateOrdersFinalizeResponse403OrdersFinalizeOffSessionChargesNotEnabled(components.OffSessionChargesNotEnabled{/* values here */})
```

### OrganizationNotReadyForPayments

```go
ordersFinalizeResponse403OrdersFinalize := apierrors.CreateOrdersFinalizeResponse403OrdersFinalizeOrganizationNotReadyForPayments(components.OrganizationNotReadyForPayments{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersFinalizeResponse403OrdersFinalize.Type {
	case apierrors.OrdersFinalizeResponse403OrdersFinalizeTypeOffSessionChargesNotEnabled:
		// ordersFinalizeResponse403OrdersFinalize.OffSessionChargesNotEnabled is populated
	case apierrors.OrdersFinalizeResponse403OrdersFinalizeTypeOrganizationNotReadyForPayments:
		// ordersFinalizeResponse403OrdersFinalize.OrganizationNotReadyForPayments is populated
}
```

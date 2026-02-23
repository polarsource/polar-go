# OrdersGenerateInvoiceResponse422OrdersGenerateInvoice

Order is not paid or is missing billing name or address.


## Supported Types

### MissingInvoiceBillingDetails

```go
ordersGenerateInvoiceResponse422OrdersGenerateInvoice := apierrors.CreateOrdersGenerateInvoiceResponse422OrdersGenerateInvoiceMissingInvoiceBillingDetails(components.MissingInvoiceBillingDetails{/* values here */})
```

### NotPaidOrder

```go
ordersGenerateInvoiceResponse422OrdersGenerateInvoice := apierrors.CreateOrdersGenerateInvoiceResponse422OrdersGenerateInvoiceNotPaidOrder(components.NotPaidOrder{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch ordersGenerateInvoiceResponse422OrdersGenerateInvoice.Type {
	case apierrors.OrdersGenerateInvoiceResponse422OrdersGenerateInvoiceTypeMissingInvoiceBillingDetails:
		// ordersGenerateInvoiceResponse422OrdersGenerateInvoice.MissingInvoiceBillingDetails is populated
	case apierrors.OrdersGenerateInvoiceResponse422OrdersGenerateInvoiceTypeNotPaidOrder:
		// ordersGenerateInvoiceResponse422OrdersGenerateInvoice.NotPaidOrder is populated
}
```

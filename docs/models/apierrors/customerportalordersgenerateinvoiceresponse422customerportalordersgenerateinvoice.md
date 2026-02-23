# CustomerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoice

Order is not paid or is missing billing name or address.


## Supported Types

### MissingInvoiceBillingDetails

```go
customerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoice := apierrors.CreateCustomerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoiceMissingInvoiceBillingDetails(components.MissingInvoiceBillingDetails{/* values here */})
```

### NotPaidOrder

```go
customerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoice := apierrors.CreateCustomerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoiceNotPaidOrder(components.NotPaidOrder{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoice.Type {
	case apierrors.CustomerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoiceTypeMissingInvoiceBillingDetails:
		// customerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoice.MissingInvoiceBillingDetails is populated
	case apierrors.CustomerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoiceTypeNotPaidOrder:
		// customerPortalOrdersGenerateInvoiceResponse422CustomerPortalOrdersGenerateInvoice.NotPaidOrder is populated
}
```

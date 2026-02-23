# CheckoutLinkCreate


## Supported Types

### CheckoutLinkCreateProductPrice

```go
checkoutLinkCreate := components.CreateCheckoutLinkCreateCheckoutLinkCreateProductPrice(components.CheckoutLinkCreateProductPrice{/* values here */})
```

### CheckoutLinkCreateProduct

```go
checkoutLinkCreate := components.CreateCheckoutLinkCreateCheckoutLinkCreateProduct(components.CheckoutLinkCreateProduct{/* values here */})
```

### CheckoutLinkCreateProducts

```go
checkoutLinkCreate := components.CreateCheckoutLinkCreateCheckoutLinkCreateProducts(components.CheckoutLinkCreateProducts{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinkCreate.Type {
	case components.CheckoutLinkCreateTypeCheckoutLinkCreateProductPrice:
		// checkoutLinkCreate.CheckoutLinkCreateProductPrice is populated
	case components.CheckoutLinkCreateTypeCheckoutLinkCreateProduct:
		// checkoutLinkCreate.CheckoutLinkCreateProduct is populated
	case components.CheckoutLinkCreateTypeCheckoutLinkCreateProducts:
		// checkoutLinkCreate.CheckoutLinkCreateProducts is populated
}
```

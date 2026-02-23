# CheckoutLinkCreateProductsMetadata


## Supported Types

### 

```go
checkoutLinkCreateProductsMetadata := components.CreateCheckoutLinkCreateProductsMetadataStr(string{/* values here */})
```

### 

```go
checkoutLinkCreateProductsMetadata := components.CreateCheckoutLinkCreateProductsMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutLinkCreateProductsMetadata := components.CreateCheckoutLinkCreateProductsMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutLinkCreateProductsMetadata := components.CreateCheckoutLinkCreateProductsMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinkCreateProductsMetadata.Type {
	case components.CheckoutLinkCreateProductsMetadataTypeStr:
		// checkoutLinkCreateProductsMetadata.Str is populated
	case components.CheckoutLinkCreateProductsMetadataTypeInteger:
		// checkoutLinkCreateProductsMetadata.Integer is populated
	case components.CheckoutLinkCreateProductsMetadataTypeNumber:
		// checkoutLinkCreateProductsMetadata.Number is populated
	case components.CheckoutLinkCreateProductsMetadataTypeBoolean:
		// checkoutLinkCreateProductsMetadata.Boolean is populated
}
```

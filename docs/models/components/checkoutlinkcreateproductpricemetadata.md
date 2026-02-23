# CheckoutLinkCreateProductPriceMetadata


## Supported Types

### 

```go
checkoutLinkCreateProductPriceMetadata := components.CreateCheckoutLinkCreateProductPriceMetadataStr(string{/* values here */})
```

### 

```go
checkoutLinkCreateProductPriceMetadata := components.CreateCheckoutLinkCreateProductPriceMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutLinkCreateProductPriceMetadata := components.CreateCheckoutLinkCreateProductPriceMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutLinkCreateProductPriceMetadata := components.CreateCheckoutLinkCreateProductPriceMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinkCreateProductPriceMetadata.Type {
	case components.CheckoutLinkCreateProductPriceMetadataTypeStr:
		// checkoutLinkCreateProductPriceMetadata.Str is populated
	case components.CheckoutLinkCreateProductPriceMetadataTypeInteger:
		// checkoutLinkCreateProductPriceMetadata.Integer is populated
	case components.CheckoutLinkCreateProductPriceMetadataTypeNumber:
		// checkoutLinkCreateProductPriceMetadata.Number is populated
	case components.CheckoutLinkCreateProductPriceMetadataTypeBoolean:
		// checkoutLinkCreateProductPriceMetadata.Boolean is populated
}
```

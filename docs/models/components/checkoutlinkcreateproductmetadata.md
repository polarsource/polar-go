# CheckoutLinkCreateProductMetadata


## Supported Types

### 

```go
checkoutLinkCreateProductMetadata := components.CreateCheckoutLinkCreateProductMetadataStr(string{/* values here */})
```

### 

```go
checkoutLinkCreateProductMetadata := components.CreateCheckoutLinkCreateProductMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutLinkCreateProductMetadata := components.CreateCheckoutLinkCreateProductMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutLinkCreateProductMetadata := components.CreateCheckoutLinkCreateProductMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinkCreateProductMetadata.Type {
	case components.CheckoutLinkCreateProductMetadataTypeStr:
		// checkoutLinkCreateProductMetadata.Str is populated
	case components.CheckoutLinkCreateProductMetadataTypeInteger:
		// checkoutLinkCreateProductMetadata.Integer is populated
	case components.CheckoutLinkCreateProductMetadataTypeNumber:
		// checkoutLinkCreateProductMetadata.Number is populated
	case components.CheckoutLinkCreateProductMetadataTypeBoolean:
		// checkoutLinkCreateProductMetadata.Boolean is populated
}
```

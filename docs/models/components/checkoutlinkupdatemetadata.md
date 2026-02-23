# CheckoutLinkUpdateMetadata


## Supported Types

### 

```go
checkoutLinkUpdateMetadata := components.CreateCheckoutLinkUpdateMetadataStr(string{/* values here */})
```

### 

```go
checkoutLinkUpdateMetadata := components.CreateCheckoutLinkUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutLinkUpdateMetadata := components.CreateCheckoutLinkUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutLinkUpdateMetadata := components.CreateCheckoutLinkUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutLinkUpdateMetadata.Type {
	case components.CheckoutLinkUpdateMetadataTypeStr:
		// checkoutLinkUpdateMetadata.Str is populated
	case components.CheckoutLinkUpdateMetadataTypeInteger:
		// checkoutLinkUpdateMetadata.Integer is populated
	case components.CheckoutLinkUpdateMetadataTypeNumber:
		// checkoutLinkUpdateMetadata.Number is populated
	case components.CheckoutLinkUpdateMetadataTypeBoolean:
		// checkoutLinkUpdateMetadata.Boolean is populated
}
```

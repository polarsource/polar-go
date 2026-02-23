# CheckoutUpdateMetadata


## Supported Types

### 

```go
checkoutUpdateMetadata := components.CreateCheckoutUpdateMetadataStr(string{/* values here */})
```

### 

```go
checkoutUpdateMetadata := components.CreateCheckoutUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutUpdateMetadata := components.CreateCheckoutUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutUpdateMetadata := components.CreateCheckoutUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutUpdateMetadata.Type {
	case components.CheckoutUpdateMetadataTypeStr:
		// checkoutUpdateMetadata.Str is populated
	case components.CheckoutUpdateMetadataTypeInteger:
		// checkoutUpdateMetadata.Integer is populated
	case components.CheckoutUpdateMetadataTypeNumber:
		// checkoutUpdateMetadata.Number is populated
	case components.CheckoutUpdateMetadataTypeBoolean:
		// checkoutUpdateMetadata.Boolean is populated
}
```

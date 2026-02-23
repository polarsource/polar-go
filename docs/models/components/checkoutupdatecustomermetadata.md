# CheckoutUpdateCustomerMetadata


## Supported Types

### 

```go
checkoutUpdateCustomerMetadata := components.CreateCheckoutUpdateCustomerMetadataStr(string{/* values here */})
```

### 

```go
checkoutUpdateCustomerMetadata := components.CreateCheckoutUpdateCustomerMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutUpdateCustomerMetadata := components.CreateCheckoutUpdateCustomerMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutUpdateCustomerMetadata := components.CreateCheckoutUpdateCustomerMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutUpdateCustomerMetadata.Type {
	case components.CheckoutUpdateCustomerMetadataTypeStr:
		// checkoutUpdateCustomerMetadata.Str is populated
	case components.CheckoutUpdateCustomerMetadataTypeInteger:
		// checkoutUpdateCustomerMetadata.Integer is populated
	case components.CheckoutUpdateCustomerMetadataTypeNumber:
		// checkoutUpdateCustomerMetadata.Number is populated
	case components.CheckoutUpdateCustomerMetadataTypeBoolean:
		// checkoutUpdateCustomerMetadata.Boolean is populated
}
```

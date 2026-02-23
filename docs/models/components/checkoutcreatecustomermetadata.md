# CheckoutCreateCustomerMetadata


## Supported Types

### 

```go
checkoutCreateCustomerMetadata := components.CreateCheckoutCreateCustomerMetadataStr(string{/* values here */})
```

### 

```go
checkoutCreateCustomerMetadata := components.CreateCheckoutCreateCustomerMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutCreateCustomerMetadata := components.CreateCheckoutCreateCustomerMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutCreateCustomerMetadata := components.CreateCheckoutCreateCustomerMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutCreateCustomerMetadata.Type {
	case components.CheckoutCreateCustomerMetadataTypeStr:
		// checkoutCreateCustomerMetadata.Str is populated
	case components.CheckoutCreateCustomerMetadataTypeInteger:
		// checkoutCreateCustomerMetadata.Integer is populated
	case components.CheckoutCreateCustomerMetadataTypeNumber:
		// checkoutCreateCustomerMetadata.Number is populated
	case components.CheckoutCreateCustomerMetadataTypeBoolean:
		// checkoutCreateCustomerMetadata.Boolean is populated
}
```

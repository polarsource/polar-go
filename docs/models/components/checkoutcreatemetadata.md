# CheckoutCreateMetadata


## Supported Types

### 

```go
checkoutCreateMetadata := components.CreateCheckoutCreateMetadataStr(string{/* values here */})
```

### 

```go
checkoutCreateMetadata := components.CreateCheckoutCreateMetadataInteger(int64{/* values here */})
```

### 

```go
checkoutCreateMetadata := components.CreateCheckoutCreateMetadataNumber(float64{/* values here */})
```

### 

```go
checkoutCreateMetadata := components.CreateCheckoutCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch checkoutCreateMetadata.Type {
	case components.CheckoutCreateMetadataTypeStr:
		// checkoutCreateMetadata.Str is populated
	case components.CheckoutCreateMetadataTypeInteger:
		// checkoutCreateMetadata.Integer is populated
	case components.CheckoutCreateMetadataTypeNumber:
		// checkoutCreateMetadata.Number is populated
	case components.CheckoutCreateMetadataTypeBoolean:
		// checkoutCreateMetadata.Boolean is populated
}
```

# RefundCreateMetadata


## Supported Types

### 

```go
refundCreateMetadata := components.CreateRefundCreateMetadataStr(string{/* values here */})
```

### 

```go
refundCreateMetadata := components.CreateRefundCreateMetadataInteger(int64{/* values here */})
```

### 

```go
refundCreateMetadata := components.CreateRefundCreateMetadataNumber(float64{/* values here */})
```

### 

```go
refundCreateMetadata := components.CreateRefundCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch refundCreateMetadata.Type {
	case components.RefundCreateMetadataTypeStr:
		// refundCreateMetadata.Str is populated
	case components.RefundCreateMetadataTypeInteger:
		// refundCreateMetadata.Integer is populated
	case components.RefundCreateMetadataTypeNumber:
		// refundCreateMetadata.Number is populated
	case components.RefundCreateMetadataTypeBoolean:
		// refundCreateMetadata.Boolean is populated
}
```

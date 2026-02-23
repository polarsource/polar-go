# CustomerCreateMetadata


## Supported Types

### 

```go
customerCreateMetadata := components.CreateCustomerCreateMetadataStr(string{/* values here */})
```

### 

```go
customerCreateMetadata := components.CreateCustomerCreateMetadataInteger(int64{/* values here */})
```

### 

```go
customerCreateMetadata := components.CreateCustomerCreateMetadataNumber(float64{/* values here */})
```

### 

```go
customerCreateMetadata := components.CreateCustomerCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerCreateMetadata.Type {
	case components.CustomerCreateMetadataTypeStr:
		// customerCreateMetadata.Str is populated
	case components.CustomerCreateMetadataTypeInteger:
		// customerCreateMetadata.Integer is populated
	case components.CustomerCreateMetadataTypeNumber:
		// customerCreateMetadata.Number is populated
	case components.CustomerCreateMetadataTypeBoolean:
		// customerCreateMetadata.Boolean is populated
}
```

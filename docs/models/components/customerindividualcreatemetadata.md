# CustomerIndividualCreateMetadata


## Supported Types

### 

```go
customerIndividualCreateMetadata := components.CreateCustomerIndividualCreateMetadataStr(string{/* values here */})
```

### 

```go
customerIndividualCreateMetadata := components.CreateCustomerIndividualCreateMetadataInteger(int64{/* values here */})
```

### 

```go
customerIndividualCreateMetadata := components.CreateCustomerIndividualCreateMetadataNumber(float64{/* values here */})
```

### 

```go
customerIndividualCreateMetadata := components.CreateCustomerIndividualCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerIndividualCreateMetadata.Type {
	case components.CustomerIndividualCreateMetadataTypeStr:
		// customerIndividualCreateMetadata.Str is populated
	case components.CustomerIndividualCreateMetadataTypeInteger:
		// customerIndividualCreateMetadata.Integer is populated
	case components.CustomerIndividualCreateMetadataTypeNumber:
		// customerIndividualCreateMetadata.Number is populated
	case components.CustomerIndividualCreateMetadataTypeBoolean:
		// customerIndividualCreateMetadata.Boolean is populated
}
```

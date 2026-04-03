# CustomerTeamCreateMetadata


## Supported Types

### 

```go
customerTeamCreateMetadata := components.CreateCustomerTeamCreateMetadataStr(string{/* values here */})
```

### 

```go
customerTeamCreateMetadata := components.CreateCustomerTeamCreateMetadataInteger(int64{/* values here */})
```

### 

```go
customerTeamCreateMetadata := components.CreateCustomerTeamCreateMetadataNumber(float64{/* values here */})
```

### 

```go
customerTeamCreateMetadata := components.CreateCustomerTeamCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch customerTeamCreateMetadata.Type {
	case components.CustomerTeamCreateMetadataTypeStr:
		// customerTeamCreateMetadata.Str is populated
	case components.CustomerTeamCreateMetadataTypeInteger:
		// customerTeamCreateMetadata.Integer is populated
	case components.CustomerTeamCreateMetadataTypeNumber:
		// customerTeamCreateMetadata.Number is populated
	case components.CustomerTeamCreateMetadataTypeBoolean:
		// customerTeamCreateMetadata.Boolean is populated
}
```

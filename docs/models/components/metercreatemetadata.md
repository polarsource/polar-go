# MeterCreateMetadata


## Supported Types

### 

```go
meterCreateMetadata := components.CreateMeterCreateMetadataStr(string{/* values here */})
```

### 

```go
meterCreateMetadata := components.CreateMeterCreateMetadataInteger(int64{/* values here */})
```

### 

```go
meterCreateMetadata := components.CreateMeterCreateMetadataNumber(float64{/* values here */})
```

### 

```go
meterCreateMetadata := components.CreateMeterCreateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch meterCreateMetadata.Type {
	case components.MeterCreateMetadataTypeStr:
		// meterCreateMetadata.Str is populated
	case components.MeterCreateMetadataTypeInteger:
		// meterCreateMetadata.Integer is populated
	case components.MeterCreateMetadataTypeNumber:
		// meterCreateMetadata.Number is populated
	case components.MeterCreateMetadataTypeBoolean:
		// meterCreateMetadata.Boolean is populated
}
```

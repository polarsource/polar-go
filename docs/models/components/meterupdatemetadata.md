# MeterUpdateMetadata


## Supported Types

### 

```go
meterUpdateMetadata := components.CreateMeterUpdateMetadataStr(string{/* values here */})
```

### 

```go
meterUpdateMetadata := components.CreateMeterUpdateMetadataInteger(int64{/* values here */})
```

### 

```go
meterUpdateMetadata := components.CreateMeterUpdateMetadataNumber(float64{/* values here */})
```

### 

```go
meterUpdateMetadata := components.CreateMeterUpdateMetadataBoolean(bool{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch meterUpdateMetadata.Type {
	case components.MeterUpdateMetadataTypeStr:
		// meterUpdateMetadata.Str is populated
	case components.MeterUpdateMetadataTypeInteger:
		// meterUpdateMetadata.Integer is populated
	case components.MeterUpdateMetadataTypeNumber:
		// meterUpdateMetadata.Number is populated
	case components.MeterUpdateMetadataTypeBoolean:
		// meterUpdateMetadata.Boolean is populated
}
```

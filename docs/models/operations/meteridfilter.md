# MeterIDFilter

Filter by meter ID.


## Supported Types

### 

```go
meterIDFilter := operations.CreateMeterIDFilterStr(string{/* values here */})
```

### 

```go
meterIDFilter := operations.CreateMeterIDFilterArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch meterIDFilter.Type {
	case operations.MeterIDFilterTypeStr:
		// meterIDFilter.Str is populated
	case operations.MeterIDFilterTypeArrayOfStr:
		// meterIDFilter.ArrayOfStr is populated
}
```

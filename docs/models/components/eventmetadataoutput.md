# EventMetadataOutput


## Supported Types

### 

```go
eventMetadataOutput := components.CreateEventMetadataOutputStr(string{/* values here */})
```

### 

```go
eventMetadataOutput := components.CreateEventMetadataOutputInteger(int64{/* values here */})
```

### 

```go
eventMetadataOutput := components.CreateEventMetadataOutputNumber(float64{/* values here */})
```

### 

```go
eventMetadataOutput := components.CreateEventMetadataOutputBoolean(bool{/* values here */})
```

### CostMetadataOutput

```go
eventMetadataOutput := components.CreateEventMetadataOutputCostMetadataOutput(components.CostMetadataOutput{/* values here */})
```

### LLMMetadata

```go
eventMetadataOutput := components.CreateEventMetadataOutputLLMMetadata(components.LLMMetadata{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventMetadataOutput.Type {
	case components.EventMetadataOutputTypeStr:
		// eventMetadataOutput.Str is populated
	case components.EventMetadataOutputTypeInteger:
		// eventMetadataOutput.Integer is populated
	case components.EventMetadataOutputTypeNumber:
		// eventMetadataOutput.Number is populated
	case components.EventMetadataOutputTypeBoolean:
		// eventMetadataOutput.Boolean is populated
	case components.EventMetadataOutputTypeCostMetadataOutput:
		// eventMetadataOutput.CostMetadataOutput is populated
	case components.EventMetadataOutputTypeLLMMetadata:
		// eventMetadataOutput.LLMMetadata is populated
}
```

# EventMetadataInput


## Supported Types

### 

```go
eventMetadataInput := components.CreateEventMetadataInputStr(string{/* values here */})
```

### 

```go
eventMetadataInput := components.CreateEventMetadataInputInteger(int64{/* values here */})
```

### 

```go
eventMetadataInput := components.CreateEventMetadataInputNumber(float64{/* values here */})
```

### 

```go
eventMetadataInput := components.CreateEventMetadataInputBoolean(bool{/* values here */})
```

### CostMetadataInput

```go
eventMetadataInput := components.CreateEventMetadataInputCostMetadataInput(components.CostMetadataInput{/* values here */})
```

### LLMMetadata

```go
eventMetadataInput := components.CreateEventMetadataInputLLMMetadata(components.LLMMetadata{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch eventMetadataInput.Type {
	case components.EventMetadataInputTypeStr:
		// eventMetadataInput.Str is populated
	case components.EventMetadataInputTypeInteger:
		// eventMetadataInput.Integer is populated
	case components.EventMetadataInputTypeNumber:
		// eventMetadataInput.Number is populated
	case components.EventMetadataInputTypeBoolean:
		// eventMetadataInput.Boolean is populated
	case components.EventMetadataInputTypeCostMetadataInput:
		// eventMetadataInput.CostMetadataInput is populated
	case components.EventMetadataInputTypeLLMMetadata:
		// eventMetadataInput.LLMMetadata is populated
}
```

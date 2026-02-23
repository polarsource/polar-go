# NetAverageOrderValue


## Supported Types

### 

```go
netAverageOrderValue := components.CreateNetAverageOrderValueInteger(int64{/* values here */})
```

### 

```go
netAverageOrderValue := components.CreateNetAverageOrderValueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch netAverageOrderValue.Type {
	case components.NetAverageOrderValueTypeInteger:
		// netAverageOrderValue.Integer is populated
	case components.NetAverageOrderValueTypeNumber:
		// netAverageOrderValue.Number is populated
}
```

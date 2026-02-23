# AverageOrderValue


## Supported Types

### 

```go
averageOrderValue := components.CreateAverageOrderValueInteger(int64{/* values here */})
```

### 

```go
averageOrderValue := components.CreateAverageOrderValueNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch averageOrderValue.Type {
	case components.AverageOrderValueTypeInteger:
		// averageOrderValue.Integer is populated
	case components.AverageOrderValueTypeNumber:
		// averageOrderValue.Number is populated
}
```

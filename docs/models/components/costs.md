# Costs


## Supported Types

### 

```go
costs := components.CreateCostsInteger(int64{/* values here */})
```

### 

```go
costs := components.CreateCostsNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch costs.Type {
	case components.CostsTypeInteger:
		// costs.Integer is populated
	case components.CostsTypeNumber:
		// costs.Number is populated
}
```

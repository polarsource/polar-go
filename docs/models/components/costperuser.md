# CostPerUser


## Supported Types

### 

```go
costPerUser := components.CreateCostPerUserInteger(int64{/* values here */})
```

### 

```go
costPerUser := components.CreateCostPerUserNumber(float64{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch costPerUser.Type {
	case components.CostPerUserTypeInteger:
		// costPerUser.Integer is populated
	case components.CostPerUserTypeNumber:
		// costPerUser.Number is populated
}
```
